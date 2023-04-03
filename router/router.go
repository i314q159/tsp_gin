package router

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func TspRouter() http.Handler {
	r := gin.Default()

	//log
	dt := time.Now().Format("2006-01-02")
	f, _ := os.Create("./log/" + dt + ".log")
	gin.DefaultWriter = io.MultiWriter(f)

	r.LoadHTMLGlob("./static/html/*")
	r.StaticFS("/dwz", http.Dir("./statics"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	//api
	r.GET("/api/v1/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})
	r.POST("/api/v1/user", func(c *gin.Context) {
		name := c.Query("name")
		email := c.Query("email")

		c.JSON(http.StatusOK, gin.H{
			"name":  name,
			"email": email,
		})
	})
	r.PUT("/api/v1/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})
	r.DELETE("/api/v1/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})

	return r
}
