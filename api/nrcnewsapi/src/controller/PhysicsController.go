package controller

import (
	"github.com/gin-gonic/gin"
	"nrcnewsapi/api/nrcnewsapi/src/scraper"
)

type PhysicsController struct {
	BaseController
}

func (t PhysicsController) InitRoute(r *gin.Engine) {
	scraper := scraper.Scraper{Endpoint: "natuurkunde"}
	physics := r.Group("/category")
	{
		physics.GET("/physics", scraper.GetAllArticles())
		physics.GET("/physics/article", scraper.GetArticle())
	}
}