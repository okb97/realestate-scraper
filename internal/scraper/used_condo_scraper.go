package scraper

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/okb97/realestate-scraper/internal/db"
	"github.com/okb97/realestate-scraper/internal/model"
	"github.com/okb97/realestate-scraper/internal/scraper/parse"
	"github.com/okb97/realestate-scraper/internal/transform"
)

func RunUsedCondoScraper(conn *sql.DB) {
	var detailCollectormodel model.DetailCollector

	mainCollector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"),
	)
	detailCollector := mainCollector.Clone()

	detailCollector.OnHTML("#mainContents", func(e *colly.HTMLElement) {
		tx, err := conn.Begin()
		if err != nil {
			log.Printf("トランザクション開始失敗: %v", err)
			return
		}
		detailCollectormodel = parse.ParseDetailCollector(e)
		usedCondoModel := transform.TransformUsedCondo(tx, detailCollectormodel)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		usedCondoID, updated, err := db.InsertUsedCondo(ctx, tx, usedCondoModel)
		if err != nil {
			tx.Rollback()
			log.Printf("InsertUsedCondo失敗:物件名=%s,URL:%s,Error:%s", usedCondoModel.UsedCondoName, usedCondoModel.Url, err)
			return
		}
		if usedCondoID != 0 {
			err = transform.TransformStationAndBusStop(tx, detailCollectormodel.Transportation, usedCondoID)
			if err != nil {
				tx.Rollback()
				log.Printf("TransformStationAndBusStop失敗:物件名=%s,URL:%s,ERROR:%s", usedCondoModel.UsedCondoName, usedCondoModel.Url, err)
				return
			}
		}
		if err := tx.Commit(); err != nil {
			log.Printf("トランザクションコミット失敗: %v", err)
			return
		}
		if updated {
			log.Printf("データの挿入が完了しました。usedCondoName:%s,URL:%s", usedCondoModel.UsedCondoName, usedCondoModel.Url)
		}

	})

	mainCollector.OnHTML(".property_unit-title a", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		detailCollector.Visit(link)
	})

	// ページネーション
	mainCollector.OnHTML(".pagination a", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "次へ") {
			nextPage := e.Request.AbsoluteURL(e.Attr("href"))
			log.Println("次ページへ:", nextPage)
			e.Request.Visit(nextPage)
		}
	})

	startURL := "https://suumo.jp/jj/bukken/ichiran/JJ010FJ001/?ar=030&bs=011&ta=14&jspIdFlg=patternShikugun&sc=14114&kb=1&kt=9999999&mb=0&mt=9999999&ekTjCd=&ekTjNm=&tj=0&cnb=0&cn=9999999&srch_navi=1"
	if err := mainCollector.Visit(startURL); err != nil {
		log.Fatal(err)
	}

}
