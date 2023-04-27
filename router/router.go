package router

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"tsp_gin/conf"
	"tsp_gin/database"

	"github.com/gin-gonic/gin"
)

func TspRouter() http.Handler {
	engine := gin.Default()

	// log
	logger()

	//api
	userAPI(engine)
	imgAPI(engine)

	return engine
}

func logger() {
	dt := time.Now().Format("2006-01-02")
	f, _ := os.Create("./log/" + dt + ".log")
	gin.DefaultWriter = io.MultiWriter(f)
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

func imgAPI(engine *gin.Engine) {
	engine.StaticFS("/img", http.Dir("./tmp"))

	imgPath(engine, "gas")
	imgPath(engine, "greed")
	imgPath(engine, "dijkstra")
}

func imgPath(engine *gin.Engine, name string) {
	engine.Any(fmt.Sprintf("/api/v1/img/%s", name), func(context *gin.Context) {
		switch context.Request.Method {
		case http.MethodGet:
			context.JSON(http.StatusOK, gin.H{
				"img":  "gas.png",
				"path": fmt.Sprintf("http://%s:%s/img/tsp_%s.png", conf.SERVER_IP, conf.SERVER_PORT, name),
			})
		}
	})
}
