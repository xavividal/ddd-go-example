package player

type Collection struct {
	items []Player
}

func NewCollection() Collection {
	return Collection{
		items: make([]Player, 0),
	}
}

func (c Collection) Add(p Player) {
	c.items = append(c.items, p)
}
