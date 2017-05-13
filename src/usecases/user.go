package usecases

import "github.com/stacknowledge/go-clean-todo/src/domain"

type UserInteractor struct {
	UserRepository domain.UserRepository
}

/*func (interactor *UserInteractor) Login(username string, password string) {

}*/

func (interactor *UserInteractor) Register(user domain.User) (int64, error) {
	return interactor.UserRepository.Store(user)
}

func (interactor *UserInteractor) Profile(id uint) *domain.User {
	return interactor.UserRepository.FindById(id)
}
