package main

import (
	"github.com/Alfeenn/todo-list/app"
	"github.com/Alfeenn/todo-list/cmd"
	"github.com/Alfeenn/todo-list/controller"
	"github.com/Alfeenn/todo-list/middleware"
	"github.com/Alfeenn/todo-list/repository"
	"github.com/Alfeenn/todo-list/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cmd.MigrateCmd()

	engine := gin.Default()
	db := app.NewDB()
	repo := repository.NewRepository()
	service := service.NewService(repo, db)
	controller := controller.NewController(service)
	engine.Use(gin.CustomRecovery(middleware.ErrorRecovery))
	baseRoute := engine.Group("")
	todo := baseRoute.Group("/todo-items")
	{
		todo.GET("", controller.FindAll)
		todo.GET("/:id", controller.FindTodoById)
		todo.POST("", controller.Create)
		todo.DELETE("/:id", controller.Delete)
		todo.PATCH("/:id", controller.Update)
	}
	activity := baseRoute.Group("/activity-groups")
	{
		activity.GET("", controller.FindAllActivity)
		activity.GET("/:id", controller.FindActivityById)
		activity.POST("", controller.CreateActivity)
		activity.DELETE("/:id", controller.DeleteActivity)
		activity.PATCH("/:idactivity", controller.UpdateActivity)

	}
	engine.Run(":3030")
}
