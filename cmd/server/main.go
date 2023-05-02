package main

import (
	"fmt"

	"github.com/gabehamasaki/finance-server/internal/database"
	"github.com/gabehamasaki/finance-server/internal/router"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading environment: %v", err)
		return
	}

	if err := database.Init(); err != nil {
		fmt.Printf("error connect to database: %v", err)
		return
	}

	router.Init()
}
