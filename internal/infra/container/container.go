package container

import (
	appPlayer "github.com/xavividal/ddd-go-example/internal/app/player"
	domainPlayer "github.com/xavividal/ddd-go-example/internal/domain/player"
	assemblerAppPlayer "github.com/xavividal/ddd-go-example/internal/infra/assembler/app/player"
	assemblerDomainPlayer "github.com/xavividal/ddd-go-example/internal/infra/assembler/domain/player"
	persistencePlayer "github.com/xavividal/ddd-go-example/internal/infra/persistence/player"
)

var instance *Container

// Container ...
type Container struct {
	App struct {
		Player struct {
			CreatePlayer          *appPlayer.CreatePlayer
			CreatePlayerAssembler appPlayer.CreatePlayerAssembler
		}
	}
	Domain struct {
		Player struct {
			Repository domainPlayer.Repository
		}
	}
	Infra struct {
		Assembler struct {
			Domain struct {
				Player struct {
					Assembler *assemblerDomainPlayer.Assembler
				}
			}
		}
	}
}

func init() {
	instance = &Container{}
	instance.Infra.Assembler.Domain.Player.Assembler = assemblerDomainPlayer.New()

	instance.Domain.Player.Repository = persistencePlayer.NewMemoryRepository()

	instance.App.Player.CreatePlayerAssembler = assemblerAppPlayer.NewCreatePlayerAssembler(
		instance.Infra.Assembler.Domain.Player.Assembler,
	)

	instance.App.Player.CreatePlayer = appPlayer.NewCreatePlayer(
		instance.Domain.Player.Repository,
		instance.App.Player.CreatePlayerAssembler,
	)
}

// Instance returns the one and only container
func Instance() *Container {
	return instance
}
