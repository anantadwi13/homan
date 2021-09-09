package main

import (
	"context"
	"github.com/anantadwi13/cli-whm/internal"
	"github.com/anantadwi13/cli-whm/internal/domain"
	"github.com/anantadwi13/cli-whm/internal/usecase"
	"log"
	"time"
)

func main() {
	config, err := domain.NewConfig(domain.ConfigParams{})
	if err != nil {
		panic(err)
	}
	app := internal.NewApp(config)
	err = app.UseCases.Init.Execute(context.TODO(), nil)
	if err != nil && err != usecase.ErrorUcInitAlreadyInitialized {
		log.Println(err)
	}
	err = app.UseCases.Up.Execute(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(5 * time.Second)
	err = app.UseCases.Down.Execute(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}
}
