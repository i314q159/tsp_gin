package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TspRouter() http.Handler {
	engine := gin.New()
	engine.Use(cors.Default())

	//api
	UserAPI(engine)
	ImgAPI(engine)

	return engine
}
