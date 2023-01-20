package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/serjbibox/friend-service/pkg/handler"
	"github.com/serjbibox/friend-service/pkg/server"
	"github.com/serjbibox/friend-service/pkg/storage"
	"github.com/serjbibox/friend-service/pkg/storage/postgresql"
)

// Конфигурация приложения
type config struct {
	PostgresConfig postgresql.PostgresConfig `json:"postgres_settings"`
}

var ctx = context.Background()

func main() {
	c, err := readConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	connString, err := postgresql.GetConnectionString(c.PostgresConfig)
	log.Println("Connection String is:", connString)
	if err != nil {
		log.Fatal(err)
	}
	db, err := postgresql.New(connString)
	if err != nil {
		log.Fatal(err)
	}
	s, err := storage.NewStoragePostgres(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	handlers, err := handler.New(s)
	if err != nil {
		log.Fatal(err)
	}
	srv := new(server.Server)
	log.Fatal(srv.Run(server.HTTP_PORT, handlers.InitRoutes()))
}

//Чтение JSON файла конфигурации
func readConfig(path string) (*config, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config config
	err = json.Unmarshal(c, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
