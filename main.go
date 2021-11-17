package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gin-weight-tracker/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Static("/assets", "./app/assets")
	r.LoadHTMLGlob("app/views/*")

	r.Use(cors.Default())
	routes.MakeAppRoutes(r)

	r.Run(":" + os.Getenv("APP_PORT"))
}
