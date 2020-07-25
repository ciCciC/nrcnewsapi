package controller

import (
	"github.com/gin-gonic/gin"
	"nrcnewsapi/api/nrcnewsapi/src/scraper"
)

type AstronomyController struct {
	BaseController
}

func (c AstronomyController) InitRoute(r *gin.Engine) {
	scraper := scraper.Scraper{Endpoint: "astronomie"}
	game := r.Group("/" + CATEGORY)
	{
		game.GET("/astronomy", scraper.GetAllArticles())
		game.POST("/astronomy/article", scraper.GetArticle())
	}
}
