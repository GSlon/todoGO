package repository

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    "github.com/sirupsen/logrus"
)

const (
    userTable = "users"
    todoitemsTable = "todoitems"
)

type Config struct {
    Host string
    Port string
    Username string
    Password string
    DBName string
    SSLMode string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
    logrus.Info(fmt.Sprintf(
        "host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
        cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

    var config string
    if cfg.Password == "" {
        config = fmt.Sprintf(
        "host=%s port=%s user=%s dbname=%s sslmode=%s",
        cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode)
    } else {
        config = fmt.Sprintf(
        "host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
        cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
    }

    db, err := sqlx.Open("postgres", config)

    if err != nil {
        return nil, err
    }

    logrus.Info("connection to postgres establish")

    return db, nil
}
