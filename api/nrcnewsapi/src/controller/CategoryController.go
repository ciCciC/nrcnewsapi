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
	categories := r.Group("/categories")
	{
		categories.GET("/", t.GetCategories)
		categories.GET("/nl", t.GetCategoriesNL)
	}
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
			{Display: "games", Topic: "game"},
			{Display: "technology", Topic: "technology"},
			{Display: "physics", Topic: "physics"},
			{Display: "astronomy", Topic: "astronomy"},
		})
}

// GetCategoriesNL godoc
// @Summary Retrieves category list in Dutch
// @Description Get category list in Dutch
// @Produce json
// @Success 200 {array} model.Category
// @Router /categories/nl [get]
func (t CategoryController) GetCategoriesNL(context *gin.Context) {
	context.JSON(http.StatusOK,
		[...]model.Category{
			{Display: ".", Topic: "nieuws"},
			{Display: ".", Topic: "games"},
			{Display: ".", Topic: "technologie"},
			{Display: ".", Topic: "natuurkunde"},
			{Display: ".", Topic: "astronomie"},
			{Display: ".", Topic: "geologie"},
			{Display: ".", Topic: "geschiedenis"},
			{Display: ".", Topic: "archeologie"},
			{Display: ".", Topic: "biologie"},
			{Display: ".", Topic: "cultuur"},
			{Display: ".", Topic: "binnenland"},
			{Display: ".", Topic: "buitenland"},
			{Display: ".", Topic: "economie"},
		})
}
