package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebAPI(engine *gin.Engine, path string) {
	engine.LoadHTMLGlob(path + "/*.html")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	engine.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	engine.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
}
