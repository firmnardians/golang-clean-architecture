package repository

import (
	"errors"
	"strconv"

	"fiber/internal/entity"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
	Create(user entity.User) error
	DeleteUser(key string) error
}

type userRepository struct {
	data []entity.User
}

func NewUserRepository() UserRepository {
	return &userRepository{
		data: []entity.User{},
	}
}

func (r *userRepository) GetAll() ([]entity.User, error) {
	return r.data, nil
}

func (r *userRepository) Create(user entity.User) error {
	user.ID = len(r.data) + 1

	r.data = append(r.data, user)
	return nil
}

func (r *userRepository) DeleteUser(key string) error {
	num, err := strconv.Atoi(key)
	if err != nil {
		return nil
	}

	users := r.data
	keyInt := num

	var found *entity.User
	for _, u := range users {
		if u.ID == keyInt {
			found = &u
			break
		}
	}

	if found == nil {
		return errors.New("USER_NOT_FOUND")
	}

	dataUsers := []entity.User{}
	for _, u := range users {
		if u.ID != keyInt {
			dataUsers = append(dataUsers, u)
		}
	}

	r.data = dataUsers
	return nil
}
