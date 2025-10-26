package usescase

import (
	"fiber/internal/entity"
	"fiber/internal/repository"
)

type UserUsecase interface {
	GetUsers() ([]entity.User, error)
	CreateUsers(user entity.User) error
	DeleteUsers(key string) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{repo: r}
}

func (u *userUsecase) GetUsers() ([]entity.User, error) {
	return u.repo.GetAll()
}

func (u *userUsecase) CreateUsers(user entity.User) error {
	return u.repo.Create(user)
}

func (u *userUsecase) DeleteUsers(key string) error {
	return u.repo.DeleteUser(key)
}
