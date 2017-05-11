package domain

import "time"

type (
	TodoRepository interface {
		Store(todo Todo)
		FindById(id uint)
		Update(todo Todo)
		Delete(id uint)
	}

	Todo struct {
		ID      uint
		Title   string
		Expires time.Time
		User    User
	}
)
