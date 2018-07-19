package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rH Router) health(c *gin.Context) {
	if err := rH.handler.Health(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
