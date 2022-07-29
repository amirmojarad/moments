package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"moments/db"
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
	// dc stands for DatabaseConnection
	dc, cancel := db.New()

	defer dc.Client.Close()
	defer cancel()

}
