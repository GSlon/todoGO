package main

import (
    "github.com/spf13/viper"
    "github.com/joho/godotenv"
    "github.com/sirupsen/logrus"
    "github.com/GSlon/todoGO" // for server.go import
    "github.com/GSlon/todoGO/internal/handler"
    "github.com/GSlon/todoGO/internal/repository"
    "github.com/GSlon/todoGO/internal/service"
    "os"
)

func initConfig() error {
    viper.AddConfigPath("config")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}

func main() {
    if err := initConfig(); err != nil {
        logrus.Fatalf(err.Error())
    }

    if err := godotenv.Load(); err != nil {
		logrus.Fatalf(err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

    srv := new(server.Server)
    if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
        logrus.Fatalf(err.Error())
    }
}
