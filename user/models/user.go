package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Username  string
	Phone	string
}
