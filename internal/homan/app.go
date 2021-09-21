//go:build wireinject
// +build wireinject

package homan

import (
	"embed"
	"github.com/anantadwi13/cli-whm/internal/homan/domain"
	"github.com/anantadwi13/cli-whm/internal/homan/domain/service"
	"github.com/anantadwi13/cli-whm/internal/homan/domain/usecase"
	service2 "github.com/anantadwi13/cli-whm/internal/homan/external/service"
	usecase2 "github.com/anantadwi13/cli-whm/internal/homan/external/usecase"
	"github.com/google/wire"
)

var (
	//go:embed template/*
	Templates embed.FS
)

var useCasesSet = wire.NewSet(
	usecase2.NewUcInit,
	usecase.NewUcUp,
	usecase.NewUcDown,
	usecase2.NewUcAdd,
	usecase2.NewUcRemove,
	wire.Struct(new(useCases), "Init", "Up", "Down", "Add", "Remove"),
)

var serviceSet = wire.NewSet(
	service2.NewCommander,
	service2.NewLocalRegistry,
	service2.NewDockerExecutor,
	service.NewStorage,
	service2.NewDockerProxy,
	wire.Struct(new(services), "Commander", "Executor", "Registry", "Storage", "Proxy"),
)

var applicationSet = wire.NewSet(
	useCasesSet,
	wire.Value(Templates),
	wire.Struct(new(App), "UseCases", "Config"),
)

type useCases struct {
	Init   usecase.UcInit
	Up     usecase.UcUp
	Down   usecase.UcDown
	Add    usecase.UcAdd
	Remove usecase.UcRemove
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
	Config   domain.Config
}

func NewApp(config domain.Config) App {
	panic(wire.Build(applicationSet, serviceSet))
}
