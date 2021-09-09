// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package internal

import (
	"github.com/anantadwi13/cli-whm/internal/domain"
	"github.com/anantadwi13/cli-whm/internal/domain/service"
	service2 "github.com/anantadwi13/cli-whm/internal/external/service"
	"github.com/anantadwi13/cli-whm/internal/usecase"
	"github.com/google/wire"
)

// Injectors from app.go:

func NewApp(config domain.Config) App {
	storage := service.NewStorage()
	registry := service2.NewLocalRegistry(config, storage)
	ucInit := usecase.NewUcInit(registry, config, storage)
	commander := service2.NewCommander()
	executor := service2.NewDockerExecutor(config, commander, registry)
	ucUp := usecase.NewUcUp(registry, executor)
	ucDown := usecase.NewUcDown(registry, executor)
	ucAdd := usecase.NewUcAdd(config, registry, executor)
	internalUseCases := useCases{
		Init: ucInit,
		Up:   ucUp,
		Down: ucDown,
		Add:  ucAdd,
	}
	app := App{
		UseCases: internalUseCases,
	}
	return app
}

// app.go:

var useCasesSet = wire.NewSet(usecase.NewUcInit, usecase.NewUcUp, usecase.NewUcDown, usecase.NewUcAdd, wire.Struct(new(useCases), "Init", "Up", "Down", "Add"))

var serviceSet = wire.NewSet(service2.NewCommander, service2.NewLocalRegistry, service2.NewDockerExecutor, service.NewStorage, service2.NewDockerProxy, wire.Struct(new(services), "Commander", "Executor", "Registry", "Storage", "Proxy"))

var applicationSet = wire.NewSet(
	useCasesSet, wire.Struct(new(App), "UseCases"),
)

type useCases struct {
	Init usecase.UcInit
	Up   usecase.UcUp
	Down usecase.UcDown
	Add  usecase.UcAdd
}

type services struct {
	Commander service2.Commander
	Executor  service.Executor
	Registry  service.Registry
	Storage   service.Storage
	Proxy     service.Proxy
}

type App struct {
	UseCases useCases
}
