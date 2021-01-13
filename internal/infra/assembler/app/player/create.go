package player

import (
	domainPlayer "github.com/xavividal/ddd-go-example/internal/domain/player"
	"github.com/xavividal/ddd-go-example/internal/infra/assembler/domain/player"
)

type CreatePlayerAssembler struct {
	assembler *player.Assembler
}

func NewCreatePlayerAssembler(p *player.Assembler) *CreatePlayerAssembler {
	return &CreatePlayerAssembler{
		assembler: p,
	}
}

func (c CreatePlayerAssembler) Assemble(p domainPlayer.Player) interface{} {
	return c.assembler.Assemble(p)
}
