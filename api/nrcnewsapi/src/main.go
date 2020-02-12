package main

import (
	"github.com/gin-gonic/gin"
	"nrcnewsapi/api/nrcnewsapi/src/route"
)

func main() {
	r := route.SetupRouter()
	// text = Lets Scrape
	welcomeTxt := "  _          _             _____                                _ \r\n | |        | |           / ____|                              | |\r\n | |     ___| |_ ___     | (___   ___ _ __ __ _ _ __   ___     | |\r\n | |    / _ \\ __/ __|     \\___ \\ / __| '__/ _` | '_ \\ / _ \\    | |\r\n | |___|  __/ |_\\__ \\     ____) | (__| | | (_| | |_) |  __/    |_|\r\n |______\\___|\\__|___/    |_____/ \\___|_|  \\__,_| .__/ \\___|    (_)\r\n                                               | |                \r\n                                               |_|                "

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"": welcomeTxt,
		})
	})
	r.Run()
}
