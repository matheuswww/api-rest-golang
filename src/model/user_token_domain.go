package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/rest_err"
)

var(
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string,*rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)
	
	claims := jwt.MapClaims{
		"id": ud.id,
		"email": ud.email,
		"name": ud.name,
		"age": ud.age,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512,claims)
	tokenString,err := token.SignedString([]byte(secret))
	if err != nil {
		return "",rest_err.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token,err=%s",err),
		)
	}

	return tokenString,nil
}

func VerifyTokenMiddleware(c *gin.Context)  {
	secret := os.Getenv(JWT_SECRET_KEY)
	tokenValue := removeBearerPrefix(c.Request.Header.Get("Authorization"))

	token,err := jwt.Parse(removeBearerPrefix(tokenValue),func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC);ok {
			return []byte(secret),nil
		}
		errRest := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(errRest.Code,errRest)
		c.Abort()
		return nil,errRest
	})
	if err != nil {
		errRest := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(errRest.Code,errRest)
		c.Abort()
		return
	}
	claims,ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(errRest.Code,errRest)
		c.Abort()
		return
	}
	userDomain := &userDomain{
		id: uint(claims["id"].(float64)),
		email: claims["email"].(string),
		name: claims["name"].(string),
		age: uint8(claims["age"].(float64)),
	}
	logger.Info(fmt.Sprintf("User authenticated: %#v",userDomain))
}

func removeBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer") {
		token = strings.TrimPrefix("Bearer",token)
	}
	return token
}