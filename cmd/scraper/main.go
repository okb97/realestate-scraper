package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/okb97/realestate-scraper/internal/db"
	"github.com/okb97/realestate-scraper/internal/scraper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".envの読み込みに失敗しました: %v", err)
	}

	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("DB接続に失敗しました: %v", err)
	}
	defer conn.Close()

	scraper.RunUsedCondoScraper(conn)
}
