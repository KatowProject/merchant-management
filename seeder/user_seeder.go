package seeder

import (
	"merchant-management/internal/model"
	"merchant-management/pkg/hash"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB, count int) error {
	for i := 1; i <= count; i++ {
		hash := hash.BcryptHasher{}

		pass, _ := hash.Hash("password")

		user := model.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Username: faker.Username(),
			Password: &pass,
		}

		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}
