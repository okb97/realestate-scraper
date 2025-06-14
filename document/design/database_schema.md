# 画像取得機能 データベース設計書

## 📊 新規テーブル設計

### property_images テーブル

```sql
-- 物件画像テーブル
DROP TABLE IF EXISTS property_images CASCADE;

CREATE TABLE property_images (
    image_id SERIAL PRIMARY KEY,
    used_condo_id INTEGER NOT NULL,
    image_type VARCHAR(20) NOT NULL DEFAULT 'main',
    image_url VARCHAR(2048),
    image_data BYTEA,
    image_size INTEGER,
    width INTEGER,
    height INTEGER,
    file_format VARCHAR(10) DEFAULT 'jpeg',
    quality INTEGER DEFAULT 75,
    hash_value VARCHAR(64),
    sort_order INTEGER DEFAULT 0,
    is_main_image BOOLEAN DEFAULT FALSE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    delete_flag BOOLEAN DEFAULT FALSE,
    
    FOREIGN KEY (used_condo_id) REFERENCES used_condos (used_condo_id) ON DELETE CASCADE,
    UNIQUE(used_condo_id, image_type, sort_order)
);
```

### インデックス設計

```sql
-- パフォーマンス最適化用インデックス
CREATE INDEX idx_property_images_condo_id ON property_images(used_condo_id);
CREATE INDEX idx_property_images_main ON property_images(used_condo_id, is_main_image) WHERE is_main_image = TRUE;
CREATE INDEX idx_property_images_type ON property_images(image_type);
CREATE INDEX idx_property_images_hash ON property_images(hash_value) WHERE hash_value IS NOT NULL;
CREATE INDEX idx_property_images_created ON property_images(created_at);
CREATE INDEX idx_property_images_size ON property_images(image_size) WHERE image_size > 0;
```

## 🏷️ フィールド定義詳細

### 基本情報フィールド
| フィールド名 | 型 | 制約 | 説明 |
|-------------|----|----|------|
| image_id | SERIAL | PRIMARY KEY | 画像ID（自動採番） |
| used_condo_id | INTEGER | NOT NULL, FK | 物件ID（外部キー） |
| image_type | VARCHAR(20) | NOT NULL | 画像種別 |
| image_url | VARCHAR(2048) | - | 元のSUUMO画像URL |

### 画像データフィールド
| フィールド名 | 型 | 制約 | 説明 |
|-------------|----|----|------|
| image_data | BYTEA | - | 画像バイナリデータ |
| image_size | INTEGER | - | ファイルサイズ（バイト） |
| width | INTEGER | - | 画像幅（ピクセル） |
| height | INTEGER | - | 画像高さ（ピクセル） |
| file_format | VARCHAR(10) | DEFAULT 'jpeg' | ファイル形式 |
| quality | INTEGER | DEFAULT 75 | 画像品質（JPEG用） |
| hash_value | VARCHAR(64) | - | 重複検出用ハッシュ |

### メタデータフィールド
| フィールド名 | 型 | 制約 | 説明 |
|-------------|----|----|------|
| sort_order | INTEGER | DEFAULT 0 | 表示順序 |
| is_main_image | BOOLEAN | DEFAULT FALSE | メイン画像フラグ |
| description | TEXT | - | 画像説明 |
| created_at | TIMESTAMP | DEFAULT NOW() | 作成日時 |
| updated_at | TIMESTAMP | DEFAULT NOW() | 更新日時 |
| delete_flag | BOOLEAN | DEFAULT FALSE | 削除フラグ |

## 📋 画像種別（image_type）定義

### mainCollector取得画像
| 種別 | 値 | 説明 | 取得枚数 |
|------|----|----|----------|
| サムネイル | thumbnail | 一覧ページのメイン画像 | 1枚 |
| 補助画像1 | main_list_1 | 一覧ページの補助画像1 | 1枚 |
| 補助画像2 | main_list_2 | 一覧ページの補助画像2 | 1枚 |

### detailCollector取得画像
| 種別 | 値 | 説明 | 取得枚数 |
|------|----|----|----------|
| 詳細メイン | detail_main | 詳細ページのメイン画像 | 1-3枚 |
| 外観 | exterior | 建物外観画像 | 2-5枚 |
| 内装 | interior | 室内画像 | 5-10枚 |
| 間取り | layout | 間取り図 | 1-2枚 |
| その他 | other | その他の画像 | 0-5枚 |

## 📈 ビュー・関数設計

### 画像統計ビュー
```sql
CREATE OR REPLACE VIEW property_image_stats AS
SELECT 
    used_condo_id,
    COUNT(*) as total_images,
    COUNT(CASE WHEN image_type = 'thumbnail' THEN 1 END) as thumbnail_count,
    COUNT(CASE WHEN image_type LIKE '%list%' THEN 1 END) as list_images_count,
    COUNT(CASE WHEN image_type LIKE 'detail%' OR image_type IN ('exterior', 'interior', 'layout', 'other') THEN 1 END) as detail_images_count,
    SUM(image_size) as total_size_bytes,
    ROUND(SUM(image_size)::NUMERIC / 1024 / 1024, 2) as total_size_mb,
    MAX(CASE WHEN is_main_image THEN image_id END) as main_image_id,
    MIN(created_at) as first_image_created,
    MAX(created_at) as last_image_created
FROM property_images 
WHERE delete_flag = FALSE
GROUP BY used_condo_id;
```

### 画像サマリビュー
```sql
CREATE OR REPLACE VIEW image_summary AS
SELECT 
    DATE(created_at) as date,
    image_type,
    COUNT(*) as image_count,
    SUM(image_size) as total_size,
    AVG(image_size) as avg_size,
    MIN(image_size) as min_size,
    MAX(image_size) as max_size
FROM property_images
WHERE delete_flag = FALSE
GROUP BY DATE(created_at), image_type
ORDER BY date DESC, image_type;
```

### 重複画像検出関数
```sql
CREATE OR REPLACE FUNCTION find_duplicate_images()
RETURNS TABLE(
    duplicate_count INTEGER,
    hash_value VARCHAR(64),
    image_urls TEXT[],
    total_size BIGINT
) AS $$
BEGIN
    RETURN QUERY
    SELECT 
        COUNT(*)::INTEGER as duplicate_count,
        pi.hash_value,
        ARRAY_AGG(pi.image_url) as image_urls,
        SUM(pi.image_size) as total_size
    FROM property_images pi
    WHERE pi.hash_value IS NOT NULL 
    AND pi.delete_flag = FALSE
    GROUP BY pi.hash_value
    HAVING COUNT(*) > 1
    ORDER BY COUNT(*) DESC;
END;
$$ LANGUAGE plpgsql;
```

### 画像削除関数
```sql
CREATE OR REPLACE FUNCTION cleanup_old_images(days_old INTEGER DEFAULT 365)
RETURNS INTEGER AS $$
DECLARE
    deleted_count INTEGER;
BEGIN
    UPDATE property_images 
    SET delete_flag = TRUE, updated_at = CURRENT_TIMESTAMP
    WHERE created_at < CURRENT_DATE - INTERVAL '1 day' * days_old
    AND delete_flag = FALSE;
    
    GET DIAGNOSTICS deleted_count = ROW_COUNT;
    RETURN deleted_count;
END;
$$ LANGUAGE plpgsql;
```

## 🔍 クエリ例

### 基本的な画像取得
```sql
-- 物件の全画像を取得
SELECT image_id, image_type, image_size, width, height, is_main_image
FROM property_images 
WHERE used_condo_id = 12345 
AND delete_flag = FALSE
ORDER BY image_type, sort_order;

-- メイン画像のみ取得
SELECT image_data 
FROM property_images 
WHERE used_condo_id = 12345 
AND is_main_image = TRUE 
AND delete_flag = FALSE
LIMIT 1;
```

### 統計情報取得
```sql
-- 日別画像取得統計
SELECT 
    DATE(created_at) as date,
    COUNT(*) as total_images,
    COUNT(DISTINCT used_condo_id) as properties_with_images,
    SUM(image_size) / 1024 / 1024 as total_mb
FROM property_images
WHERE delete_flag = FALSE
GROUP BY DATE(created_at)
ORDER BY date DESC
LIMIT 30;

-- 画像種別統計
SELECT 
    image_type,
    COUNT(*) as count,
    AVG(image_size) as avg_size_bytes,
    SUM(image_size) / 1024 / 1024 as total_mb
FROM property_images
WHERE delete_flag = FALSE
GROUP BY image_type
ORDER BY count DESC;
```

### パフォーマンス監視
```sql
-- 大きな画像の検出
SELECT used_condo_id, image_type, image_size, width, height
FROM property_images
WHERE image_size > 5 * 1024 * 1024  -- 5MB以上
AND delete_flag = FALSE
ORDER BY image_size DESC;

-- 重複画像の確認
SELECT * FROM find_duplicate_images()
WHERE duplicate_count > 2;
```

## 💾 ストレージ最適化

### パーティショニング（将来的な拡張）
```sql
-- 月別パーティション（大量データ対応）
CREATE TABLE property_images_y2024m01 PARTITION OF property_images
FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');

CREATE TABLE property_images_y2024m02 PARTITION OF property_images
FOR VALUES FROM ('2024-02-01') TO ('2024-03-01');
```

### 圧縮・アーカイブ戦略
```sql
-- 古い画像の圧縮フラグ
ALTER TABLE property_images ADD COLUMN is_compressed BOOLEAN DEFAULT FALSE;
ALTER TABLE property_images ADD COLUMN original_size INTEGER;

-- アーカイブテーブル
CREATE TABLE property_images_archive (LIKE property_images INCLUDING ALL);
```

## 🔐 セキュリティ・権限設計

### アクセス権限
```sql
-- 読み取り専用ユーザー
CREATE ROLE image_reader;
GRANT SELECT ON property_images TO image_reader;
GRANT SELECT ON property_image_stats TO image_reader;

-- アプリケーションユーザー
CREATE ROLE image_app;
GRANT SELECT, INSERT, UPDATE ON property_images TO image_app;
GRANT USAGE ON SEQUENCE property_images_image_id_seq TO image_app;
```

### データ保護
```sql
-- 個人情報を含む可能性のある画像の暗号化（将来的な機能）
ALTER TABLE property_images ADD COLUMN encryption_key_id INTEGER;
ALTER TABLE property_images ADD COLUMN is_encrypted BOOLEAN DEFAULT FALSE;
```