package services

import (
	"cloud_storage/internal/domain"
	"cloud_storage/internal/sec"
)

type User struct {
	rep domain.UserRepo
}

func NewUser(rep domain.UserRepo) *User {
	return &User{rep: rep}
}

func (u *User) Register(username, pass, email string) error {

	passhash, err := sec.PasswordHash(pass)
	if err != nil {
		return err
	}

	return u.rep.Register(username, passhash, email)
}
