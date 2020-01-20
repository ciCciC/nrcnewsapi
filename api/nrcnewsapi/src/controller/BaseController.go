package controller

import "github.com/gin-gonic/gin"

type BaseController interface {
	InitRoute(r *gin.Engine)
}