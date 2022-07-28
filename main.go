package main

import (
	"moments/conf"
	"moments/db"
)

func main() {
	conf.LoadDbEnv()
	// dc stands for DatabaseConnection
	dc, cancel := db.New()

	defer dc.Client.Close()
	defer cancel()

}
