package repository

import "github.com/jinzhu/gorm"

type Repository struct {
	UserRepository *UserRepository
}

func InitRepository(db *gorm.DB) *Repository {
	userRepository := ProvideUserRepository(db)
	return &Repository{UserRepository: &userRepository}
}
