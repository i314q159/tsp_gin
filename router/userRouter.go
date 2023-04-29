package router

import (
	"fmt"
	"net/http"
	"tsp_gin/conf"
	"tsp_gin/database"

	"github.com/gin-gonic/gin"
)

func UserAPI(engine *gin.Engine) {
	engine.Any(fmt.Sprintf("/api/%s/user/login", conf.API_VERSION), func(context *gin.Context) {
		switch context.Request.Method {
		case http.MethodPost:
			database.UserLogin(context)
		}
	})

	engine.Any(fmt.Sprintf("/api/%s/user/register", conf.API_VERSION), func(context *gin.Context) {
		switch context.Request.Method {
		case http.MethodPost:
			database.UserRegister(context)
		}
	})
}
