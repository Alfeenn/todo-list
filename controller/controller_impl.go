package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Alfeenn/todo-list/helper"
	"github.com/Alfeenn/todo-list/model"
	"github.com/Alfeenn/todo-list/model/web"
	"github.com/Alfeenn/todo-list/service"
	"github.com/gin-gonic/gin"
)

type ControllerImpl struct {
	ServiceModel service.Service
}

func NewController(c service.Service) Controller {
	return &ControllerImpl{
		ServiceModel: c,
	}
}

func (c *ControllerImpl) Create(g *gin.Context) {
	request := model.Todo{}
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	err := g.ShouldBindJSON(&request)
	log.Print(request)
	log.Print("error", err)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			web.WebResponse{
				Status:  "Bad Request",
				Message: "title cannot be null",
			})
	} else {

		resp := c.ServiceModel.CreateToDo(g.Request.Context(), request)
		response := web.WebResponse{
			Status:  "Success",
			Message: "Success",
			Data:    resp,
		}
		g.JSON(201, &response)
	}
}

func (c *ControllerImpl) Update(g *gin.Context) {
	request := model.Todo{}
	err := g.ShouldBindJSON(&request)
	log.Print(request)
	//check if bind json error
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			web.WebResponse{
				Status:  "Bad Request",
				Message: "title cannot be null",
			})
	} else {
		stringId := g.Param("id")
		id, err := strconv.Atoi(stringId)
		helper.PanicIfErr(err)
		request.Id = id
		log.Print(request)
		result := c.ServiceModel.UpdateToDo(g.Request.Context(), request)
		response := web.WebResponse{
			Status:  "Success",
			Message: "Success",
			Data:    result,
		}
		g.JSON(http.StatusOK, response)
	}

}

func (c *ControllerImpl) Delete(g *gin.Context) {
	stringId := g.Params.ByName("id")
	id, err := strconv.Atoi(stringId)
	helper.PanicIfErr(err)
	c.ServiceModel.DeleteToDo(g.Request.Context(), id)
	response := web.WebResponse{
		Status:  "Success",
		Message: "Success",
	}
	g.JSON(http.StatusOK, response)
}

func (c *ControllerImpl) FindTodoById(g *gin.Context) {
	stringId := g.Params.ByName("id")
	id, err := strconv.Atoi(stringId)
	helper.PanicIfErr(err)

	result := c.ServiceModel.FindTodo(g.Request.Context(), id)

	response := web.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	}
	g.JSON(http.StatusOK, response)
}

func (c *ControllerImpl) FindCourseByCategory(g *gin.Context) {
	id := g.Params.ByName("category")
	log.Print(id)
	if id == "" {
		g.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{
				"code": http.StatusNotFound,
				"msg":  "category not found"})
	} else {

		// result := c.ServiceModel.FindActivityById(g.Request.Context(), id)
		response := web.WebResponse{
			Status:  "Success",
			Message: "Success",
			// Data:   result,
		}
		g.JSON(http.StatusOK, response)

	}

}

func (c *ControllerImpl) FindAll(g *gin.Context) {
	id := g.Query("activity_group_id")
	if id == "" {
		result := c.ServiceModel.FindAllToDo(g.Request.Context())
		response := web.WebResponse{
			Status:  "Success",
			Message: "Success",
			Data:    result,
		}
		g.JSON(http.StatusOK, response)
	} else {
		intId, _ := strconv.Atoi(id)
		result := c.ServiceModel.FindTodo(g.Request.Context(), intId)
		response := web.WebResponse{
			Status:  "Success",
			Message: "Success",
			Data:    result,
		}
		g.JSON(http.StatusOK, response)
	}

}

func (c *ControllerImpl) CreateActivity(g *gin.Context) {
	request := model.Activity{}
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	err := g.ShouldBindJSON(&request)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			web.WebResponse{
				Status:  "Bad Request",
				Message: "title cannot be null",
			})
	} else {

		resp := c.ServiceModel.CreateActivity(g.Request.Context(), request)
		response := web.WebResponse{
			Status:  "Success",
			Message: "Success",
			Data:    resp,
		}
		g.JSON(201, response)
	}
}

func (c *ControllerImpl) FindAllActivity(g *gin.Context) {

	result := c.ServiceModel.FindAllActivity(g.Request.Context())
	response := web.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	}
	g.JSON(http.StatusOK, response)

}

func (c *ControllerImpl) UpdateActivity(g *gin.Context) {
	request := model.Activity{}
	request.UpdatedAt = time.Now()
	err := g.ShouldBindJSON(&request)
	//check if bind json error
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			web.WebResponse{
				Status:  "Bad Request",
				Message: "title cannot be null",
			})
	} else {
		stringId := g.Param("idactivity")
		id, err := strconv.Atoi(stringId)
		helper.PanicIfErr(err)
		request.Id = id
		log.Print(request)
		result := c.ServiceModel.UpdateActivity(g.Request.Context(), request)
		response := web.WebResponse{
			Status:  "Success",
			Message: "Success",
			Data:    result,
		}
		g.JSON(http.StatusOK, response)
	}

}

func (c *ControllerImpl) FindActivityById(g *gin.Context) {
	stringId := g.Params.ByName("id")
	id, err := strconv.Atoi(stringId)
	helper.PanicIfErr(err)

	result := c.ServiceModel.FindActivityById(g.Request.Context(), id)

	response := web.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	}
	g.JSON(http.StatusOK, response)

}

func (c *ControllerImpl) DeleteActivity(g *gin.Context) {
	stringId := g.Params.ByName("id")
	id, err := strconv.Atoi(stringId)
	helper.PanicIfErr(err)
	c.ServiceModel.DeleteActivity(g.Request.Context(), id)
	response := web.WebResponse{
		Status:  "Success",
		Message: "Success",
	}
	g.JSON(http.StatusOK, response)
}
