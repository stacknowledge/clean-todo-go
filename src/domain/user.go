package domain

type UserRepository interface {
	Store(user User) (int64, error)
	FindById(id uint) *User
	/*FindByEmail(email string) *User
	Deactivate(id uint)
	Activate(id uint)*/
}

type User struct {
	ID       uint
	Email    string
	Name     string
	Password string
}
