package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nrcnewsapi/api/nrcnewsapi/src/model"
)

type CategoryController struct {
	BaseController
}

func (c CategoryController) InitRoute(r *gin.Engine) {
	r.GET("/categories", c.GetCategories)
}

func (c CategoryController) GetCategories(context *gin.Context) {
	context.JSON(http.StatusOK,
		[...]model.Category{
			{Display: "latest news", Topic: "news"},
			{Display: "games", Topic: "games"},
			{Display: "technology", Topic: "technology"},
			{Display: "physics", Topic: "physics"},
			{Display: "astronomy", Topic: "astronomy"},
		})
}
