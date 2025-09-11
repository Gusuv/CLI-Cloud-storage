package domain

type UserRepo interface {
	Register(username, passhash, email string) error
}
