# 不動産スクレーパー

SUUMO（日本の不動産サイト）から中古マンションデータを収集し、収集したデータにアクセスするためのAPIを提供するWebアプリケーションです。

## 特徴

- SUUMOの中古マンション情報を自動収集
- PostgreSQLデータベースでデータ管理
- REST API経由でのデータアクセス
- React TypeScript製のWebフロントエンド
- 駅・バス停の交通情報も含む詳細データ
- 地域別・エリア別の柔軟なスクレイピング対象選択
- コマンドライン引数による自動化対応

## アーキテクチャ

### バックエンド（Go）
- **API サーバー**: Ginフレームワークを使用したREST API（ポート8080）
- **スクレーパー**: Collyフレームワークを使用したWebスクレーピング
- **データベース**: PostgreSQL with トランザクション管理

### フロントエンド
- React + TypeScript + Vite
- マンション一覧表示機能

## セットアップ

### 必要な環境
- Go 1.24.3+
- PostgreSQL
- Node.js & npm

### データベース設定
```bash
# PostgreSQLテーブル作成
psql -d your_database -f internal/SQL/DDL/CREATE_TABLE.sql
```

### 環境変数設定
プロジェクトルートに`.env`ファイルを作成し、データベース接続情報を設定してください。

### 依存関係のインストール
```bash
# Go依存関係
go mod tidy

# フロントエンド依存関係
cd frontend
npm install
```

## 使用方法

### データ収集（スクレーパー）

#### インタラクティブモード
```bash
go run cmd/scraper/main.go
```

#### コマンドライン引数による自動実行
```bash
# ヘルプ表示
go run cmd/scraper/main.go -help

# 利用可能地域一覧表示
go run cmd/scraper/main.go -list

# 東京都心部全体をスクレイピング
go run cmd/scraper/main.go -area=1

# 複数地域をスクレイピング
go run cmd/scraper/main.go -area=1,2,3

# 特定の区市町村のみスクレイピング
go run cmd/scraper/main.go -area=1:100      # 千代田区のみ
go run cmd/scraper/main.go -area=1:100,101  # 千代田区と中央区
go run cmd/scraper/main.go -area=7:100,101  # 横浜市鶴見区・神奈川区

# 地域とURLの混在指定
go run cmd/scraper/main.go -area=1,7:100,8:100,101
```

#### スクレイピング対象地域

**東京都（地域ID: 1-6）**
- 1: 東京都心部（千代田区、中央区、港区、新宿区、文京区、渋谷区）
- 2: 東京23区東部（台東区、墨田区、江東区、荒川区、足立区、葛飾区、江戸川区）
- 3: 東京23区南部（品川区、目黒区、大田区、世田谷区）
- 4: 東京23区西部（中野区、杉並区、練馬区）
- 5: 東京23区北部（豊島区、北区、板橋区）
- 6: 東京都下（八王子市、立川市、武蔵野市など38市町村）

**神奈川県（地域ID: 7-10）**
- 7: 横浜市（18区）
- 8: 川崎市（7区）
- 9: 相模原市（3区）
- 10: 神奈川県その他（横須賀市、平塚市、鎌倉市など20市町村）

**URL採番システム**
- 各地域内のURL IDは100から順番に採番
- 例：1:100=千代田区、1:101=中央区、7:100=横浜市鶴見区

### APIサーバー起動
```bash
go run cmd/api/main.go
```

### フロントエンド開発サーバー
```bash
cd frontend
npm run dev
```

## 設定ファイル

### CLAUDE.md
プロジェクトルートの`CLAUDE.md`ファイルには、Claude Code用の指示が含まれています：
- データベーステーブル構造の編集制限
- プロジェクト概要とアーキテクチャ説明
- 共通コマンドと環境設定手順

### 地域設定ファイル
- `config/areas.go`: スクレイピング対象地域とURL設定
- `config/area_selector.go`: 地域選択とコマンドライン引数処理

## API エンドポイント

### 中古マンション一覧取得
```
GET /api/used-condos
```

**レスポンス例:**
```json
{
  "used_condos": [
    {
      "id": 1,
      "used_condo_name": "サンプルマンション",
      "price": 5500,
      "madori": "3LDK",
      "menseki": 70.25,
      "chikusu": 15,
      "address": "東京都新宿区...",
      "stations": [...],
      "bus_stops": [...]
    }
  ]
}
```

## データベーススキーマ

主要テーブル：
- `used_condos`: 物件データ（価格、間取り、築年数、面積など）
- `station`, `train_line`: 鉄道情報（駅名、路線名）
- `bus_stop`, `bus_line`: バス情報（バス停名、路線名）
- `address`: 住所情報（都道府県、市区町村、町域）
- `used_condos_stations`, `used_condos_bus_stop`: 物件と交通機関の関連テーブル

### データ変換・検証機能
- 住所データの正規化とマスターデータとの照合
- 交通情報（駅・バス停）の自動抽出と登録
- 価格・面積等の数値データの検証と変換
- 築年数の自動計算

## 技術スタック

**バックエンド:**
- Go 1.24.3+
- Gin (Web framework)
- Colly (Web scraping)
- PostgreSQL (データベース)
- godotenv (環境変数管理)
- database/sql (データベース接続)

**フロントエンド:**
- React 18
- TypeScript
- Vite (ビルドツール)
- ESLint (コード品質管理)

**開発・運用:**
- Docker対応可能な構成
- トランザクション管理によるデータ整合性保証
- エラーハンドリングとログ出力
- 段階的なページネーション処理

## ライセンス

このプロジェクトは個人開発用です。