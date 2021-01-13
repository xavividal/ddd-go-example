package player

import (
	"github.com/xavividal/ddd-go-example/internal/domain/player"
	"time"
)

type Assembler struct {}

func New() *Assembler {
	return &Assembler{}
}

func (a Assembler) Assemble(p player.Player) DTO {
	return DTO{
		ID: p.ID().Value(),
		Name: p.Name(),
		CreatedAt: p.CreatedAt().Format(time.RFC850),
	}
}
