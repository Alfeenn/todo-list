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

func NewMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("X-API-KEY") == "RAHASIA" {
			return
		}

		ctx.Next()
	}
}

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
