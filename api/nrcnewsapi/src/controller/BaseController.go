package controller

import (
	"github.com/gin-gonic/gin"
)

type BaseController interface {
	InitRoute(r *gin.Engine)
}

const (
	CATEGORY    = "category"
	GAMES       = "/categorie/games"
	TECHNOLOGY  = "/categorie/technologie"
	PHYSICS     = "/categorie/natuurkunde"
	ASTRONOMY   = "/categorie/astronomie"
	LATEST_NEWS = "/nieuws"
)
