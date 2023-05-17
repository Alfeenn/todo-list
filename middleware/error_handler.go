package middleware

import (
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
			Status:  "error",
			Message: exception.Err,
		})
		return true
	} else {
		return false
	}
}

func InternalServer(g *gin.Context, err interface{}) {
	g.AbortWithStatusJSON(500, web.WebResponse{
		Status: "Internal Server Error",
		Data:   err,
	})
}
