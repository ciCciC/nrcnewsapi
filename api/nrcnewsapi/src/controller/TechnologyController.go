package controller

import (
	"github.com/gin-gonic/gin"
	"nrcnewsapi/api/nrcnewsapi/src/scraper"
)

type TechnologyController struct {
	BaseController
}

func (t TechnologyController) InitRoute(r *gin.Engine) {
	scraper := scraper.Scraper{Endpoint: TECHNOLOGY, Topic: "technologie"}
	technology := r.Group("/" + CATEGORY)
	{
		technology.GET("/technology", t.allArticles(scraper))
		technology.POST("/technology/article", t.article(scraper))
	}
}

// allArticles godoc
// @Summary Retrieves all articles
// @Produce json
// @Success 200 {array} model.ArticleItem
// @Router /category/technology [get]
func (t TechnologyController) allArticles(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetAllArticles()
}

// article godoc
// @Summary Retrieves article
// @Description Get a article with article item payload
// @Accept  json
// @Produce json
// @Param articleitem body model.ArticleItem true "Get article"
// @Success 200 {object} model.Article
// @Router /category/technology/article [post]
func (t TechnologyController) article(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetArticle()
}
