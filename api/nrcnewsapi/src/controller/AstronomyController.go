package controller

import (
	"github.com/gin-gonic/gin"
	"nrcnewsapi/api/nrcnewsapi/src/scraper"
)

type AstronomyController struct {
	BaseController
}

func (t AstronomyController) InitRoute(r *gin.Engine) {
	scraper := scraper.Scraper{Endpoint: ASTRONOMY, Topic: "astronomie"}
	game := r.Group("/" + CATEGORY)
	{
		game.GET("/astronomy", t.allArticles(scraper))
		game.POST("/astronomy/article", t.article(scraper))
	}
}

// allArticles godoc
// @Summary Retrieves all articles
// @Produce json
// @Success 200 {array} model.ArticleItem
// @Router /category/astronomy [get]
func (t AstronomyController) allArticles(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetAllArticles()
}

// article godoc
// @Summary Retrieves article
// @Description Get a article with article item payload
// @Accept  json
// @Produce json
// @Param articleitem body model.ArticleItem true "Get article"
// @Success 200 {object} model.Article
// @Router /category/astronomy/article [post]
func (t AstronomyController) article(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetArticle()
}
