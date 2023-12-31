package repository

import (
	"rest-api-redis/pkg/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUsertRepostiory(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (p *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := p.DB.Debug().Model(models.User{}).Find(&users).Error
	return users, err
}

func (p *UserRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	err := p.DB.Debug().Model(&models.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func (p *UserRepository) Create(user models.User) (models.User, error) {
	err := p.DB.Debug().Model(&models.User{}).Create(&user).Error
	return user, err
}

func (p *UserRepository) Update(user models.User) (models.User, error) {
	err := p.DB.Debug().Model(&models.User{}).Save(&user).Error
	return user, err
}

func (p *UserRepository) Delete(user models.User) {
	p.DB.Debug().Model(&models.User{}).Delete(&user)
}
