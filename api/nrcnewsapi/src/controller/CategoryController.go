package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nrcnewsapi/api/nrcnewsapi/src/model"
)

type CategoryController struct {
	BaseController
}

func (t CategoryController) InitRoute(r *gin.Engine) {
	r.GET("/categories", t.GetCategories)
}

// GetCategories godoc
// @Summary Retrieves category list
// @Description Get category list
// @Produce json
// @Success 200 {array} model.Category
// @Router /categories [get]
func (t CategoryController) GetCategories(context *gin.Context) {
	context.JSON(http.StatusOK,
		[...]model.Category{
			{Display: "latest news", Topic: "news"},
			{Display: "games", Topic: "games"},
			{Display: "technology", Topic: "technology"},
			{Display: "physics", Topic: "physics"},
			{Display: "astronomy", Topic: "astronomy"},
		})
}
