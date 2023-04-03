package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterB() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": "Welcome server 02",
		})
	})
	return r
}
