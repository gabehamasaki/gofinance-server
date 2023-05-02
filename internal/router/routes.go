package router

import "github.com/gabehamasaki/finance-server/internal/handlers"

func routes() {

	handlers.Init()

	v1 := app.Group("/api/v1")
	{
		v1.POST("/account", handlers.CreateUser)
	}
}
