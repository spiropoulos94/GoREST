package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
	"go-api/utils"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// sto signup tha ftiaxneis user kai tha tou dineis pisw ena JWT token

func Signup(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(jsonData)
	}

	user := models.User{}

	json.Unmarshal(jsonData, &user)

	userExists := utils.UserExists(user.Email)

	fmt.Println("userExists")
	fmt.Println(userExists)

	if !userExists {
		models.DB.Create(&user)
		c.JSON(200, gin.H{
			"message": "user successfully created",
			"user":    user,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "user already exists",
		})
	}

}
