package repositories

import (
	"learn_gorm/models"

	"gorm.io/gorm"
)

type userRepo struct{}

type IUserRepo interface {
	Create(tx *gorm.DB, user *models.User) error
	FindOne(tx *gorm.DB, id uint) (*models.User, error)
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}

func (r *userRepo) Create(tx *gorm.DB, user *models.User) error {
	return tx.Create(user).Error
}

func (r *userRepo) FindOne(tx *gorm.DB, id uint) (*models.User, error) {
	var user models.User
	err := tx.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
