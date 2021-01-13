package player

import "github.com/xavividal/ddd-go-example/internal/domain/player"

type MemoryRepository struct {
	players map[string]player.Player
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		players: make(map[string]player.Player),
	}
}

func (m MemoryRepository) NextIdentity() player.Identity {
	return player.NewIdentity("ABC")
}

func (m *MemoryRepository) Save(p player.Player) {
	m.players[p.ID().Value()] = p
}
