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
	physics := r.Group("/" + CATEGORY)
	{
		physics.GET("/physics", scraper.GetAllArticles())
		physics.POST("/physics/article", scraper.GetArticle())
	}
}
