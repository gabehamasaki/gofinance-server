package main

import (
	"fmt"

	"github.com/gabehamasaki/finance-server/internal/router"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading environment: %v", err)
		return
	}

	router.Init()
}
