package conf

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type conf struct {
}

func LoadDbEnv() error {
	err := godotenv.Load("conf/db.env")
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return err
}
