package hook

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() {
	dt := time.Now().Format("2006-01-02")
	f, _ := os.Create("./log/" + dt + ".log")
	gin.DefaultWriter = io.MultiWriter(f)
}
