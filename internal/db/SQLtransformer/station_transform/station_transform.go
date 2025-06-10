package main

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/okb97/realestate-scraper/internal/utils"
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
	filepath := "internal/db/SQLtransformer/data/station20250430free.csv"
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file) // ← transform は不要
	reader.Comma = ','            // CSVの区切り文字がカンマの場合
	reader.FieldsPerRecord = -1   // 行ごとのカラム数違いを許容

	// ヘッダー行をスキップ
	_, err = reader.Read()
	if err != nil {
		panic(err)
	}
	lineNumber := 1
	for i := 0; i < 10236; i++ {
		record, err := reader.Read()
		lineNumber++

		if len(record) < 9 {
			log.Printf("⚠️ スキップ：行のカラム数が不足しています (%dカラム): %v", len(record), lineNumber)
			continue
		}

		// 東京都・神奈川県以外はスキップ
		prefCode := record[6]
		if prefCode != "13" && prefCode != "14" {
			continue
		}
		address := record[8] // 住所
		// prefCodeに応じて都県名を補完（addressに"東京都"や"神奈川県"が含まれていない場合）
		if prefCode == "13" && !strings.HasPrefix(address, "東京都") {
			address = "東京都" + address
		} else if prefCode == "14" && !strings.HasPrefix(address, "神奈川県") {
			address = "神奈川県" + address
		}

		// CSVの各フィールドを取得
		trainLineId := record[5]   // 路線ID
		stationId := record[0]     // 駅ID
		stationName := record[2]   // 駅名
		postalCodeRaw := record[7] // 郵便番号
		postalCode := strings.ReplaceAll(postalCodeRaw, "-", "")

		pref, city, town, _ := utils.DivideAddress((address))
		addressId, _ := utils.GetAddressID(pref, city, town)
		createdAt := time.Now().Format("2006-01-02 15:04:05")

		if addressId == -1 {
			log.Printf("❌ address_idが取得できませんでした: 元住所=%s, pref=%s, city=%s, town=%s", record[8], pref, city, town)
			continue
		}

		var exists bool
		err = db.QueryRow(`SELECT EXISTS (SELECT 1 FROM station WHERE station_id = $1)`, stationId).Scan(&exists)
		if exists {
			//log.Printf("⚠️ station_id=%s,station_name=%s は既に存在するためスキップします", stationId, stationName)
			continue
		}

		log.Printf("Trying to insert: station_id=%s, train_line_id=%s, station_name=%s", stationId, trainLineId, stationName)
		query := `
INSERT INTO station (
    station_id, train_line_id, station_name, postal_code, address_id, address, 
    delete_flag, register_date_time, update_date_time, register_function, update_function
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

		_, err = db.Exec(query,
			stationId,
			trainLineId,
			stationName,
			postalCode,
			addressId,
			address,
			false,
			createdAt,
			createdAt,
			"station_transform",
			"station_transform",
		)
		if err != nil {
			log.Fatalf("Insert error: %v", err)
		}
	}

}
