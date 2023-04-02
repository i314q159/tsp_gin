package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./static/html/*")
	r.Static("/dwz", "./statics")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// api
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

	//port
	r.Run(":8080")
}
