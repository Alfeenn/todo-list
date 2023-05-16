package model

import (
	"time"
)

type User struct {
	Id       string `gorm:"primary_key; column:id"`
	Username string `gorm:"not null; unique; size:32"`
	Password string `gorm:"not null; size:70"`
	Name     string `gorm:"not null; size:20"`
	Age      int64  `gorm:"not null; size:20"`
	Phone    int64  `gorm:"not null; size:20"`
	Role     string `gorm:"not null; default:user size:10"`
}

type Todo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Priority  string    `json:"priority"`
	Isactive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type Activity struct {
	Id        int       `gorm:"primary_key; column:id" json:"id"`
	Title     string    `gorm:"not null; size:30" json:"title" binding:"required"`
	Email     string    `gorm:"not null; size:30" json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
