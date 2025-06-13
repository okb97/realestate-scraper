package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/okb97/realestate-scraper/config"
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

	// コマンドライン引数をチェック
	urls, areaName, err := config.ParseCommandLineArgs()
	if err != nil {
		log.Fatalf("コマンドライン引数エラー: %v", err)
	}
	
	// コマンドライン引数が指定されていない場合はインタラクティブモード
	if urls == nil {
		selector := config.NewAreaSelector()
		
		// 地域概要を表示
		selector.GetAreaSummary()
		
		// 地域選択
		urls, areaName, err = selector.SelectArea()
		if err != nil {
			log.Fatalf("地域選択エラー: %v", err)
		}
		
		// 終了が選択された場合
		if areaName == "exit" {
			fmt.Println("スクレイピングを終了します。")
			os.Exit(0)
		}
	}
	
	fmt.Printf("\n=== %s のスクレイピングを開始します ===\n", areaName)
	fmt.Printf("対象URL数: %d件\n\n", len(urls))
	
	// 選択された地域のスクレイピングを実行
	scraper.RunUsedCondoScraperWithURLs(conn, urls)
	
	fmt.Printf("\n=== %s のスクレイピングが完了しました ===\n", areaName)
}
