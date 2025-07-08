package service

import (
	"errors"
	"merchant-management/internal/model"
	"merchant-management/internal/repository"
	"merchant-management/pkg/hash"
)

type UserService interface {
	GetAllUser() ([]model.User, error)
	GetUserByID(id uint) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(id uint, data model.User) (*model.User, error)
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
	var hasher hash.Hasher = hash.BcryptHasher{}
	_, err := s.repo.FindByUsername(user.Username, false)
	if err != nil && err.Error() != "user not found" {
		return nil, err
	}

	existingUser, err := s.repo.FindByEmail(user.Email, false)
	if err != nil && err.Error() != "user not found" {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := hasher.Hash(user.Password)
    if err != nil {
        return nil, err
    }

    user.Password = hashedPassword

	createdUser, err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *userService) UpdateUser(id uint, user model.User) (*model.User, error) {
	// Check if the user exists
	existingUser, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("user not found")
	}

	// Check for duplicate username
	userWithSameUsername, err := s.repo.FindByUsername(user.Username, false)
	if err != nil && err.Error() != "user not found" {
		return nil, err
	}
	if userWithSameUsername != nil && userWithSameUsername.ID != id {
		return nil, errors.New("username already in use")
	}

	// Check for duplicate email
	userWithSameEmail, err := s.repo.FindByEmail(user.Email, false)
	if err != nil && err.Error() != "user not found" {
		return nil, err
	}
	if userWithSameEmail != nil && userWithSameEmail.ID != id {
		return nil, errors.New("email already in use")
	}

	user.ID = existingUser.ID
	updatedUser, err := s.repo.Update(&user)
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
