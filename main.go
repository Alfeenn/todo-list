package main

import (
	"github.com/Alfeenn/todo-list/app"
	"github.com/Alfeenn/todo-list/cmd"
	"github.com/Alfeenn/todo-list/controller"
	"github.com/Alfeenn/todo-list/repository"
	"github.com/Alfeenn/todo-list/service"
	"github.com/gin-gonic/gin"
)

func main() {
	migrate := cmd.MigrateCmd()
	if migrate {
		return
	}
	engine := gin.New()
	db := app.NewDB()
	repo := repository.NewRepository()
	service := service.NewService(repo, db)
	controller := controller.NewController(service)
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	baseRoute := engine.Group("")
	admin := baseRoute.Group("/admin")
	{
		admin.GET("/course", controller.FindAll)
		admin.POST("/course", controller.Create)
	}
	activity := baseRoute.Group("/activity-groups")
	{
		activity.GET("", controller.FindAllActivity)
		activity.GET("/:id", controller.FindActivityById)
		activity.POST("", controller.CreateActivity)
		activity.DELETE("/:id", controller.DeleteActivity)
		activity.PATCH("/:idactivity", controller.UpdateActivity)
		activity.GET("/class/:id", controller.FindCourseById)

	}
	engine.Run("localhost:8000")
}
