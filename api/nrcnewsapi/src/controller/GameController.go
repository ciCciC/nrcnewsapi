package controller

import (
	"github.com/gin-gonic/gin"
	"nrcnewsapi/api/nrcnewsapi/src/scraper"
)

type GameController struct {
	BaseController
}

func (t GameController) InitRoute(r *gin.Engine) {
	scraper := scraper.Scraper{Endpoint: GAMES, Topic: "game"}
	game := r.Group("/" + CATEGORY)
	{
		game.GET("/game", t.allArticles(scraper))
		game.POST("/game/article", t.article(scraper))
	}
}

// allArticles godoc
// @Summary Retrieves all articles
// @Produce json
// @Success 200 {array} model.ArticleItem
// @Router /category/games [get]
func (t GameController) allArticles(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetAllArticles()
}

// article godoc
// @Summary Retrieves article
// @Description Get a article with article item payload
// @Accept  json
// @Produce json
// @Param articleitem body model.ArticleItem true "Get article"
// @Success 200 {object} model.Article
// @Router /category/games/article [post]
func (t GameController) article(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetArticle()
}
