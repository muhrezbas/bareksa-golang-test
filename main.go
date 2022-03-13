package main

import (
	"bareksa-api/config"
	"bareksa-api/pkg/mysql"
	"bareksa-api/router"
	"os"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// mysql ...
	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *gorm.DB
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

	NewMysql := mysql.Config{
		Host: cfg.GetString("mysql", "host"),
		Port: cfg.GetInt("mysql", "port"),
		User: cfg.GetString("mysql", "user"),
		Pass: cfg.GetString("mysql", "pass"),
		DB:   cfg.GetString("mysql", "db"),
	}

	db, err = NewMysql.Connect()
	if err != nil {
		panic(err)
	}

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
