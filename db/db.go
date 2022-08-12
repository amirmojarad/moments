package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"moments/ent"
	"os"
	"time"
)

func dsnTest() (string, error) {
	host := os.Getenv("TEST_DB_HOST")
	if host == "" {
		err := errors.New("DB_Host is empty in env file")
		log.Error().Err(err).Msg("")
		return "", err
	}

	password := os.Getenv("TEST_DB_PASSWORD")
	if password == "" {
		err := errors.New("DB_PASSWORD is empty in env file")
		log.Error().Err(err).Msg("")
		return "", err
	}

	user := os.Getenv("TEST_DB_USER")
	if user == "" {
		err := errors.New("DB_USER is empty in env file")
		log.Error().Err(err).Msg("")
		return "", err
	}
	name := os.Getenv("TEST_DB_NAME")
	if name == "" {
		err := errors.New("DB_NAME is empty in env file")
		log.Error().Err(err).Msg("")
		return "", err
	}
	port := os.Getenv("TEST_DB_PORT")
	if port == "" {
		err := errors.New("DB_PORT is empty in env file")
		log.Error().Err(err).Msg("")
		return "", err
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		host, user, password, name, port), nil
}

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

func loadDbEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return err
}

// New make database connection
func New() (*DatabaseConnection, context.CancelFunc) {
	loadDbEnv()
	dsn, err := dsn()
	if err != nil {
		return nil, nil
	}
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Error().Err(err).Msg("error while opening database connection")
	}
	client = (*client).Debug()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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

func NewTestDB() (*DatabaseConnection, context.CancelFunc) {
	dsn, err := dsnTest()
	if err != nil {
		return nil, nil
	}
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Error().Err(err).Msg("error while opening database connection")
	}
	client = (*client).Debug()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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

func Clear() {
	dsn, _ := dsnTest()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal().Err(err)
	}
	_, err = db.Query("DROP DATABASE moments_test")
}
