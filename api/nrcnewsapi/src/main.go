package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"nrcnewsapi/api/nrcnewsapi/src/docs"
	"nrcnewsapi/api/nrcnewsapi/src/route"
)

// @title NRC News Swagger API
// @version 1.0
// @description Swagger API for NRC news API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email williamhall2894@gmail.com

// @license.name MIT
// @license.url https://github.com/ciCciC/nrcnewsapi

// @BasePath /api/v1
func main() {
	docs.SwaggerInfo.Title = "NRC news API"

	r := route.SetupRouter()
	welcomeTxt := "  _          _             _____                                _ \r\n | |        | |           / ____|                              | |\r\n | |     ___| |_ ___     | (___   ___ _ __ __ _ _ __   ___     | |\r\n | |    / _ \\ __/ __|     \\___ \\ / __| '__/ _` | '_ \\ / _ \\    | |\r\n | |___|  __/ |_\\__ \\     ____) | (__| | | (_| | |_) |  __/    |_|\r\n |______\\___|\\__|___/    |_____/ \\___|_|  \\__,_| .__/ \\___|    (_)\r\n                                               | |                \r\n                                               |_|                "

	r.GET("/",
		func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
			c.JSON(http.StatusOK, gin.H{
				"": welcomeTxt,
			})
		})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":5011")
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}
