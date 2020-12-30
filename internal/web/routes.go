package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s Server) statusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
