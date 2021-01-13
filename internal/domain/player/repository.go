package player

type Repository interface {
	NextIdentity() Identity
	Save(p Player)
}
