package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/p-12s/own-golang-manual/0-golang-test-assignment/wildberries/http-api"
	"github.com/p-12s/own-golang-manual/0-golang-test-assignment/wildberries/http-api/pkg/handler"
	"github.com/p-12s/own-golang-manual/0-golang-test-assignment/wildberries/http-api/pkg/repository"
	"github.com/p-12s/own-golang-manual/0-golang-test-assignment/wildberries/http-api/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// TODO добавить доку
func main() {
	// TODO переделать логирование на Zap
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error init configs: %s\n", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(common.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Service Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	logrus.Print("Service Shutting Down")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
