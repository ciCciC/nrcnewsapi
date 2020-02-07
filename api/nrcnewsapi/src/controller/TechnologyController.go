package controller

import (
	"github.com/gin-gonic/gin"
	"nrcnewsapi/api/nrcnewsapi/src/scraper"
)

type TechnologyController struct {
	BaseController
}

func (t TechnologyController) InitRoute(r *gin.Engine) {
	scraper := scraper.Scraper{Endpoint: "technologie"}
	technology := r.Group("/category")
	{
		technology.GET("/technology", scraper.GetAllArticles())
		technology.GET("/technology/article", scraper.GetArticle())
	}
}
