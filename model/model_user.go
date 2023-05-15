package model

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
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Priority  string `json:"priority"`
	Isactive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
}

type Activity struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
