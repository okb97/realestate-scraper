package main

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".envの読み込みに失敗しました: %v", err)
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("DB接続失敗: %v", err)
	}
	defer db.Close()

	// 駅データCSVファイルのオープン
	file, err := os.Open("internal/db/SQLtransformer/data/line20250430free.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ',' // CSVの区切り文字がカンマの場合

	// ヘッダー行をスキップ
	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("CSV読み込みエラー: %v", err)
		}
		if len(record) < 10 {
			log.Printf("⚠️ 不正な行（列が足りません）: %+v", record)
			continue
		}

		// CSVの各フィールドを取得
		trainLineId := record[0] // 路線ID
		trainLineName := record[2]
		createdAt := time.Now().Format("2006-01-02 15:04:05")

		log.Printf("Trying to insert: train_line_id=%s, train_line_name=%s", trainLineId, trainLineName)
		query := `
INSERT INTO train_line (
    train_line_id, train_line_name, 
    delete_flag, register_date_time, update_date_time, register_function, update_function
) VALUES ($1, $2, $3, $4, $5, $6, $7)`

		_, err = db.Exec(query,
			trainLineId,
			trainLineName,
			false,
			createdAt,
			createdAt,
			"train_line_transform",
			"train_line_transform",
		)
		if err != nil {
			log.Fatalf("Insert error: %v", err)
		}
	}
}
