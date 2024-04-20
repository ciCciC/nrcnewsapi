package controller

import (
	"github.com/gin-gonic/gin"
	"nrcnewsapi/api/nrcnewsapi/src/scraper"
)

type AnyController struct {
	BaseController
}

func (t AnyController) InitRoute(r *gin.Engine) {
	scraper := scraper.Scraper{Endpoint: "/categorie", Topic: "", State: "/full"}
	anyCategory := r.Group("/" + CATEGORY)
	{
		anyCategory.GET("/all/:name", t.allArticles(scraper))
		anyCategory.POST("/all/article", t.article(scraper))
		anyCategory.POST("/all/article/fallback", t.fallback(scraper))
	}
}

// allArticles godoc
// @Summary Retrieves all articles
// @Produce json
// @Success 200 {array} model.ArticleItem
// @Router /all/:name [get]
func (t AnyController) allArticles(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetAllArticles()
}

// article godoc
// @Summary Retrieves article
// @Description Get an article with article item payload
// @Accept  json
// @Produce json
// @Param articleitem body model.ArticleItem true "Get article"
// @Success 200 {object} model.Article
// @Router /all/article [post]
func (t AnyController) article(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetArticle()
}

// article godoc
// @Summary Retrieves article using a fallback (older articles)
// @Description Get an article with article item payload
// @Accept  json
// @Produce json
// @Param articleitem body model.ArticleItem true "Get article"
// @Success 200 {object} model.Article
// @Router /all/article/fallback [post]
func (t AnyController) fallback(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetArticleFallback()
}
