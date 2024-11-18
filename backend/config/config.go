package config

import (
	"text/template"
	"time"
	"github.com/AndreanDjabbar/CaysFashion/backend/database/mysql_service"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/routes"
	"github.com/AndreanDjabbar/CaysFashion/backend/migrations"
	"github.com/AndreanDjabbar/CaysFashion/backend/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var log = logger.SetUpLogger()

func EnvInit() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Failed to load .env file", "error", err)
		return
	}
}

func DBInit() {
	db := mysql_service.GetDB()
	migrations.Migrate(db)
}

func RouteInit() *gin.Engine {
	route := gin.Default()
	route.SetFuncMap(template.FuncMap{
		//... 
	})
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"POST", "GET", "OPTIONS"}, 
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
	}))
	route.MaxMultipartMemory = 8 << 20
	routes.SetUpMainRoutes(route)
	return route
}