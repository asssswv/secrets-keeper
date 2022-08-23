package main

import (
	"context"
	"os"
	"os/signal"
	sc "secrets_keeper/app"
	"secrets_keeper/app/pkg/handler"
	"secrets_keeper/app/pkg/repository/redis_repo"
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

	rdb := redis_repo.NewRedisClient(redis_repo.Config{
		Host:     viper.GetString("rdb.host"),
		Port:     viper.GetString("rdb.port"),
		Password: viper.GetString("rdb.password"),
	})

	repos := redis_repo.NewRepository(rdb)
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
	viper.SetConfigType("yml")
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
