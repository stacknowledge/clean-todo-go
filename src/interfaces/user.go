package interfaces

import (
	"fmt"

	"github.com/stacknowledge/clean-todo-go/src/domain"
	"github.com/stacknowledge/clean-todo-go/src/infrastructure"
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

	return hydrateUserRepositoryDataToDomain(row)
}

func (repo *DbUserRepository) FindByEmail(email string) *domain.User {
	row := repo.DbHandler.Query(fmt.Sprintf(
		"SELECT id, name, email, password FROM users WHERE email = '%v' LIMIT 1",
		email,
	))

	return hydrateUserRepositoryDataToDomain(row)
}

func hydrateUserRepositoryDataToDomain(row infrastructure.Row) *domain.User {
	var (
		id                    uint
		email, name, password string
	)

	defer row.Close()
	row.Next()
	row.Scan(&id, &name, &email, &password)

	return &domain.User{
		ID:       id,
		Email:    email,
		Name:     name,
		Password: password,
	}
}
