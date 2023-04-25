package router

import (
	"io"
	"net/http"
	"os"
	"time"
	"tsp_gin/database"

	"github.com/gin-gonic/gin"
)

func TspRouter() http.Handler {
	engine := gin.Default()

	//log
	dt := time.Now().Format("2006-01-02")
	f, _ := os.Create("./log/" + dt + ".log")
	gin.DefaultWriter = io.MultiWriter(f)

	engine.LoadHTMLGlob("./static/html/*")

	engine.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	engine.GET("/login", func(ccontext *gin.Context) {
		ccontext.HTML(http.StatusOK, "login.html", nil)
	})
	engine.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	//api
	userAPI(engine)

	return engine
}

func userAPI(engine *gin.Engine) {
	engine.Any("/api/v1/user", func(context *gin.Context) {
		switch context.Request.Method {
		case http.MethodPost:
			database.AddUser(context)
		case http.MethodPut:
			database.UpdateUser(context)
		}
	})
}
