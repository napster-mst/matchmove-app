package main

import (
	"fmt"
	"log"

	"github.com/go-pg/pg"
	"matchmove.com/demo/common/config"
	"matchmove.com/demo/handlers"
	"matchmove.com/demo/seeder"
)

func main() {
	conf := config.GetConfigs()
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprint(conf.Database.Host + ":" + conf.Database.Port),
		User:     conf.Database.User,
		Password: conf.Database.Password,
		Database: conf.Database.Database,
	})

	seeder.Seed(db)
	log.Println("server loaded at: ", conf.Port.Port)
	handlers.Routes(conf.Port.Port, db)
}
