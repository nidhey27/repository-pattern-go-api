package repository

import (
	"rest-api-redis/pkg/models"

	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	FindAll(offset int, limit int) ([]models.User, error)
	FindByID(id uint) (models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id uint) error
}

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (p *UserRepository) FindAll(page int, limit int) ([]models.User, error) {
	var users []models.User
	err := p.DB.Model(models.User{}).Offset(limit).Limit(page).Find(&users).Error
	return users, err
}

func (p *UserRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	err := p.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func (p *UserRepository) Create(user *models.User) (*models.User, error) {
	err := p.DB.Model(&models.User{}).Create(&user).Error
	return user, err
}

func (p *UserRepository) Update(user *models.User) (*models.User, error) {
	err := p.DB.Model(&models.User{}).Save(&user).Error
	return user, err
}

func (p *UserRepository) Delete(id uint) error {
	return p.DB.Model(&models.User{}).Delete(&models.User{}, "id = ?", id).Error
}
