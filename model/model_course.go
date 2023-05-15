package model

import "mime/multipart"

type Course struct {
	Id        string                `form:"id" gorm:"" json:"id"`
	Name      string                `form:"name" gorm:"not null; size:20" json:"name" binding:"required"`
	Price     int                   `form:"price" gorm:"not null; size:50" json:"price"`
	Category  string                `form:"category" gorm:"not null; size:20" json:"category"`
	File      *multipart.FileHeader `form:"file"  json:"file" binding:"required"`
	Thumbnail string                `form:"thumbnail" gorm:"not null; size:50" json:"thumbnail"`
}

type Class struct {
	UserId   string `form:"userid" gorm:"size:40" json:"user_id"`
	Users    User   `gorm:"foreignKey:UserId" json:"-"`
	CourseId string `gorm:"size:40" json:"course_id"`
	Courses  Course `gorm:"foreignKey:CourseId" json:"-"`
}
