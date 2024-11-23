package main

import (
	"neuro-most/media-service/config"
	"neuro-most/media-service/internal/infra"
)

func main() {
	cfg, err := config.NewLoadConfig()
	if err != nil {
		panic(err)
	}
	infra.Config(cfg).Database().Serve().Start()
}
