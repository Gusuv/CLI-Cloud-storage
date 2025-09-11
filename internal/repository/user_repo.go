package repository

import (
	"cloud_storage/internal/models"
	"gorm.io/gorm"
)

type UserRep struct {
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRep {
	return &UserRep{Db: db}
}

func (rep *UserRep) Register(username, passhash, email string) error {
	return rep.Db.Create(&models.User{
		Username:     username,
		PasswordHash: passhash,
		Email:        email,
	}).Error
}
