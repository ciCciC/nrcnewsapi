package route

import (
	"github.com/gin-gonic/gin"
	"nrcnewsapi/api/nrcnewsapi/src/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	cArr := []controller.BaseController{
		new(controller.TechnologyController),
		new(controller.GameController),
		new(controller.PhysicsController),
		new(controller.CategoryController),
	}

	for _, baseController := range cArr {
		baseController.InitRoute(r)
	}

	return r
}
