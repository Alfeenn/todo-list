package middleware

import (
	"log"

	"github.com/Alfeenn/todo-list/model/web"
	"github.com/gin-gonic/gin"
)

type ErrNotFound struct {
	Err string
}

func NewNotFound(error string) ErrNotFound {
	return ErrNotFound{
		Err: error,
	}
}

func ErrorRecovery(g *gin.Context, err interface{}) {
	if NotFoundError(g, err) {
		return
	}
	InternalServer(g, err)
}

func NotFoundError(g *gin.Context, err interface{}) bool {
	exception, ok := err.(ErrNotFound)

	if ok {
		g.AbortWithStatusJSON(404, web.WebResponse{
			Status:  "Not Found",
			Message: exception.Err,
		})
		return true
	} else {
		g.Next()
		return false
	}
}

func InternalServer(g *gin.Context, err interface{}) {
	log.Print("err middleware", err)
	g.AbortWithStatusJSON(500, web.WebResponse{
		Status: "Bad Request",
		Data:   &err,
	})
	g.Next()
}
