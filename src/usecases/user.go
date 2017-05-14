package usecases

import (
	"errors"

	"github.com/stacknowledge/clean-todo-go/src/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	UserRepository domain.UserRepository
}

func (interactor *UserInteractor) Login(email string, password string) (*domain.User, error) {
	user := interactor.UserRepository.FindByEmail(email)

	if error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); error != nil {
		return nil, errors.New("Bad credentials!")
	}

	return user, nil
}

func (interactor *UserInteractor) Register(user domain.User) (int64, error) {
	return interactor.UserRepository.Store(user)
}

func (interactor *UserInteractor) Profile(id uint) *domain.User {
	return interactor.UserRepository.FindById(id)
}
