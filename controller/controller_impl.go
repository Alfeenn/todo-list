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
	req, err := helper.UploadFile(g)

	if err != nil {

		g.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
	} else {

		resp := c.ServiceModel.CreateToDo(g.Request.Context(), req)
		response := web.WebResponse{
			Code:   http.StatusCreated,
			Status: "CREATED",
			Data:   resp,
		}
		g.JSON(http.StatusOK, response)
	}
}

func (c *ControllerImpl) Update(g *gin.Context) {
	req, err := helper.UploadFile(g)
	//check if bind json error
	if err != nil {

		g.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
	} else {
		result := c.ServiceModel.UpdateToDo(g.Request.Context(), req)
		g.JSON(http.StatusOK, result)
	}

}

func (c *ControllerImpl) Delete(g *gin.Context) {
	id := g.Param("id")
	c.ServiceModel.DeleteToDo(g.Request.Context(), id)
	g.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "Successfully delete data"})
}

func (c *ControllerImpl) DeleteActivity(g *gin.Context) {
	stringId := g.Params.ByName("id")
	id, err := strconv.Atoi(stringId)
	helper.PanicIfErr(err)
	c.ServiceModel.DeleteActivity(g.Request.Context(), id)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}
	g.JSON(http.StatusOK, response)
}

func (c *ControllerImpl) FindCourseById(g *gin.Context) {
	id := g.Params.ByName("id")
	if id == "" {
		g.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{
				"code": http.StatusNotFound,
				"msg":  "Id not found"})
	} else {
		result := c.ServiceModel.FindTodo(g.Request.Context(), id)
		response := web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   result,
		}
		g.JSON(http.StatusOK, response)
	}

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
			Code:   http.StatusOK,
			Status: "OK",
			// Data:   result,
		}
		g.JSON(http.StatusOK, response)

	}

}

func (c *ControllerImpl) FindAll(g *gin.Context) {

	result := c.ServiceModel.FindAllToDo(g.Request.Context())
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	g.JSON(http.StatusOK, response)

}

func (c *ControllerImpl) CreateActivity(g *gin.Context) {
	request := model.Activity{}
	request.CreatedAt = time.Now()
	err := g.ShouldBindJSON(&request)
	if err != nil {

		g.AbortWithStatusJSON(http.StatusBadRequest,
			web.WebResponse{
				Status:  "Bad request",
				Message: "title cannot be null",
			})
	} else {

		resp := c.ServiceModel.CreateActivity(g.Request.Context(), request)
		response := web.WebResponse{
			Code:   http.StatusCreated,
			Status: "CREATED",
			Data:   resp,
		}
		g.JSON(http.StatusOK, response)
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
	err := g.ShouldBindJSON(&request)

	//check if bind json error
	if err != nil {

		g.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
	} else {
		stringId := g.Param("idactivity")
		id, err := strconv.Atoi(stringId)
		helper.PanicIfErr(err)
		request.Id = id
		log.Print(request)
		result := c.ServiceModel.UpdateActivity(g.Request.Context(), request)
		g.JSON(http.StatusOK, result)
	}

}

func (c *ControllerImpl) FindActivityById(g *gin.Context) {
	stringId := g.Params.ByName("id")
	id, err := strconv.Atoi(stringId)
	helper.PanicIfErr(err)
	if id == 0 {
		g.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{
				"code": http.StatusNotFound,
				"msg":  "Id not found"})
	} else {
		result := c.ServiceModel.FindActivityById(g.Request.Context(), id)
		response := web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   result,
		}
		g.JSON(http.StatusOK, response)
	}

}
