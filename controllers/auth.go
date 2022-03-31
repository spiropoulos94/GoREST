package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
	"go-api/utils"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// newToken encodes User struct into a JWT string

type Claims struct {
	models.User
	jwt.StandardClaims
}

// func newToken(user models.User) string {

// }

//  parse token reads a jwt token and returns a models.User struct

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
		user.Password, _ = utils.HashPassword(user.Password)
		models.DB.Create(&user)

		c.JSON(201, gin.H{
			"message": "user successfully created",
			"user":    user,
			// "token":   token,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "user already exists",
		})
	}

}
