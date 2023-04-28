package api

import (
	"fmt"
	"net/http"
	"tsp_gin/conf"

	"github.com/gin-gonic/gin"
)

func ImgAPI(engine *gin.Engine) {
	engine.StaticFS("/img", http.Dir("./tmp"))

	engine.Any(fmt.Sprintf("/api/%s/img/", conf.API_VERSION), func(context *gin.Context) {
		switch context.Request.Method {
		case http.MethodGet:
			context.JSON(http.StatusOK, gin.H{
				"img":  "",
				"path": "",
			})
		}
	})

	imgPath(engine, "gas")
	imgPath(engine, "greedy")
	imgPath(engine, "dijkstra")
}

func imgPath(engine *gin.Engine, imgName string) {
	engine.Any(fmt.Sprintf("/api/%s/img/%s", conf.API_VERSION, imgName), func(context *gin.Context) {
		switch context.Request.Method {
		case http.MethodGet:
			Algorithm(imgName)
			context.JSON(http.StatusOK, gin.H{
				"img":  "gas.png",
				"path": fmt.Sprintf("http://%s:%s/img/tsp_%s.png", conf.SERVER_IP, conf.SERVER_PORT, imgName),
			})
		}
	})
}
