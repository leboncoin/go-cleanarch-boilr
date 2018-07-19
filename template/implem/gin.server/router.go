package server

import (
	"github.com/gin-gonic/gin"
	
	"{{GitServer}}/{{Organization}}/{{Name}}/uc"
)

type Router struct {
	handler uc.Handler
}

func NewRouter(handler uc.Handler) Router {
	return Router{
		handler: handler,
	}
}

func (rH Router) SetRoutes(router *gin.Engine) {
	router.GET("/health", rH.health)
}