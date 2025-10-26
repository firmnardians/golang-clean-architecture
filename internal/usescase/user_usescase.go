package usescase

import (
	"errors"
	"strconv"

	"fiber/internal/entity"
	"fiber/internal/repository"
)

type UserUsecase interface {
	GetUsers() ([]entity.User, error)
	CreateUsers(user entity.User) error
	DeleteUsers(key string) error
	UpdateUsers(user entity.User) error
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
	id, err := strconv.Atoi(key)
	if err != nil {
		return errors.New("INVALID_ID")
	}

	user, err := u.repo.FindUser(id)
	if err != nil {
		return err
	}

	return u.repo.DeleteUser(user.ID)
}

func (u *userUsecase) UpdateUsers(user entity.User) error {
	return u.repo.UpdateUser(user)
}
