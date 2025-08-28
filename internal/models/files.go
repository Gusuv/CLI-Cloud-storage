package models

import "time"

type File struct {
	Id        int64     `gorm:"primary_key"`
	UserId    int64     `gorm:"not null;index"`
	User      User      `gorm:"foreignKey:UserId"`
	Name      string    `gorm:"not null"`
	Path      string    `gorm:"not null"`
	Size      int64     `gorm:"not null"`
	MimeType  string    `gorm:"size:100"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
