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
	}
}

// allArticles godoc
// @Summary Retrieves all articles
// @Produce json
// @Success 200 {array} model.ArticleItem
// @Router /category/:category [get]
func (t AnyController) allArticles(scrape scraper.Scraper) gin.HandlerFunc {
	return scrape.GetAllArticles()
}
