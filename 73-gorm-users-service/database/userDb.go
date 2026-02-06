package database

import (
	"user-service/models"

	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

type IUserDB interface {
	Create(user *models.User) (*models.User, error)
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db}
}

func (u *UserDB) Create(user *models.User) (*models.User, error) {
	err := u.DB.AutoMigrate(models.User{}) // It creates a table in the database
	if err != nil {
		return nil, err
	}
	tx := u.DB.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
