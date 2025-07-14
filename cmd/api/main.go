package main

import (
	"github.com/dangLuan01/rebuild-api-movie28/internal/app"
	"github.com/dangLuan01/rebuild-api-movie28/internal/config"
)

func main() {
	
	app.LoadEnv()

	cfg := config.NewConfig()

	application := app.NewApplication(cfg)

	if err := application.Run(); err != nil {
		panic(err)
	}
}