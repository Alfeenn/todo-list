package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Alfeenn/todo-list/helper"
	"github.com/Alfeenn/todo-list/middleware"
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

func (c *ControllerImpl) DeleteUser(g *gin.Context) {
	id := g.Param("id")
	c.ServiceModel.DeleteUser(g.Request.Context(), id)
	g.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "Successfully delete data"})
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

		result := c.ServiceModel.FindCourseByCategory(g.Request.Context(), id)
		response := web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   result,
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

func (c *ControllerImpl) UserSignIn(g *gin.Context) {
	key := strconv.AppendBool([]byte(model.Key), true)
	requestservice := web.RequestLogin{
		Username: g.Request.FormValue("username"),
		Password: g.Request.FormValue("password"),
	}
	//check form input
	err := g.ShouldBind(&requestservice)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
	} else {
		//proceed to login
		user := c.ServiceModel.Login(g.Request.Context(), requestservice)
		match := helper.CheckHashPassword(user.Password, requestservice.Password)
		var data map[string]interface{}
		//check if password match
		if !match {
			g.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code": http.StatusInternalServerError,
					"msg":  "Password not match"})
		} else {
			//set token
			tokenstring := helper.GenerateToken(g, key, user)
			data = map[string]interface{}{
				"Authorization": tokenstring,
			}
			g.JSON(http.StatusOK, web.WebResponse{
				Code:   200,
				Status: "OK",
				Data:   data,
			})
		}
	}
}

func (c *ControllerImpl) GetAccessList(g *gin.Context) {
	enforcer := middleware.UserPolicy()

	adapter := enforcer.GetAllObjects()

	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   adapter,
	}
	g.JSON(http.StatusOK, response)

}

func (c *ControllerImpl) Register(g *gin.Context) {
	enforcer := middleware.UserPolicy()
	req := web.CategoryRequest{}
	age, _ := strconv.Atoi(g.Request.FormValue("age"))
	req.Age = int64(age)
	phone, _ := strconv.Atoi(g.Request.FormValue("phone"))
	req.Phone = int64(phone)

	err := g.ShouldBind(&req)
	log.Print(req)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
	} else {

		req.Password, _ = helper.HashPassword(req.Password)
		resp := c.ServiceModel.Register(g.Request.Context(), req)
		enforcer.AddGroupingPolicy(fmt.Sprint(resp.Username), resp.Role)
		response := web.WebResponse{
			Code:   http.StatusCreated,
			Status: "CREATED",
			Data:   resp,
		}
		g.JSON(http.StatusOK, response)
	}

}

func (c *ControllerImpl) GetCourse(g *gin.Context) {
	id := g.Params.ByName("idcourse")
	req := model.Class{}
	sub, existed := g.Get("id")
	if !existed {
		g.AbortWithStatusJSON(401, gin.H{"code": 401, "msg": "User hasn't logged in yet"})
		return
	}
	req.UserId = fmt.Sprint(sub)
	log.Print(req.UserId)
	if id == "" {
		g.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{
				"code": http.StatusNotFound,
				"msg":  "Id not found"})
	} else {
		result := c.ServiceModel.GetCourse(g.Request.Context(), req, id)
		response := web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   result,
		}
		g.JSON(http.StatusOK, response)
	}

}
