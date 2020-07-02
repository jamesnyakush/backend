package store

type Store interface {
	Users() UserStore
}

type UserStore interface {
	New()
	Get()
}