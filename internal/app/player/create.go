package player

import "github.com/xavividal/ddd-go-example/internal/domain/player"

type CreatePlayerRQ struct {
	Name string
}

type CreatePlayerAssembler interface {
	Assemble(p player.Player) interface{}
}

type CreatePlayer struct {
	players   player.Repository
	assembler CreatePlayerAssembler
}

func NewCreatePlayer(p player.Repository, a CreatePlayerAssembler) *CreatePlayer {
	return &CreatePlayer{
		players: p,
		assembler: a,
	}
}

func (c CreatePlayer) Execute(rq CreatePlayerRQ) interface{} {
	p := player.New(
		c.players.NextIdentity(),
		rq.Name)

	c.players.Save(p)

	return c.assembler.Assemble(p)
}
