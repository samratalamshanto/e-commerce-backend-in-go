package main

import (
	"ecom-backend/internal/common"
	"ecom-backend/internal/router"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	common.InitDB()
	r := gin.Default()

	router.LoadRouter(r)

	fmt.Println("Server running at :9090")
	r.Run(":9090")
}
