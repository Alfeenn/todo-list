package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func Authorize(obj string, act string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get current user/subject
		sub, existed := c.Get("username")
		if !existed {
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "msg": "User hasn't logged in yet"})
			return
		}

		// Load policy from Database
		err := enforcer.LoadPolicy()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"code": 500, "msg": "Failed to load policy from DB"})
			return
		}

		// Casbin enforces policy
		ok, err := enforcer.Enforce(fmt.Sprint(sub), obj, act)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"code": 500, "msg": "Error occurred when authorizing user"})
			return
		}

		if !ok {
			log.Print(fmt.Sprint(sub), obj, act)
			c.AbortWithStatusJSON(403, gin.H{"code": http.StatusForbidden, "msg": "You are not authorized"})
			return
		}
		c.Next()
	}
}
