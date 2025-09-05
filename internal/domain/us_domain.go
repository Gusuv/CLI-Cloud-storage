package domain

import "cloud_storage/internal/repository"

type UserRepo interface {
	repository.UserRep
}
