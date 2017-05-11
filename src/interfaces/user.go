package interfaces

import (
	"fmt"

	"github.com/stacknowledge/go-clean-todo/src/domain"
	infrastructure "github.com/stacknowledge/go-clean-todo/src/infrastructure/database"
)

type DbUserRepository infrastructure.DbRepository

func NewDbUserRepository(dbHandlers map[string]infrastructure.DbHandler) *DbUserRepository {
	return &DbUserRepository{
		DbHandlers: dbHandlers,
		DbHandler:  dbHandlers["UserRepository"],
	}
}

func (repo *DbUserRepository) Store(user domain.User) (int64, error) {
	return repo.DbHandler.Execute(fmt.Sprintf(
		"INSERT INTO users (email, name, password) VALUES ('%v', '%v', '%v')",
		user.Email, user.Name, user.Password,
	))
}

func (repo *DbUserRepository) FindById(userId uint) *domain.User {
	row := repo.DbHandler.Query(fmt.Sprintf(
		"SELECT id, name, email, password FROM users WHERE id = '%d' LIMIT 1",
		userId,
	))

	var id uint
	var email, name, password string

	row.Next()
	row.Scan(&id, &name, &email, &password)

	return &domain.User{
		ID:       id,
		Email:    email,
		Name:     name,
		Password: password,
	}
}
