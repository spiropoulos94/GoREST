package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckHeaderForJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("MIDDLEWARE LOGGGGINNNNNNG")
		c.Next()
	}
}
