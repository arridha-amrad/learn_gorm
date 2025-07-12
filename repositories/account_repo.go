package repositories

import (
	"learn_gorm/models"

	"gorm.io/gorm"
)

type accountRepo struct{}

type IAccountRepo interface {
	Create(tx *gorm.DB, account *models.Account) error
	FindOne(tx *gorm.DB, id int) (*models.Account, error)
	FindByUserId(tx *gorm.DB, userId uint) (*models.Account, error)
	UpdateByUserId(tx *gorm.DB, userId uint, account models.Account) error
}

func NewAccountRepo() IAccountRepo {
	return &accountRepo{}
}

func (r *accountRepo) Create(tx *gorm.DB, account *models.Account) error {
	return tx.Create(account).Error
}

func (r *accountRepo) FindOne(tx *gorm.DB, id int) (*models.Account, error) {
	var account models.Account
	err := tx.First(&account, id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepo) FindByUserId(tx *gorm.DB, userId uint) (*models.Account, error) {
	var account models.Account
	err := tx.Where(&models.Account{UserID: userId}).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepo) UpdateByUserId(tx *gorm.DB, userId uint, account models.Account) error {
	var acc models.Account
	if err := tx.Where(&models.Account{UserID: userId}).First(&acc).Error; err != nil {
		return err
	}
	return tx.Model(&acc).Updates(account).Error
}
