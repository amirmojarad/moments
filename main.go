package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"moments/api"
)

func loadDbEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return err
}

func main() {
	loadDbEnv()
	engine := api.RunEngine()
	engine.Run(":8080")
}
