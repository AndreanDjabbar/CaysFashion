package main

import (
	"fmt"
	"os"

	"github.com/AndreanDjabbar/CaysFashion/backend/config"
	"github.com/AndreanDjabbar/CaysFashion/backend/pkg/logger"
)

func main() {
	config.EnvInit()
	config.DBInit()
	routes := config.RouteInit()
	logger := logger.SetUpLogger()

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Starting server", "host", host, "port", port)
	if err := routes.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		logger.Error(
			"Failed to start server",
			"error", err,
		)
		panic(err)
    }
}