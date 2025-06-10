# 不動産スクレーパー

SUUMO（日本の不動産サイト）から中古マンションデータを収集し、収集したデータにアクセスするためのAPIを提供するWebアプリケーションです。

## 特徴

- SUUMOの中古マンション情報を自動収集
- PostgreSQLデータベースでデータ管理
- REST API経由でのデータアクセス
- React TypeScript製のWebフロントエンド
- 駅・バス停の交通情報も含む詳細データ

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
```bash
go run cmd/scraper/main.go
```

### APIサーバー起動
```bash
go run cmd/api/main.go
```

### フロントエンド開発サーバー
```bash
cd frontend
npm run dev
```

## API エンドポイント

### 中古マンション一覧取得
```
GET /api/used-condos
```

## データベーススキーマ

主要テーブル：
- `used_condos`: 物件データ
- `station`, `train_line`: 鉄道情報
- `bus_stop`, `bus_line`: バス情報
- `address`: 住所情報
- `used_condos_stations`, `used_condos_bus_stop`: 関連テーブル

## 技術スタック

**バックエンド:**
- Go
- Gin (Web framework)
- Colly (Web scraping)
- PostgreSQL
- godotenv

**フロントエンド:**
- React
- TypeScript
- Vite
- ESLint

## ライセンス

このプロジェクトは個人開発用です。