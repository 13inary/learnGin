package token

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("hello token")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func getSet(c *gin.Context) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: 6,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Printf("token = %+v\n", token)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("sign token error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
