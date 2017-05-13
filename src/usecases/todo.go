package usecases

import "github.com/stacknowledge/go-clean-todo/src/domain"

type TodoInteractor struct {
	TodoRepository domain.TodoRepository
}

func (interactor *TodoInteractor) Store(todo *domain.Todo) (int64, error) {
	return interactor.TodoRepository.Store(todo)
}

func (interactor *TodoInteractor) Delete(id uint) error {
	return interactor.TodoRepository.Delete(id)
}

func (interactor *TodoInteractor) FindById(id uint) *domain.Todo {
	return interactor.TodoRepository.FindById(id)
}

func (interactor *TodoInteractor) Update(todo *domain.Todo) (int64, error) {
	return interactor.TodoRepository.Update(todo)
}
