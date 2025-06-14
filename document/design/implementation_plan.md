# 画像取得機能 実装計画書

## 📋 修正・新規作成ファイル一覧

### 🆕 新規作成ファイル

#### 1. 画像処理関連
- `internal/scraper/image/downloader.go` - 画像ダウンロード処理
- `internal/scraper/image/processor.go` - 画像処理（リサイズ、圧縮）
- `internal/scraper/image/storage.go` - 画像保存管理
- `internal/model/property_image.go` - 画像データモデル

#### 2. 設定・管理関連
- `config/special_properties.go` - 特定物件管理
- `config/image_config.go` - 画像設定管理
- `data/special_properties.txt` - 特定物件一覧ファイル

#### 3. データベース関連
- `internal/SQL/DDL/CREATE_IMAGES_TABLE.sql` - 画像テーブル定義
- `internal/db/property_image_db.go` - 画像DB操作

#### 4. 画像保存ディレクトリ
- `data/images/thumbnails/` - サムネイル画像
- `data/images/main/` - 一覧ページ画像  
- `data/images/detail/` - 詳細ページ画像

### 🔧 修正対象ファイル

#### 1. メインスクレーパー
- `internal/scraper/used_condo_scraper.go`
  - ImageDownloaderの組み込み
  - mainCollectorに画像取得処理追加
  - detailCollectorに特定物件画像取得追加

#### 2. データベース関連
- `internal/db/used_condo_db.go`
  - 画像保存処理の追加

#### 3. 設定関連
- `cmd/scraper/main.go`
  - 画像設定の読み込み
  - 特定物件管理の初期化

## 🔧 主要関数・メソッド設計

### ImageDownloader（internal/scraper/image/downloader.go）

```go
type ImageDownloader struct {
    client     *http.Client
    config     *ImageConfig
    maxRetries int
    timeout    time.Duration
}

// メイン機能
func NewImageDownloader(config *ImageConfig) *ImageDownloader
func (id *ImageDownloader) DownloadMainListImages(e *colly.HTMLElement, usedCondoID int) ([]*ImageInfo, error)
func (id *ImageDownloader) DownloadDetailImages(e *colly.HTMLElement, usedCondoID int) ([]*ImageInfo, error)

// 内部処理
func (id *ImageDownloader) downloadImage(imageURL string, imageType ImageType, sortOrder int) (*ImageInfo, error)
func (id *ImageDownloader) fetchImageData(imageURL string) ([]byte, error)
func (id *ImageDownloader) analyzeImage(imageURL string, imageType ImageType, data []byte, sortOrder int) (*ImageInfo, error)
func (id *ImageDownloader) isDuplicateURL(images []*ImageInfo, url string) bool
```

### SpecialPropertyManager（config/special_properties.go）

```go
type SpecialPropertyManager struct {
    properties []SpecialProperty
    filePath   string
}

// メイン機能
func NewSpecialPropertyManager(filePath string) *SpecialPropertyManager
func (spm *SpecialPropertyManager) LoadFromFile() error
func (spm *SpecialPropertyManager) IsSpecialProperty(propertyName, propertyURL string) bool
func (spm *SpecialPropertyManager) GetSpecialProperties() []SpecialProperty

// 管理機能
func (spm *SpecialPropertyManager) AddSpecialProperty(property SpecialProperty)
func (spm *SpecialPropertyManager) SaveToFile() error
func (spm *SpecialPropertyManager) GetStats() map[string]int
```

### ImageProcessor（internal/scraper/image/processor.go）

```go
type ImageProcessor struct {
    config *ImageConfig
}

// 画像処理
func NewImageProcessor(config *ImageConfig) *ImageProcessor
func (ip *ImageProcessor) ProcessImage(data []byte, imageType ImageType) ([]byte, error)
func (ip *ImageProcessor) ResizeImage(data []byte, maxWidth, maxHeight int) []byte
func (ip *ImageProcessor) CompressImage(data []byte, quality int) []byte
func (ip *ImageProcessor) GenerateHash(data []byte) string
func (ip *ImageProcessor) ValidateImage(data []byte) error
```

### PropertyImageDB（internal/db/property_image_db.go）

```go
// データベース操作
func InsertPropertyImage(ctx context.Context, tx *sql.Tx, image *PropertyImage) error
func UpdatePropertyImage(ctx context.Context, tx *sql.Tx, image *PropertyImage) error
func DeletePropertyImagesByCondoID(ctx context.Context, tx *sql.Tx, usedCondoID int) error
func GetPropertyImagesByCondoID(ctx context.Context, db *sql.DB, usedCondoID int) ([]*PropertyImage, error)
func GetImageStatistics(ctx context.Context, db *sql.DB) (*ImageStatistics, error)
```

## 🔄 処理フロー詳細

### mainCollector修正箇所

```go
// used_condo_scraper.go の mainCollector.OnHTML 内に追加
mainCollector.OnHTML(".property_unit-title a", func(e *colly.HTMLElement) {
    // 既存の処理...
    link := e.Request.AbsoluteURL(e.Attr("href"))
    
    // 🆕 画像取得処理を追加
    if imageDownloader.IsEnabled() {
        images, err := imageDownloader.DownloadMainListImages(e, 0) // usedCondoIDは後で更新
        if err != nil {
            log.Printf("一覧画像取得失敗: %v", err)
        } else {
            // 取得した画像をコンテキストに保存（詳細処理で使用）
            e.Request.Ctx.Put("images", images)
        }
    }
    
    detailCollector.Visit(link)
})
```

### detailCollector修正箇所

```go
// used_condo_scraper.go の detailCollector.OnHTML 内に追加
detailCollector.OnHTML("#mainContents", func(e *colly.HTMLElement) {
    // 既存のトランザクション処理...
    
    // 物件データの保存後、usedCondoIDが確定
    usedCondoID, updated, err := db.InsertUsedCondo(ctx, tx, usedCondoModel)
    if err != nil {
        // エラー処理...
        return
    }
    
    // 🆕 画像保存処理を追加
    if imageDownloader.IsEnabled() && usedCondoID != 0 {
        // mainCollectorで取得した画像を保存
        if mainImages, ok := e.Request.Ctx.GetAny("images").([]*ImageInfo); ok {
            if err := imageDownloader.SaveImagesToDB(tx, usedCondoID, mainImages); err != nil {
                log.Printf("一覧画像保存失敗: %v", err)
            }
        }
        
        // 特定物件の場合は詳細画像も取得
        if specialPropertyManager.IsSpecialProperty(usedCondoModel.UsedCondoName, usedCondoModel.Url) {
            detailImages, err := imageDownloader.DownloadDetailImages(e, usedCondoID)
            if err != nil {
                log.Printf("詳細画像取得失敗: %v", err)
            } else {
                if err := imageDownloader.SaveImagesToDB(tx, usedCondoID, detailImages); err != nil {
                    log.Printf("詳細画像保存失敗: %v", err)
                }
            }
        }
    }
    
    // 既存の交通データ処理...
})
```

## ⚙️ 設定ファイル設計

### config/image_config.json
```json
{
    "enabled": true,
    "max_concurrent": 3,
    "retry_count": 3,
    "timeout_seconds": 30,
    "max_image_size": 10485760,
    "jpeg_quality": 75,
    "storage_path": "./data/images",
    "enable_resize": true,
    "thumbnail_size": [320, 240],
    "standard_size": [640, 480],
    "high_res_size": [1024, 768],
    "enable_deduplication": true,
    "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36",
    "rate_limit": {
        "requests_per_second": 5,
        "burst_size": 10
    }
}
```

### data/special_properties.txt（サンプル）
```
# 特定物件一覧
# 形式: 物件名|URL|キーワード1,キーワード2|説明

# 高級マンションシリーズ
パークコート||パークコート,三井,高級|三井不動産の高級マンション
ザ・パークハウス||パークハウス,三菱地所,高級|三菱地所の高級マンション
プラウド||プラウド,野村,高級|野村不動産の高級マンション
ブリリア||ブリリア,東京建物,高級|東京建物の高級マンション

# タワーマンション
タワー||タワー,高層|タワーマンション全般
スカイ||スカイ,高層|高層マンション

# 特定物件（URL指定）
パークマンション青山|https://suumo.jp/ms/chuko/tokyo/nc_12345678/|パークマンション,青山|個別指定物件
```

## 🧪 テスト計画

### 単体テスト
- [ ] ImageDownloader各メソッドのテスト
- [ ] SpecialPropertyManagerのテスト
- [ ] 画像処理関数のテスト

### 結合テスト
- [ ] mainCollector + 画像取得の統合テスト
- [ ] detailCollector + 特定物件判定の統合テスト
- [ ] データベース保存の統合テスト

### システムテスト
- [ ] 実際のSUUMOページでのE2Eテスト
- [ ] 大量データでの性能テスト
- [ ] エラー条件での動作テスト

## 📊 監視・ログ設計

### ログ出力項目
```
[INFO] 画像取得開始: 物件ID=12345, 種別=main_list
[INFO] 画像ダウンロード成功: URL=https://..., サイズ=245KB
[WARN] 画像ダウンロード失敗: URL=https://..., 試行=2/3, エラー=timeout
[ERROR] 画像保存失敗: 物件ID=12345, エラー=database connection lost
[INFO] 特定物件検出: 物件名=パークコート青山, 詳細画像取得開始
[INFO] 画像取得完了: 物件ID=12345, 一覧=3枚, 詳細=15枚, 総サイズ=2.3MB
```

### 統計情報
- 日次画像取得数
- エラー率
- 平均ダウンロード時間
- ストレージ使用量

## 🚀 段階的実装アプローチ

### Step 1: 基盤準備
1. データベーステーブル作成
2. 基本的な設定ファイル作成
3. ディレクトリ構造準備

### Step 2: 基本機能実装
1. ImageDownloaderの基本機能
2. SpecialPropertyManagerの実装
3. mainCollectorの画像取得追加

### Step 3: 詳細機能実装
1. detailCollectorの特定物件処理
2. 画像処理・最適化機能
3. エラーハンドリング強化

### Step 4: 最適化・運用準備
1. 並行処理最適化
2. 監視・ログ機能
3. 性能チューニング