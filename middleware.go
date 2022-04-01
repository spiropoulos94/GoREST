package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckHeaderForJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("checking header authorization")
		c.Next()
	}
}
