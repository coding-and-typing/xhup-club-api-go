package main

import (
	"github.com/coding-and-typing/xhup-club-api-go/internal/config"
	"github.com/coding-and-typing/xhup-club-api-go/models"
	"github.com/coding-and-typing/xhup-club-api-go/router"
)

func main() {
	configPath := "config.dev.yaml"
	config.Init(configPath)

	// TODO fix config.Conf.DB nil pointer here
	err := models.Database(&config.Conf.DB)
	if err != nil {
		panic(err)
	}

	engine := router.InitRouter()
	engine.Run(":8080")
}
