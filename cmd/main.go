package main

import (
	"go-mini-ecommerce/config"
	"go-mini-ecommerce/internal/infrastructure/db"
	"go-mini-ecommerce/internal/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	mysql, err := db.NewMysqlConnection(cfg)
	if err != nil {
		log.Println(err)
	}

	httpServer := server.NewHttpServer(cfg, mysql)
	if err = httpServer.Run(); err != nil {
		log.Fatal(err)
	}
}
