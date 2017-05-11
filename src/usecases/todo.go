package usecases

import "github.com/stacknowledge/go-clean-todo/src/domain"

type TodoInteractor struct {
	todoRepository domain.TodoRepository
}

func (interactor *TodoInteractor) Add(todo domain.Todo) {
	interactor.todoRepository.Store(todo)
}

func (interactor *TodoInteractor) Delete(id uint) {
	interactor.todoRepository.Delete(id)
}

func (interactor *TodoInteractor) Update(todo domain.Todo) {
	interactor.todoRepository.Update(todo)
}
