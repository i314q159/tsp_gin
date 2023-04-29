package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebAPI(engine *gin.Engine) {
	engine.LoadHTMLGlob("./web/html/*")
	engine.StaticFS("./web", http.Dir("./web"))

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	webAPI(engine, "login")
	webAPI(engine, "gas")

	engine.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
}

func webAPI(engine *gin.Engine, html string) {
	engine.GET(fmt.Sprintf("/%s", html), func(c *gin.Context) {
		c.HTML(http.StatusOK, fmt.Sprintf("%s.html", html), nil)
	})
}
