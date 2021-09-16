//go:build wireinject
// +build wireinject

package internal

import (
	"embed"
	"github.com/anantadwi13/cli-whm/internal/domain"
	domainService "github.com/anantadwi13/cli-whm/internal/domain/service"
	domainUsecase "github.com/anantadwi13/cli-whm/internal/domain/usecase"
	externalService "github.com/anantadwi13/cli-whm/internal/external/service"
	externalUsecase "github.com/anantadwi13/cli-whm/internal/external/usecase"
	"github.com/google/wire"
)

var (
	//go:embed template/*
	Templates embed.FS
)

var useCasesSet = wire.NewSet(
	externalUsecase.NewUcInit,
	domainUsecase.NewUcUp,
	domainUsecase.NewUcDown,
	externalUsecase.NewUcAdd,
	externalUsecase.NewUcRemove,
	wire.Struct(new(useCases), "Init", "Up", "Down", "Add", "Remove"),
)

var serviceSet = wire.NewSet(
	externalService.NewCommander,
	externalService.NewLocalRegistry,
	externalService.NewDockerExecutor,
	domainService.NewStorage,
	externalService.NewDockerProxy,
	wire.Struct(new(services), "Commander", "Executor", "Registry", "Storage", "Proxy"),
)

var applicationSet = wire.NewSet(
	useCasesSet,
	wire.Value(Templates),
	wire.Struct(new(App), "UseCases", "Config"),
)

type useCases struct {
	Init   domainUsecase.UcInit
	Up     domainUsecase.UcUp
	Down   domainUsecase.UcDown
	Add    domainUsecase.UcAdd
	Remove domainUsecase.UcRemove
}

type services struct {
	Commander externalService.Commander
	Executor  domainService.Executor
	Registry  domainService.Registry
	Storage   domainService.Storage
	Proxy     domainService.Proxy
}

type App struct {
	UseCases useCases
	Config   domain.Config
}

func NewApp(config domain.Config) App {
	panic(wire.Build(applicationSet, serviceSet))
}
