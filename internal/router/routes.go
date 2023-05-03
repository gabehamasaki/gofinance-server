package router

import (
	"github.com/gabehamasaki/finance-server/internal/handlers"
	"github.com/gabehamasaki/finance-server/internal/middlewares"
)

func routes() {

	handlers.Init()

	v1 := app.Group("/api/v1")
	{

		v1.POST("/auth", handlers.Auth)
		v1.GET("/auth/validate", handlers.ValidateToken)

		v1.POST("/account", handlers.CreateUser)

		private := v1.Group("", middlewares.AuthMiddleware)
		{
			private.GET("/account", handlers.ShowUser)
			private.GET("/account/transactions", handlers.ShowAccountTransactions)
			private.DELETE("/account", handlers.DeleteAccount)
			private.PATCH("/account", handlers.UpdateAccount)

			private.POST("/transaction", handlers.CreateTransaction)
			private.GET("/transaction/:id", handlers.ShowTransaction)
			private.DELETE("/transaction/:id", handlers.DeleteTransaction)
			private.PATCH("/transaction/:id", handlers.UpdateTransactio)
		}
	}
}
