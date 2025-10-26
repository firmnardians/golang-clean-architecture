package repository

import (
	"errors"

	"fiber/internal/entity"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
	Create(user entity.User) error
	FindUser(id int) (*entity.User, error)
	DeleteUser(id int) error
	UpdateUser(data entity.User) error
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

func (r *userRepository) FindUser(id int) (*entity.User, error) {
	users := r.data

	var found *entity.User
	for _, u := range users {
		if u.ID == id {
			found = &u
			break
		}
	}

	if found == nil {
		return found, errors.New("USER_NOT_FOUND")
	}

	return found, nil
}

func (r *userRepository) DeleteUser(id int) error {
	users := r.data

	dataUsers := []entity.User{}
	for _, u := range users {
		if u.ID != id {
			dataUsers = append(dataUsers, u)
		}
	}

	r.data = dataUsers
	return nil
}

func (r *userRepository) UpdateUser(user entity.User) error {
	for i, u := range r.data {
		if u.ID == user.ID {
			r.data[i] = user
			return nil
		}
	}

	return errors.New("USER_NOT_FOUND")
}
