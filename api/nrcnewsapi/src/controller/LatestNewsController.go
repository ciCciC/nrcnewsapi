package controller

import (
	"github.com/gin-gonic/gin"
	"nrcnewsapi/api/nrcnewsapi/src/scraper"
)

type LatestNewsController struct {
	BaseController
}

func (t LatestNewsController) InitRoute(r *gin.Engine) {
	scraper := scraper.Scraper{Endpoint: LATEST_NEWS, Topic: "nieuws"}
	latestNews := r.Group("/")
	{
		latestNews.GET("news", scraper.GetAllArticles())
	}
}
