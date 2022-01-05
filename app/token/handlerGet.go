package token

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func getGet(c *gin.Context) {
	tokenStr := c.GetHeader("Auth")
	if tokenStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "not have auth",
		})
		c.Abort()
		return
	}

	token, claims, err := ParseToken(tokenStr)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "not have auth",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, claims)
}

func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
