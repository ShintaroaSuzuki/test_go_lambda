package services

import (
	"test-go-lambda/models"
	"test-go-lambda/repositories"
)

type UserService interface {
	GetUserByID(id string) (*models.User, error)
}

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (service *UserServiceImpl) GetUserByID(id string) (*models.User, error) {
	return service.userRepository.FindByID(id)
}
