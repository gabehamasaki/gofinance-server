package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func Init() {
	app = gin.Default()

	routes()

	app.Run(os.Getenv("PORT"))
}
