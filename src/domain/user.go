package domain

type (
	User struct {
		ID       uint
		Email    string
		Name     string
		Password string
	}

	UserRepository interface {
		Store(user *User) (int64, error)
		FindById(id uint) *User
		FindByEmail(email string) *User
	}
)
