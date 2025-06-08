package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password *string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"default:'user'"`
	CreatedAt gorm.DeletedAt `json:"created_at" gorm:"index"`
	UpdatedAt gorm.DeletedAt `json:"updated_at" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
