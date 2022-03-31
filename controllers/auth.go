package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(jsonData)
	}

	user := models.User{}

	json.Unmarshal(jsonData, &user)

	c.JSON(200, gin.H{"user": user})

}
