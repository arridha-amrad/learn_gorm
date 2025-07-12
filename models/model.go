package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string
	Email   string `gorm:"unique"`
	Account Account
}

type Account struct {
	gorm.Model
	UserID  uint `gorm:"unique"`
	Balance float64
}
