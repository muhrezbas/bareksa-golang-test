package main

import (
	"bareksa-api/config"
	"bareksa-api/router"
	"os"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	cfg config.Interface
	err error
)

func init() {
	_, err := os.Stat(".env")
	if !os.IsNotExist(err) {
		godotenv.Load(".env")
	}

	cfg = config.NewConfig()
}

func main() {

	routers := gin.Default()
	routers.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "UPDATE"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routers.Use(gzip.Gzip(gzip.DefaultCompression))
	app := router.Context{
		R: routers,
	}
	app.LoadRoutes()
	app.R.Run(":" + os.Getenv("APP_PORT"))
}
