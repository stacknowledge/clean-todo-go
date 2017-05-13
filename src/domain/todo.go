package domain

type (
	TodoRepository interface {
		Store(todo *Todo) (int64, error)
		FindById(id uint) *Todo
		Update(todo *Todo) (int64, error)
		Delete(id uint) error
	}

	Todo struct {
		ID      uint
		Title   string
		Expires string
		User    *User
	}
)
