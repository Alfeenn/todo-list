package middleware

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Alfeenn/todo-list/helper"
	"github.com/Alfeenn/todo-list/model"
	"github.com/Alfeenn/todo-list/model/web"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := strconv.AppendBool([]byte(model.Key), true)
		claim := helper.ClaimToken(ctx, key)
		if claim.Username == "" {
			response := web.WebResponse{Code: http.StatusUnauthorized, Status: "UNAUTHORIZED"}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				response)
		}
		log.Print(gin.H{"data": claim.Username})
	}
}

func NotFound(c *gin.Context, err error) {
	response := web.WebResponse{
		Status:  "error",
		Message: err.Error(),
	}
	c.AbortWithStatusJSON(http.StatusNotFound, response)
}

func BadRequest(c *gin.Context, err error) {
	response := web.WebResponse{
		Status:  "error",
		Message: "title cannot be null",
	}
	c.AbortWithStatusJSON(http.StatusNotFound, response)
}
