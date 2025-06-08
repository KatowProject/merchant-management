package repository

import (
	"errors"
	"merchant-management/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) (model.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return *user, nil
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) FindByUsernameOrEmail(username string) (*model.User, error) {
	var user model.User
	result := r.db.Where("username = ? OR email = ?", username, user).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found") // User not found
		}
		return nil, result.Error // Other error
	}
	return &user, nil
}

func (r *UserRepository) Update(user *model.User) (model.User, error) {
	// Update the user in the database
	result := r.db.Save(user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return *user, nil
}

func (r *UserRepository) Delete(id uint) error {
	// Delete the user from the database
	result := r.db.Delete(&model.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
