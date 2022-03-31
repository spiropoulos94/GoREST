package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
	"go-api/utils"
	"io/ioutil"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// newToken encodes User struct into a JWT string

type Claims struct {
	models.User
	jwt.StandardClaims
}

func newToken(user models.User) (string, error) {

	expirationTime := time.Now().Add(120 * time.Hour) // 5 days

	claims := &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err
	}

	return tokenString, nil
}

func parseToken(token_string string) *Claims {

	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	return claims
}

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

		token, _ := newToken(user)

		parseToken(token)

		// 	c.JSON(201, gin.H{
		// 		"message": "user successfully created",
		// 		"user":    user,
		// 		"token":   token,
		// 	})
		// } else {
		c.JSON(400, gin.H{
			"message": "user already exists",
		})
	}

}
