package controller

import (
	"github.com/gin-gonic/gin"
	"nrcnewsapi/api/nrcnewsapi/src/scraper"
)

type BaseController interface {
	InitRoute(r *gin.Engine)
	allArticles(scrape scraper.Scraper) gin.HandlerFunc
	article(scrape scraper.Scraper) gin.HandlerFunc
}

const (
	CATEGORY    = "category"
	GAMES       = "/categorie/games"
	TECHNOLOGY  = "/categorie/technologie"
	PHYSICS     = "/categorie/natuurkunde"
	ASTRONOMY   = "/categorie/astronomie"
	LATEST_NEWS = "/nieuws"
)
