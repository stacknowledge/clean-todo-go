package interfaces

import (
	"fmt"

	"github.com/stacknowledge/clean-todo-go/src/domain"
	"github.com/stacknowledge/clean-todo-go/src/infrastructure"
)

type DbTodoRepository infrastructure.DbRepository

func NewDbTodoRepository(dbHandlers map[string]infrastructure.DbHandler) *DbTodoRepository {
	return &DbTodoRepository{
		DbHandlers: dbHandlers,
		DbHandler:  dbHandlers["TodoRepository"],
	}
}

func (repository *DbTodoRepository) Store(todo *domain.Todo) (int64, error) {
	return repository.DbHandler.Execute(fmt.Sprintf(
		"INSERT INTO todos (title, expires_at, user_id) VALUES ('%v', '%v', '%d')",
		todo.Title, todo.Expires, todo.User,
	))
}

func (repository *DbTodoRepository) Update(todo *domain.Todo) (int64, error) {
	return repository.DbHandler.Execute(fmt.Sprintf(
		"UPDATE todos SET title = '%v', expires_at = '%v' WHERE id = %d",
		todo.Title, todo.Expires, todo.ID,
	))
}

func (repository *DbTodoRepository) Delete(id uint) error {
	_, deleteError := repository.DbHandler.Execute(fmt.Sprintf(
		"DELETE FROM todos WHERE id = %d",
		id,
	))

	return deleteError
}

func (repository *DbTodoRepository) FindById(id uint) *domain.Todo {
	row := repository.DbHandler.Query(fmt.Sprintf(
		"SELECT id, title, expires_at FROM todos WHERE id = '%d' LIMIT 1",
		id,
	))

	return hydrateTodoRepositoryDataToDomain(row)
}

func hydrateTodoRepositoryDataToDomain(row infrastructure.Row) *domain.Todo {
	var (
		id             uint
		title, expires string
	)

	defer row.Close()
	row.Next()
	row.Scan(&id, &title, &expires)

	return &domain.Todo{
		ID:      id,
		Title:   title,
		Expires: expires,
	}
}
