// app/models/user.go
package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (User) TableName() string {
	return "users"
}
