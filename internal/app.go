//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/anantadwi13/cli-whm/internal/domain"
	domainService "github.com/anantadwi13/cli-whm/internal/domain/service"
	"github.com/anantadwi13/cli-whm/internal/domain/usecase"
	externalService "github.com/anantadwi13/cli-whm/internal/external/service"
	"github.com/google/wire"
)

var useCasesSet = wire.NewSet(
	usecase.NewUcInit,
	usecase.NewUcUp,
	usecase.NewUcDown,
	usecase.NewUcAdd,
	wire.Struct(new(useCases), "Init", "Up", "Down", "Add"),
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
	wire.Struct(new(App), "UseCases"),
)

type useCases struct {
	Init usecase.UcInit
	Up   usecase.UcUp
	Down usecase.UcDown
	Add  usecase.UcAdd
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
}

func NewApp(config domain.Config) App {
	panic(wire.Build(applicationSet, serviceSet))
}
