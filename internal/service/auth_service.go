package service

import (
	"errors"
	"merchant-management/internal/model"
	"merchant-management/internal/repository"
	"merchant-management/pkg/hash"
	"merchant-management/pkg/jwt"
)

type AuthService interface {
	Login(username, password string) (*model.User, string, error)                 // Login returns a JWT token
	Register(username, name, email, password string) (*model.User, string, error) // Register creates a new user and returns a JWT token
}

type authService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(username, password string) (*model.User, string, error) {
	var hasher hash.Hasher = hash.BcryptHasher{}

	user, err := s.repo.FindByUsernameOrEmail(username)
	if err != nil {
		return nil, "", err
	}

	if user == nil {
		return nil, "", errors.New("user not found")
	}

	isValid, err := hasher.Compare(user.Password, password)
	if err != nil || !isValid {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := jwt.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *authService) Register(username, name, email, password string) (*model.User, string, error) {
	var hasher hash.Hasher = hash.BcryptHasher{}

	// Check if user already exists
	existingUser, err := s.repo.FindByUsernameOrEmail(username)
	if err != nil {
		return nil, "", err
	}
	if existingUser != nil {
		return nil, "", errors.New("user already exists")
	}

	hashedPassword, err := hasher.Hash(password)
	if err != nil {
		return nil, "", err
	}

	user := &model.User{
		Username: username,
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	newUser, err := s.repo.Create(user)
	if err != nil {
		return nil, "", err
	}

	token, err := jwt.GenerateToken(newUser.ID, newUser.Role)
	if err != nil {
		return nil, "", err
	}

	return newUser, token, nil
}
