package service

import (
	"merchant-management/internal/model"
	"merchant-management/internal/repository"
)

type UserService interface {
	GetAllUser() ([]model.User, error)
	GetUserByID(id uint) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// Create implements Service.
func (s *userService) GetAllUser() ([]model.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetUserByID(id uint) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) CreateUser(user *model.User) (*model.User, error) {
	createdUser, err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *userService) UpdateUser(user *model.User) (*model.User, error) {
	updatedUser, err := s.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (s *userService) DeleteUser(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
