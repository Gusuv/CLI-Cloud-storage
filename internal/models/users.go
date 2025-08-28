package models

import "time"

type User struct {
	Id           int64     `gorm:"primary_key"`
	Username     string    `gorm:"uniqueIndex;not null"`
	Email        string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	Created      time.Time `gorm:"autoCreateTime"`
}
