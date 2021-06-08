package main

import (
	"github.com/joho/godotenv"
	"github.com/p-12s/own-golang-manual/0-golang-test-assignment/wildberries/http-api"
	"github.com/p-12s/own-golang-manual/0-golang-test-assignment/wildberries/http-api/pkg/handler"
	"github.com/p-12s/own-golang-manual/0-golang-test-assignment/wildberries/http-api/pkg/repository"
	"github.com/p-12s/own-golang-manual/0-golang-test-assignment/wildberries/http-api/pkg/service"
	"github.com/spf13/viper"
	"log"
	_ "github.com/lib/pq"
	"os"
)

// TODO добавить доку
// TODO переделать логирование на Zap

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error init configs: %s\n", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("Failed to initialize DB: $s\n", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(common.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
