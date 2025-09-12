package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/config"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

func init() {
	num64, err := strconv.ParseUint(os.Getenv("LOG_LEVEL"), 10, 32)
	if err != nil {
		logrus.Error(err)
	}
	num32 := uint32(num64)
	logrus.SetLevel(logrus.Level(num32))
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}

func main() {
	db := config.NewDB()
	validator := config.NewValidator()
	app, err := config.NewFiber()
	cfg := config.NewConfig()
	if err != nil {
		logrus.Error(err)
	}

	config.Bootstrap(&config.BootstrapConfig{
		DB:        db,
		Validator: validator,
		App:       app,
		Config:    cfg,
	})

	webPort := os.Getenv("PORT")
	err = app.Listen(fmt.Sprintf(":%s", webPort))
	if err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}
