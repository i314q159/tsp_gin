package router

import (
	"io"
	"net/http"
	"os"
	"time"
	"tsp_gin/api"

	"github.com/gin-gonic/gin"
)

func TspRouter() http.Handler {
	engine := gin.Default()

	// log
	logger()

	//api
	api.UserAPI(engine)
	api.ImgAPI(engine)

	return engine
}

func WebRouter() http.Handler {
	engine := gin.Default()
	api.WebAPI(engine)
	return engine
}

func logger() {
	dt := time.Now().Format("2006-01-02")
	f, _ := os.Create("./log/" + dt + ".log")
	gin.DefaultWriter = io.MultiWriter(f)
}
