package db

import (
	"context"
	"entgo.io/ent/entc/integration/ent"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func dsn() (string, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		err := errors.New("DB_Host is empty in env file")
		log.Error().Err(err).Msg("")
		return "", err
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		err := errors.New("DB_PASSWORD is empty in env file")
		log.Error().Err(err).Msg("")
		return "", err
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		err := errors.New("DB_USER is empty in env file")
		log.Error().Err(err).Msg("")
		return "", err
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		err := errors.New("DB_NAME is empty in env file")
		log.Error().Err(err).Msg("")
		return "", err
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		err := errors.New("DB_PORT is empty in env file")
		log.Error().Err(err).Msg("")
		return "", err
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		host, user, password, name, port), nil
}

type DatabaseConnection struct {
	Client *ent.Client
	Ctx    *context.Context
}

// New make database connection
func New() (*DatabaseConnection, context.CancelFunc) {
	dsn, err := dsn()
	if err != nil {
		return nil, nil
	}
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Error().Err(err).Msg("error while opening database connection")
	}
	client = (*client).Debug()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Error().Err(err).Msg("failed printing schema changes")
	}
	// create ent schemas into database if there are not exists in db. this function uses context.Background.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Error().Err(err).Msg("failed creating schema changes")
	}
	return &DatabaseConnection{
		client, &ctx,
	}, cancel
}
