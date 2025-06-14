# 画像取得機能 設計書

## 📋 要件定義

### 🎯 基本要件

#### 1. mainCollectorでの画像取得
- **対象**: 物件一覧ページの各物件
- **取得枚数**: 3枚/物件
- **処理タイミング**: 物件一覧ページのスクレイピング時
- **保存対象**: 基本的な物件画像（サムネイル品質）

#### 2. 特定物件の詳細画像取得
- **対象**: 「特定物件一覧」に登録された物件
- **取得場所**: detailCollector（物件詳細ページ）
- **取得枚数**: 詳細ページ内の全画像（通常10-20枚）
- **保存対象**: 高品質な詳細画像

### 💡 追加推奨要件

#### 3. 画像品質・サイズ管理
- **画像形式**: JPEG優先（PNG, WebP対応）
- **品質設定**: JPEG品質75%で統一
- **サイズ制限**: 1画像あたり最大10MB
- **リサイズ**: 必要に応じて自動リサイズ
  - サムネイル: 320×240px
  - 通常画像: 640×480px（一覧ページ）
  - 高解像度: 1024×768px（詳細ページ）

#### 4. エラーハンドリング・信頼性
- **リトライ機能**: ダウンロード失敗時に3回まで再試行
- **タイムアウト**: 1画像あたり30秒でタイムアウト
- **不正画像**: 破損画像の検出・スキップ
- **継続性**: 一部画像の失敗時も処理継続

#### 5. パフォーマンス最適化
- **並行処理**: 最大3枚の画像を同時ダウンロード
- **キャッシュ**: URL/ハッシュベースの重複画像検出
- **進捗表示**: ダウンロード状況のログ出力
- **メモリ管理**: 大量画像処理時のメモリ使用量制御

#### 6. データ管理・整合性
- **重複排除**: 同一画像の重複保存を防止
- **メタデータ管理**: ファイルサイズ、寸法、形式の記録
- **関連性**: 物件データとの紐づけ管理
- **バックアップ**: 外部ストレージへの自動保存（オプション）

#### 7. 設定・運用管理
- **機能ON/OFF**: 画像取得機能の有効/無効切り替え
- **特定物件管理**: 外部ファイルでの特定物件一覧管理
- **ストレージ設定**: 保存先パスの設定
- **容量監視**: ストレージ使用量の監視・警告

## 🏗️ システム設計

### 📁 ディレクトリ構造
```
realestate-scraper/
├── internal/
│   ├── scraper/
│   │   ├── image/                    # 新規作成
│   │   │   ├── downloader.go        # 画像ダウンロード
│   │   │   ├── processor.go         # 画像処理
│   │   │   └── storage.go           # 画像保存管理
│   │   └── used_condo_scraper.go    # 修正対象
│   ├── model/
│   │   └── property_image.go        # 新規作成
│   └── SQL/DDL/
│       └── CREATE_IMAGES_TABLE.sql  # 新規作成
├── config/
│   ├── special_properties.go        # 新規作成
│   └── image_config.go              # 新規作成
├── document/
│   └── design/
│       └── image_acquisition_design.md  # 本文書
└── data/
    ├── special_properties.txt        # 新規作成
    └── images/                       # 新規作成
        ├── thumbnails/
        ├── main/
        └── detail/
```

### 🗃️ データベース設計

#### property_images テーブル
```sql
CREATE TABLE property_images (
    image_id SERIAL PRIMARY KEY,
    used_condo_id INTEGER NOT NULL,
    image_type VARCHAR(20) NOT NULL,     -- 'thumbnail', 'main_list', 'detail'
    image_url VARCHAR(2048),             -- 元のSUUMO画像URL
    image_data BYTEA,                    -- 画像バイナリデータ
    image_size INTEGER,                  -- ファイルサイズ（バイト）
    width INTEGER,                       -- 画像幅
    height INTEGER,                      -- 画像高さ
    file_format VARCHAR(10),             -- 'jpeg', 'png', 'webp'
    quality INTEGER,                     -- 画像品質（1-100）
    sort_order INTEGER,                  -- 表示順序
    is_main_image BOOLEAN,               -- メイン画像フラグ
    hash_value VARCHAR(64),              -- 重複検出用ハッシュ
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    delete_flag BOOLEAN,
    FOREIGN KEY (used_condo_id) REFERENCES used_condos (used_condo_id)
);
```

### 🔧 主要コンポーネント設計

#### 1. ImageDownloader（画像ダウンローダー）
```go
type ImageDownloader struct {
    client     *http.Client
    config     *ImageConfig
    maxRetries int
    timeout    time.Duration
}

// メインメソッド
func (id *ImageDownloader) DownloadMainListImages(e *colly.HTMLElement) ([]*ImageInfo, error)
func (id *ImageDownloader) DownloadDetailImages(e *colly.HTMLElement) ([]*ImageInfo, error)
func (id *ImageDownloader) SaveImagesToDB(tx *sql.Tx, usedCondoID int, images []*ImageInfo) error
```

#### 2. SpecialPropertyManager（特定物件管理）
```go
type SpecialPropertyManager struct {
    properties []SpecialProperty
    filePath   string
}

// メインメソッド
func (spm *SpecialPropertyManager) LoadFromFile() error
func (spm *SpecialPropertyManager) IsSpecialProperty(name, url string) bool
func (spm *SpecialPropertyManager) GetSpecialProperties() []SpecialProperty
```

#### 3. ImageProcessor（画像処理）
```go
type ImageProcessor struct {
    config *ImageConfig
}

// メインメソッド
func (ip *ImageProcessor) ResizeImage(data []byte, maxWidth, maxHeight int) []byte
func (ip *ImageProcessor) CompressImage(data []byte, quality int) []byte
func (ip *ImageProcessor) GenerateHash(data []byte) string
```

### 📊 処理フロー設計

#### mainCollectorでの画像取得フロー
```
1. 物件一覧ページをスクレイピング
2. 各物件について：
   a. 物件の基本情報を取得
   b. 3枚の画像URLを抽出
   c. 画像をダウンロード・処理
   d. データベースに保存
3. 特定物件判定
   a. 特定物件一覧と照合
   b. 該当する場合は詳細ページフラグを設定
```

#### detailCollectorでの画像取得フロー
```
1. 物件詳細ページをスクレイピング
2. 特定物件判定をチェック
3. 特定物件の場合：
   a. 詳細ページ内の全画像URLを抽出
   b. 既存画像との重複チェック
   c. 新規画像をダウンロード・処理
   d. 高解像度画像として保存
```

### 🔍 画像種別・セレクタ設計

#### mainCollector用セレクタ
```go
var mainListSelectors = []string{
    ".property_unit-image img",           // メイン画像
    ".property_unit-photo img",           // 補助画像1
    ".cassetteitem_other-thumbnail img",  // 補助画像2
}
```

#### detailCollector用セレクタ
```go
var detailPageSelectors = map[string]string{
    "main":     ".slider-main img, .photo-main img",
    "exterior": ".photo-exterior img, [data-type='exterior'] img",
    "interior": ".photo-interior img, [data-type='interior'] img", 
    "layout":   ".photo-layout img, [data-type='layout'] img",
    "other":    ".property-photo img, .detail-images img",
}
```

### ⚙️ 設定管理設計

#### ImageConfig構造体
```go
type ImageConfig struct {
    Enabled           bool     `json:"enabled"`             // 画像取得機能有効/無効
    MaxConcurrent     int      `json:"max_concurrent"`      // 同時ダウンロード数
    RetryCount        int      `json:"retry_count"`         // リトライ回数
    TimeoutSeconds    int      `json:"timeout_seconds"`     // タイムアウト秒数
    MaxImageSize      int64    `json:"max_image_size"`      // 最大画像サイズ
    JpegQuality       int      `json:"jpeg_quality"`        // JPEG品質
    StoragePath       string   `json:"storage_path"`        // 保存先パス
    EnableResize      bool     `json:"enable_resize"`       // リサイズ有効/無効
    ThumbnailSize     [2]int   `json:"thumbnail_size"`      // サムネイルサイズ [幅, 高さ]
    StandardSize      [2]int   `json:"standard_size"`       // 標準サイズ
    HighResSize       [2]int   `json:"high_res_size"`       // 高解像度サイズ
    EnableDeduplication bool   `json:"enable_deduplication"` // 重複排除有効/無効
}
```

## 📈 ストレージ要件

### 容量設計
- **年間総容量**: 約111GB（画像込み）
- **月間増加量**: 約9GB
- **1物件平均**: 約780KB（画像6枚想定）

### パフォーマンス要件
- **ダウンロード速度**: 平均1MB/s以上
- **処理能力**: 1物件あたり30秒以内
- **メモリ使用量**: 同時処理時最大500MB

## 🚨 リスク・制約事項

### 技術的制約
1. **SUUMO側の制限**: レート制限、IP制限の可能性
2. **画像URL変更**: SUUMOの画像URL構造変更リスク
3. **容量制限**: ローカルストレージの容量制限

### 運用上の考慮事項
1. **法的制約**: 著作権、利用規約の遵守
2. **負荷対策**: SUUMOサーバーへの負荷軽減
3. **データ管理**: 個人情報を含む画像の適切な管理

## 🎯 実装優先度

### Phase 1（高優先度）
- [x] 基本的な画像ダウンロード機能
- [x] データベーステーブル作成
- [x] 特定物件一覧管理

### Phase 2（中優先度）
- [ ] 画像処理・リサイズ機能
- [ ] エラーハンドリング強化
- [ ] 重複排除機能

### Phase 3（低優先度）
- [ ] 並行処理最適化
- [ ] 外部ストレージ連携
- [ ] 管理画面・統計機能

## 📝 関連ドキュメント

- [データベース設計書](./database_design.md)
- [API仕様書](./api_specification.md)
- [運用マニュアル](./operation_manual.md)