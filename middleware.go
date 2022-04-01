package main

import (
	"fmt"
	"go-api/controllers"
	"go-api/models"
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

		user := controllers.ParseToken(token)

		dbStoredUser := models.User{}

		result := models.DB.Select("id", "name", "email", "age").First(&dbStoredUser, user.Id)

		if result.Error != nil {
			fmt.Println("Errow while getting store user from DB")
			c.JSON(404, gin.H{
				"error":   result.Error.Error(),
				"message": "could not find jwt stored user in database",
			})
			c.Abort()
			return
		}

		c.Set("user", &dbStoredUser)

		c.Next()
	}
}
