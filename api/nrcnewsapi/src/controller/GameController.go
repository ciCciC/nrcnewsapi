package controller

import (
	"github.com/gin-gonic/gin"
	"nrcnewsapi/api/nrcnewsapi/src/scraper"
)

type GameController struct {
	BaseController
}

func (t GameController) InitRoute(r *gin.Engine) {
	scraper := scraper.Scraper{Endpoint: "games"}
	game := r.Group("/" + CATEGORY)
	{
		game.GET("/games", scraper.GetAllArticles())
		game.GET("/games/article", scraper.GetArticle())
	}
}
