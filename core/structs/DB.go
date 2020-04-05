package core

import "github.com/jinzhu/gorm"

type DB struct {
	Connections map[string]*gorm.DB
	Credentials map[string]DBConnection
}
