package repositories

import "gorm.io/gorm"

type baseRepo struct {
	DB *gorm.DB
}

type IBaseRepo interface {
	Begin() *gorm.DB
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
}

func NewBaseRepo(db *gorm.DB) IBaseRepo {
	return &baseRepo{DB: db}
}

func (r *baseRepo) Begin() *gorm.DB {
	return r.DB.Begin()
}

func (r *baseRepo) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *baseRepo) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}
