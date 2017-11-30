package entities

import (
	"time"
)

var id uint64 = 1

// User
type User struct {
	UID        uint64    `gorm:"primary_key"`
	Username   string    `gorm:"not null"`
	Password   string    `gorm:"not null"`
	CreateTime time.Time `gorm:"not null"`
}

// NewUser returns a new user with a new uid
func NewUser(username, password string) *User {
	u := User{
		UID:        id,
		Username:   username,
		Password:   password,
		CreateTime: time.Now(),
	}
	id++
	return &u
}
