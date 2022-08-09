package tests

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"moments/db"
	"testing"
)

func loadDbEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return err
}

func TestDatabaseConnection(t *testing.T) {
	err := loadDbEnv()
	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}
	dc, cancel := db.NewTestDB()
	if dc.Client == nil || dc.Ctx == nil || cancel == nil {
		t.Log("failed to connect to database")
		return
	}
}
