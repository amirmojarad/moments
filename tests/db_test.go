package tests

import (
	"moments/conf"
	"moments/db"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	err := conf.LoadDbEnv()
	if err != nil {
		t.Log(err)
	}
	client, ctx, cancel := db.New()
	if client == nil || ctx == nil || cancel == nil {
		t.Log("failed to connect to database")
		return
	}
}
