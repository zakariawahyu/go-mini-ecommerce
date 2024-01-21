package main

import (
	"go-mini-ecommerce/config"
	"go-mini-ecommerce/internal/infrastructure/cache"
	"go-mini-ecommerce/internal/infrastructure/db"
	"go-mini-ecommerce/internal/server"
	"log"
)

// @title Golang Mini Ecommerce API
// @version 1.0
// @description Go Mini Ecommerce
// @termsOfService http://swagger.io/terms/
// @contact.name Zakaria Wahyu
// @contact.email zakarianur6@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1/
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	mysql, err := db.NewMysqlConnection(cfg)
	if err != nil {
		log.Println(err)
	}

	redis := cache.NewRedisConnection(cfg)

	httpServer := server.NewHttpServer(cfg, mysql, redis)
	if err = httpServer.Run(); err != nil {
		log.Fatal(err)
	}
}
