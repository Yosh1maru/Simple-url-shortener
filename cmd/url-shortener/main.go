package main

import (
	"database/sql"
	"fmt"
	"github.com/codingconcepts/env"
	_ "github.com/lib/pq"
	"log"
	"url-shortener/internal"
	"url-shortener/internal/api/controller"
	"url-shortener/internal/api/response"
	"url-shortener/internal/api/server"
	"url-shortener/internal/postgresql"
)

func main() {
	var cfg Config
	if err := env.Set(&cfg); err != nil {
		log.Fatal("Failed to parse config.")
	}

	connectionDb := startDbConnection(cfg)

	encryptor := internal.NewEncryptor(cfg.EncryptorSalt)
	urlSource := postgresql.NewUrlSource(connectionDb)
	urlController := controller.NewUrlApiController(encryptor, response.NewJsonResponse(), urlSource)
	srv := server.NewServer(cfg.HttpServerHost+":"+cfg.HttpServerPort, urlController)
	srv.Start()
}

func startDbConnection(cfg Config) *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbName,
	)
	connectionDb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("DB connection error : %s", err)
	}

	return connectionDb
}
