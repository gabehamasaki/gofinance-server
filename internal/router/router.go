package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func Init() {
	app = gin.Default()

	routes()

	app.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
