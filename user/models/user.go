package user

import (
	"time"
)

type User struct {
	ID              uint `gorm:"primary_key"`
	FirstName       string
	LastName        string
	Username        string
	Type            string
	Email           string
	EmailVerifiedAt time.Time
	RememberToken   string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}
