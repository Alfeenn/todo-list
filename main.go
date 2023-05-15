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
	migrate, enforcer := cmd.MigrateCmd()
	if migrate {
		return
	}
	engine := gin.New()
	db := app.NewDB()
	repo := repository.NewRepository()
	service := service.NewService(repo, db)
	controller := controller.NewController(service)
	auth := middleware.NewMiddleware()
	engine.Use(gin.Logger())
	baseRoute := engine.Group("/api", auth)
	{
		baseRoute.POST("/register", controller.Register)
		baseRoute.GET("/login", controller.UserSignIn)
	}
	admin := baseRoute.Group("/admin", middleware.AuthJWT())
	{
		admin.GET("/course", middleware.Authorize("course", "read", enforcer), controller.FindAll)
		admin.POST("/course", middleware.Authorize("course", "write", enforcer), controller.Create)
		admin.GET("/course/:id", middleware.Authorize("course", "read", enforcer), controller.FindCourseById)
		admin.GET("/course/category/:category", middleware.Authorize("course", "read", enforcer), controller.FindCourseByCategory)
		admin.PUT("/:id", middleware.Authorize("course", "write", enforcer), controller.Update)
		admin.POST("/course/:id", middleware.Authorize("course", "delete", enforcer), controller.Delete)
		admin.POST("/:id", middleware.Authorize("course", "delete", enforcer), controller.DeleteUser)
		admin.GET("/acl", middleware.Authorize("course", "read", enforcer), controller.GetAccessList)
	}
	user := baseRoute.Group("/user", middleware.AuthJWT())
	{
		user.GET("/course", middleware.Authorize("course", "read", enforcer), controller.FindAll)
		user.POST("/class/:idcourse", middleware.Authorize("class", "write", enforcer), controller.GetCourse)
		user.GET("/course/category/:category", middleware.Authorize("course", "read", enforcer), controller.FindCourseByCategory)
		user.GET("/class/:id", middleware.Authorize("class", "write", enforcer), controller.FindCourseById)
		user.GET("/acl", middleware.Authorize("course", "read", enforcer), controller.GetAccessList)
	}
	engine.Run("localhost:8000")
}
