package player

import "time"

type Player struct {
	id        Identity
	name      string
	createdAt time.Time
}

func New(id Identity, name string) Player {
	return Player{
		id: id,
		name: name,
		createdAt: time.Now(),
	}
}

func (p Player) ID() Identity {
	return p.id
}

func (p Player) Name() string {
	return p.name
}

func (p Player) CreatedAt() time.Time {
	return p.createdAt
}
