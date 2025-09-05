package repository

import "gorm.io/gorm"

type UserRep struct {
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRep {
	return &UserRep{Db: db}
}

func (rep *UserRep) Register(username, passhash, email string) error {

}
