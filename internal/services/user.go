package services

import (
	"cloud_storage/internal/domain"
	"cloud_storage/internal/security"
)

type User struct {
	rep domain.UserRepo
}

func NewUser(rep domain.UserRepo) *User {
	return &User{rep: rep}
}

func (u *User) Register(username, pass, email string) error {

	passhash, err := security.PasswordHash(pass)
	if err != nil {
		return err
	}

	return u.rep.CreateUser(username, passhash, email)
}

func (u *User) Login(username, pass string) error {

	return
}
