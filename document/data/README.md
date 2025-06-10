# データ仕様書

このドキュメントは`internal/db/SQLtransformer/data/`に格納されているデータファイルの仕様を記載しています。

## ディレクトリ構造

```
data/
├── addresses/          # 住所データ
├── bus_stops/         # バス停データ  
├── line/              # 路線データ（鉄道・バス）
└── stations/          # 駅データ
```

## データファイル一覧

| カテゴリ | ファイル名 | 概要 |
|---------|------------|------|
| 駅 | [station20250430free.csv](./stations.md) | 全国駅データ（駅データ.jp） |
| 駅 | [station_utf8.csv](./stations.md) | 駅データ（UTF-8版） |
| 路線 | [line20250430free.csv](./lines.md) | 全国鉄道路線データ |
| 路線 | [tokyo_bus_routes_kokudonumerical.csv](./bus_routes.md) | 東京都バス路線（国土数値情報） |
| 路線 | [kanagawa_bus_routes_kokudonumerical.csv](./bus_routes.md) | 神奈川県バス路線（国土数値情報） |
| バス停 | [tokyo_bus_stops_kokudonumerical.csv](./bus_stops.md) | 東京都バス停（国土数値情報） |
| バス停 | [kanagawa_bus_stops_kokudonumerical.csv](./bus_stops.md) | 神奈川県バス停（国土数値情報） |
| 住所 | [13TOKYO.CSV](./addresses.md) | 東京都住所データ |
| 住所 | [14KANAGAWA.CSV](./addresses.md) | 神奈川県住所データ |

## データソース

- **駅データ.jp**: 鉄道駅・路線情報（https://ekidata.jp/）
- **国土数値情報**: バス停・バス路線情報（国土交通省）
- **住所データ**: 郵便番号・住所情報

## 更新履歴

- 2025-06-10: 初版作成
- 2025-06-10: 国土数値情報バスデータ追加