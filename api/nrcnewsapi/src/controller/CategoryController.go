package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
	BaseController
}

func (c CategoryController) InitRoute(r *gin.Engine) {
	r.GET("/categories", c.GetCategories)
}

func (c CategoryController) GetCategories(context *gin.Context) {
	context.JSON(http.StatusOK,
		[...]string{"games", "technology", "physics", "astronomy"})
}
