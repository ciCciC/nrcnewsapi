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
	latestNews := r.Group("/" + CATEGORY)
	{
		latestNews.GET("/news", t.allArticles(scraper))
		latestNews.POST("/news/article", t.article(scraper))
	}
}

// allArticles godoc
// @Summary Retrieves all articles
// @Produce json
// @Success 200 {array} model.ArticleItem
// @Router /category/news [get]
func (t LatestNewsController) allArticles(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetAllArticles()
}

// article godoc
// @Summary Retrieves article
// @Description Get a article with article item payload
// @Accept  json
// @Produce json
// @Param articleitem body model.ArticleItem true "Get article"
// @Success 200 {object} model.Article
// @Router /category/news/article [post]
func (t LatestNewsController) article(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetArticle()
}
