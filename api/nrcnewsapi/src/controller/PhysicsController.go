package controller

import (
	"github.com/gin-gonic/gin"
	"nrcnewsapi/api/nrcnewsapi/src/scraper"
)

type PhysicsController struct {
	BaseController
}

func (t PhysicsController) InitRoute(r *gin.Engine) {
	scraper := scraper.Scraper{Endpoint: PHYSICS, Topic: "natuurkunde"}
	physics := r.Group("/" + CATEGORY)
	{
		physics.GET("/physics", t.allArticles(scraper))
		physics.POST("/physics/article", t.article(scraper))
	}
}

// allArticles godoc
// @Summary Retrieves all articles
// @Produce json
// @Success 200 {array} model.ArticleItem
// @Router /category/physics [get]
func (t PhysicsController) allArticles(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetAllArticles()
}

// article godoc
// @Summary Retrieves article
// @Description Get a article with article item payload
// @Accept  json
// @Produce json
// @Param articleitem body model.ArticleItem true "Get article"
// @Success 200 {object} model.Article
// @Router /category/physics/article [post]
func (t PhysicsController) article(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetArticle()
}
