package main

import (
	"fmt"
	"go-api/controllers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckHeaderForJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("MIDDLEWARE checking for authorization headers")

		authoriazationHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authoriazationHeader, "Bearer ") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "You are not authorized",
			})

			c.Abort()
			return
		}

		token := strings.Split(authoriazationHeader, " ")[1]

		fmt.Println(token)

		user := controllers.ParseToken(token)

		c.Set("user", user)

		c.Next()
	}
}
