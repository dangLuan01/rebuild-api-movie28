package main

import (
	"log"

	"github.com/dangLuan01/rebuild-api-movie28/internal/app"
	"github.com/dangLuan01/rebuild-api-movie28/internal/config"
	"github.com/dangLuan01/rebuild-api-movie28/internal/db"
)

func main() {
	
	app.LoadEnv()

	if err := db.InitDB(); err != nil {
		log.Fatalf("unable to connect to sql")
	}

	es, err := db.InitES();
	if err != nil {
		log.Fatalf("unable to connect to EslasticSearch:%v", err)
	}
	
	cfg := config.NewConfig()

	application := app.NewApplication(cfg, db.DB, es)

	if err := application.Run(); err != nil {
		panic(err)
	}
}