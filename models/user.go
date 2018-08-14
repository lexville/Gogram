package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User model contains the id, created at, deleted at
// name and email
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
