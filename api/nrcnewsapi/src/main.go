package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"nrcnewsapi/api/nrcnewsapi/src/docs"
	"nrcnewsapi/api/nrcnewsapi/src/route"
)

// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email williamhall2894@gmail.com
// @license.name MIT
// @license.url https://github.com/ciCciC/nrcnewsapi
func main() {
	initApiDoc()

	r := route.SetupRouter()

	initBasePath(r)
	initSwaggerPath(r)

	r.Run(":5011")
}

func initApiDoc() {
	docs.SwaggerInfo.Title = "NRC Scraper API"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Description = "Scrapes NRC news articles"
	docs.SwaggerInfo.Version = "1.0"
}

func initBasePath(r *gin.Engine) {
	r.LoadHTMLGlob("assets/*")
	r.GET("/",
		func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
			c.HTML(http.StatusOK, "index.html", gin.H{
				"": "",
			})
		})
}

func initSwaggerPath(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}
