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
		technology.GET("/technology", scraper.GetAllArticles())
		technology.POST("/technology/article", scraper.GetArticle())
	}
}
