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
		//technology.GET("/technology", scraper.GetAllArticles())
		technology.GET("/technology", t.allTechnologyArticles(scraper))
		technology.POST("/technology/article", t.technologyArticle(scraper))
		//technology.POST("/technology/article", scraper.GetArticle())
	}
}

// allTechnologyArticles godoc
// @Summary Retrieves all Technology articles
// @Produce json
// @Success 200 {array} model.ArticleItem
// @Router /category/technology [get]
func (t TechnologyController) allTechnologyArticles(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetAllArticles()
}

// technologyArticle godoc
// @Summary Retrieves Technology article
// @Produce json
// @Param articleItem path ArticleItem true "Article Item payload"
// @Success 200 {object} model.Article
// @Router /category/technology/article [post]
func (t TechnologyController) technologyArticle(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetArticle()
}
