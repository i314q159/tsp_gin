package router

import (
	"net/http"
	"tsp_gin/database"
	"tsp_gin/hook"

	"github.com/gin-gonic/gin"
)

func TspRouter() http.Handler {
	engine := gin.Default()

	// log
	hook.Logger()

	// web page
	webPage(engine)

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

func webPage(engine *gin.Engine) {
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
}
