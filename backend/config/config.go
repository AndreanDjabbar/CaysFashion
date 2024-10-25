package config

import (
	"text/template"
	"github.com/AndreanDjabbar/CaysFashion/backend/database"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/routes"
	"github.com/AndreanDjabbar/CaysFashion/backend/migrations"
	"github.com/AndreanDjabbar/CaysFashion/backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var log = logger.SetUpLogger()

func DBInit() {
	db := database.GetDB()
	migrations.Migrate(db)
}

func RouteInit() *gin.Engine {
	route := gin.Default()
	route.SetFuncMap(template.FuncMap{
		//... 
	})
	route.MaxMultipartMemory = 8 << 20
	routes.SetUpMainRoutes(route)
	return route
}

func EnvInit() {
	err := godotenv.Load()
	if err != nil {
		log.Error(
			"Failed to load .env file",
			"error", err,
		)
		return
	}
}