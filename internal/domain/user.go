package domain

type UserRepo interface {
	CreateUser(username, passhash, email string) error
	CheckUser(username string) (bool, error)
	GetUserId(username string) (int64, error)
}
