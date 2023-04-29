package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TspRouter() http.Handler {
	engine := gin.Default()

	//api
	UserAPI(engine)
	ImgAPI(engine)

	return engine
}
