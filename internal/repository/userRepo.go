package repository

import (
	"cloud_storage/internal/models"
	"fmt"
	"gorm.io/gorm"
)

type UserRep struct {
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRep {
	return &UserRep{Db: db}
}

func (rep *UserRep) CreateUser(username, passhash, email string) error {
	return rep.Db.Create(&models.User{
		Username:     username,
		PasswordHash: passhash,
		Email:        email,
	}).Error
}

func (rep *UserRep) CheckUser(username string) (bool, error) {
	err := rep.Db.Where("username = ?", username).First(&models.User{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, fmt.Errorf("user not found")
	}

	return true, nil
}

func (rep *UserRep) GetUserId(username string) (int64, error) {
	var user models.User
	err, b := rep.Db.Where("username = ?", username).Get("id")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Id, err, b)
	return user.Id, nil

}
