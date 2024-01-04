package router

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"tsp_gin/conf"

	"github.com/gin-gonic/gin"
)

func ImgAPI(engine *gin.Engine) {
	engine.StaticFS("/img", http.Dir("./tmp"))

	imgPath(engine, "gas")
	imgPath(engine, "greedy")
	imgPath(engine, "dijkstra")
}

func imgPath(engine *gin.Engine, imgName string) {
	engine.Any(fmt.Sprintf("/api/%s/img/%s", conf.API_VERSION, imgName), func(context *gin.Context) {
		switch context.Request.Method {
		case http.MethodGet:
			cp := context.QueryArray("cp")

			Algorithm(imgName, cp)
			context.JSON(http.StatusOK, gin.H{
				"img":               imgName,
				"path":              fmt.Sprintf("http://%s:%s/img/%s.png", conf.SERVER_IP, conf.SERVER_PORT, imgName),
				"coordinate_points": cp,
			})
		}
	})
}

func Algorithm(algorithm string, cp []string) {
	pyName := fmt.Sprintf("./lib/%s.py", algorithm)

	// ["1,2" "3,4"] => "1,2 3,4"
	args := strings.Replace(strings.Trim(fmt.Sprint(cp), "[]"), " ", " ", -1)

	cmd := exec.Command("python3", pyName, args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(algorithm)
	fmt.Println(string(out))
}
