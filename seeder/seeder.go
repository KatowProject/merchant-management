package seeder

import (
	"log"
	"merchant-management/internal/model"
	"merchant-management/pkg/hash"

	"gorm.io/gorm"
)

func RunAll(db *gorm.DB) {
	log.Println("Starting database seeding...")

	hasher := hash.BcryptHasher{}
	pass, _ := hasher.Hash("password")

	// create 1 admin user
	adminUser := &model.User{
		Name:     "Admin User",
		Email:    "admin@mail.com",
		Username: "admin",
		Password: &pass,
		Role:     "admin",
	}

	if err := db.Create(adminUser).Error; err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	if err := SeedUsers(db, 20); err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}

	log.Println("Database seeding completed.")
}
