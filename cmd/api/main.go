package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/okb97/realestate-scraper/internal/api"
	"github.com/okb97/realestate-scraper/internal/db"
)

func main() {
	log.Println("ここまでは正常")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("`.env`の読み込みに失敗しました")
	}

	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	db.Conn = dbConn
	defer dbConn.Close()

	r := gin.Default()

	// CORS許可（シンプルに全許可）
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// ルート登録
	r.GET("/api/used-condos", api.GetUsedCondos)

	// サーバ起動
	r.Run(":8080")
}
