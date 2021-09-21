package main

import (
	"context"
	"github.com/anantadwi13/homan/internal/homan"
	"github.com/anantadwi13/homan/internal/homan/domain"
	"github.com/anantadwi13/homan/internal/homan/domain/model"
	domainService "github.com/anantadwi13/homan/internal/homan/domain/service"
	"github.com/anantadwi13/homan/internal/homan/domain/usecase"
	"github.com/anantadwi13/homan/internal/homan/external/service"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config, errConfig := domain.NewConfig(domain.ConfigParams{})
	if errConfig != nil {
		panic(errConfig)
	}
	app := homan.NewApp(config)

	var err usecase.Error

	err = app.UseCases.Remove.Execute(context.TODO(), &usecase.UcRemoveParams{
		Name: "my-blog",
	})
	if err != nil {
		log.Println(err)
	}

	err = app.UseCases.Down.Execute(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}

	err = app.UseCases.Init.Execute(context.TODO(), nil)
	if err != nil && err != usecase.ErrorUcInitAlreadyInitialized {
		log.Println(err)
	}

	err = app.UseCases.Up.Execute(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}

	err = app.UseCases.Add.Execute(context.TODO(), &usecase.UcAddParams{
		Name:        "my-blog",
		Domain:      "example.local",
		ServiceType: usecase.ServiceTypeBlog,
	})
	if err != nil {
		log.Println(err)
	}

	proxy := service.NewDockerProxy(config, service.NewDockerExecutor(
		config,
		service.NewCommander(),
		service.NewLocalRegistry(config, domainService.NewStorage()),
		domainService.NewStorage(),
	))

	sign := make(chan os.Signal, 1)

	signal.Notify(sign, syscall.SIGINT, syscall.SIGKILL)

	_ = proxy.Execute(context.TODO(), func(proxy *model.ProxyDetail) error {
		select {
		case <-time.After(24 * time.Hour):
		case <-sign:
		}
		return nil
	})
}
