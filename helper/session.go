package helper

import (
	"net/http"
	"strings"
	"time"

	"github.com/Alfeenn/todo-list/model"
	"github.com/Alfeenn/todo-list/model/web"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(c *gin.Context, key interface{}, user web.CatResp) model.Session {
	timeT := time.Now().Add(10 * time.Minute)

	claims := &model.Token{
		Id:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(timeT),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(key)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		panic(err)
	}
	return model.Session{
		Id:       claims.Id,
		Username: claims.Username,
		Expiry:   timeT,
		Token:    tokenString,
	}
}

func ClaimToken(c *gin.Context, key interface{}) *model.Token {
	var authorizationToken string
	bearerToken := c.GetHeader("Authorization")
	if strings.HasPrefix(bearerToken, "Bearer ") {
		if bearerToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
		authorizationToken = strings.TrimPrefix(bearerToken, "Bearer ")
	}
	// Get the JWT string from the cookie
	tknStr := authorizationToken
	// Initialize a new instance of `Claims`
	claims := &model.Token{}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {

		return key, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return &model.Token{}
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, []byte("invaild token"))
		return &model.Token{}
	}
	if !tkn.Valid {

		c.AbortWithStatus(http.StatusUnauthorized)
		return &model.Token{}
	} else {
		c.Set("id", claims.Id)
		c.Set("username", claims.Username)
	}
	// Finally, return the welcome message to the user, along with their
	// username given in the token
	// w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username))

	return claims
}
