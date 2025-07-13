package repositories

import (
	"learn_gorm/models"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

type IUserRepo interface {
	Create(tx *gorm.DB, user *models.User) error
	FindOne(tx *gorm.DB, id uint) (*models.User, error)
	FindOneWithAccount(tx *gorm.DB, id uint) (*models.User, error)
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{
		db: db,
	}
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

func (r *userRepo) FindOneWithAccount(tx *gorm.DB, id uint) (*models.User, error) {
	if tx == nil {
		tx = r.db
	}
	var user models.User
	if err := tx.Preload("Account", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("user_id", "Balance")
	}).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
