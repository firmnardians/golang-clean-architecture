package repository

import "fiber/internal/entity"

type UserRepository interface {
	GetAll() ([]entity.User, error)
	Create(user entity.User) error
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
