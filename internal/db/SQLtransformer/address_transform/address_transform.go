package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const KANAGAWA_ADDRESS = "14KANAGAWA.CSV"
const TOKYO_ADDRESS = "13TOKYO.CSV"
const basePath = "internal/db/SQLtransformer/data/"

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

	files := []string{KANAGAWA_ADDRESS, TOKYO_ADDRESS}
	for _, fileName := range files {
		err := processCSV(db, fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー：%v\n", err)
		}
	}
}

func processCSV(db *sql.DB, fileName string) error {
	fullPath := filepath.Join(basePath, fileName)
	file, err := os.Open(fullPath)
	if err != nil {
		return fmt.Errorf("ファイルオープン失敗: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(transform.NewReader(bufio.NewReader(file), japanese.ShiftJIS.NewDecoder()))
	_, _ = reader.Read() // ヘッダーを読み飛ばす

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		postal_code := record[2]
		prefecture := record[6]
		city := record[7]
		town := record[8]

		now := time.Now()

		_, err = db.Exec(`
			INSERT INTO address (
				postal_code, prefecture, city, town, delete_flag, register_date_time, update_date_time
			) VALUES ($1, $2, $3, $4, $5, $6, $7)
		`, postal_code, prefecture, city, town, false, now, now)

		if err != nil {
			log.Printf("INSERTエラー: %v", err)
		}
	}
	return nil
}
