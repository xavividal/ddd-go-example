package player

type Identity struct {
	value string
}

func NewIdentity(v string) Identity {
	return Identity{value: v}
}

func (i Identity) Value() string {
	return i.value
}
