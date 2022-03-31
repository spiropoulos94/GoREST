package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
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

	// tsekare an yparxei hdh

	result := models.DB.Where("email = ?", "jinzhu").First(&user)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	// an den yparxei user me ayto to mail ftiakse ton

	c.JSON(200, gin.H{"user": user})

}
