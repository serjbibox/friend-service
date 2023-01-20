package postgresql

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var elog = log.New(os.Stderr, "postgresql error\t", log.Ldate|log.Ltime|log.Lshortfile)
var ilog = log.New(os.Stdout, "postgresql info\t", log.Ldate|log.Ltime)

//Структура дял хранения конфигурации PostgresSQL
type PostgresConfig struct {
	Username string `json:"username"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"dbname"`
	SslMode  string `json:"sslmode"`
	EnvName  string `json:"envname"`
}

//Формирует строку подключения к PostgresSQL
func GetConnectionString(c PostgresConfig) (string, error) {
	pwd := os.Getenv(c.EnvName)
	if pwd == "" {
		elog.Println("error reading password from os environment")
		return "", errors.New("error reading password from os environment")
	}
	return "postgres://" +
			c.Username + ":" +
			pwd + "@" +
			c.Host + ":" +
			c.Port + "/" +
			c.DbName + "?sslmode=" +
			c.SslMode,
		nil
}

//Конструктор пула подключений PostgreSQL
func New(constr string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, constr)
	if err != nil {
		return nil, err
	}
	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}
	ilog.Println("connected to postgres database")
	return db, nil
}
