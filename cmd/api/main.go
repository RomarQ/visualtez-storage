package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	DB "github.com/romarq/visualtez-storage/internal/data"
	API "github.com/romarq/visualtez-storage/internal/data/api"
	LOG "github.com/romarq/visualtez-storage/internal/logger"

	_ "github.com/romarq/visualtez-storage/docs"
)

// InitializeAPI - Initialize REST API
// @title Visualtez Storage API
// @version 1.0
// @description API documentation
// @BasePath /
func main() {
	configuration := GetConfig()
	LOG.SetupLogger(configuration.Log.Location, configuration.Log.Level)

	LOG.Info("Initializing Storage API...")

	e := echo.New()

	e.Use(middleware.CORS())

	client, err := DB.New(configuration.DB.URL)
	if err != nil {
		LOG.Fatal("Could not connect to database: %v", err)
	}
	defer client.Disconnect(context.TODO())

	database := client.Database("visualtez")
	sharingsAPI := API.InitSharingsAPI(database)

	// API Documentation
	e.GET("/doc/*", echoSwagger.WrapHandler)

	// API Endpoints
	e.GET("/sharings/:hash", sharingsAPI.GetSharing)
	e.POST("/sharings", sharingsAPI.InsertSharing)

	// Start REST API Service
	go func() {
		if err := e.Start(":" + GetConfig().Port); err != nil && err != http.ErrServerClosed {
			LOG.Fatal("Shutting down REST API service: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Using a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// Wait for the signal
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		LOG.Fatal("Error during shutdown: %v", err)
	}
}
