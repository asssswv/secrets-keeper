package main

import (
	"context"
	"os"
	"os/signal"
	sc "secrets_keeper/app"
	"secrets_keeper/app/pkg/handler"
	"secrets_keeper/app/pkg/repository"
	"secrets_keeper/app/pkg/service"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}


	mem := make(map[string]string)

	repos := repository.NewRepository(mem)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(sc.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error ocurred while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error ocurred on server shutting down: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
