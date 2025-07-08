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

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
    result := r.db.Create(user)
    if result.Error != nil {
        return nil, result.Error
    }
    return user, nil
}

func (r *UserRepository) FindAll() ([]model.User, error) {
    var users []model.User
    result := r.db.Omit("password").Find(&users)
    if result.Error != nil {
        return nil, result.Error
    }
    return users, nil
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
    var user model.User
    result := r.db.Omit("password").First(&user, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func (r *UserRepository) FindByEmail(email string, includePassword bool) (*model.User, error) {
    var user model.User
    query := r.db
    if !includePassword {
        query = query.Omit("password")
    }
    result := query.Where("email = ?", email).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, result.Error
    }
    return &user, nil
}

func (r *UserRepository) FindByUsername(username string, includePassword bool) (*model.User, error) {
    var user model.User
    query := r.db
    if !includePassword {
        query = query.Omit("password")
    }
    result := query.Where("username = ?", username).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, result.Error
    }
    return &user, nil
}

func (r *UserRepository) FindByUsernameOrEmail(username string) (*model.User, error) {
    var user model.User
    result := r.db.Where("username = ? OR email = ?", username, username).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, result.Error
    }
    return &user, nil
}

func (r *UserRepository) Update(user *model.User) (*model.User, error) {
    result := r.db.Save(user)
    if result.Error != nil {
        return nil, result.Error
    }
    return user, nil
}

func (r *UserRepository) Delete(id uint) error {
    result := r.db.Delete(&model.User{}, id)
    return result.Error
}