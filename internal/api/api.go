package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okb97/realestate-scraper/internal/db"
)

func GetUsedCondos(c *gin.Context) {
	//log.Println("ここまでは正常")
	condos, err := db.GetAllUsedCondos()
	if err != nil {
		log.Printf("DB取得エラー: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB getch error"})
		return
	}
	c.JSON(http.StatusOK, condos)
}
