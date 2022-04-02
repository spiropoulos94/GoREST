package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
	"go-api/utils"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

func NewToken(user models.User) (string, error) {

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

func ParseToken(token_string string, c *gin.Context) (*Claims, error) {

	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not validate token",
		})

		c.Abort()
		return nil, nil
	}

	return claims, nil
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

	if !userExists {
		user.Password, _ = utils.HashPassword(strings.TrimSpace(user.Password))
		models.DB.Create(&user)

		token, _ := NewToken(user)

		c.JSON(201, gin.H{
			"message": "user successfully created",
			"user":    user,
			"jwt":     token,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "user already exists",
		})
	}

}

func Signin(c *gin.Context) {

	storedUser := models.User{}
	reqBodyData := models.User{}

	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(jsonData, &reqBodyData)

	email := strings.TrimSpace(reqBodyData.Email)
	password := strings.TrimSpace(reqBodyData.Password)

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "need email and password",
		})
		return
	}

	userExists := utils.UserExists(reqBodyData.Email)

	if !userExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user does not exist",
		})
		return
	} else {
		// sygkrine to password sto request me to pass tou user
		models.DB.Where("email = ?", reqBodyData.Email).First(&storedUser)

		if utils.CheckPasswordHash(reqBodyData.Password, storedUser.Password) {
			token, _ := NewToken(storedUser)
			c.JSON(200, gin.H{
				"message": "password match",
				"jwt":     token,
			})
			return
		} else {
			c.JSON(200, gin.H{
				"message": "password doesnt match",
			})
			return
		}

	}

}
