package main

import (
	"fmt"
	"github.com/xavividal/ddd-go-example/internal/app/player"
	assemblerDomainPlayer "github.com/xavividal/ddd-go-example/internal/infra/assembler/domain/player"
	assemblerAppPlayer "github.com/xavividal/ddd-go-example/internal/infra/assembler/app/player"
	persistencePlayer "github.com/xavividal/ddd-go-example/internal/infra/persistence/player"
)

func main() {
	rq := player.CreatePlayerRQ{
		Name: "Han Solo",
	}

	pa := assemblerDomainPlayer.New()

	rs := player.NewCreatePlayer(
		persistencePlayer.NewMemoryRepository(),
		assemblerAppPlayer.NewCreatePlayerAssembler(pa),
		).Execute(rq)

	fmt.Println(rs)
}
