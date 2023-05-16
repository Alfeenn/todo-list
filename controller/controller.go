package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
	FindCourseByCategory(g *gin.Context)
	FindCourseById(g *gin.Context)
	FindAll(g *gin.Context)
	CreateActivity(g *gin.Context)
	UpdateActivity(g *gin.Context)
	DeleteActivity(g *gin.Context)
	FindAllActivity(g *gin.Context)
	FindActivityById(g *gin.Context)
}
