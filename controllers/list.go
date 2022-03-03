package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func CreateList(c *gin.Context) {
	var list models.List
	// var items []models.Item

	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(jsonData, &list)

	fmt.Println("list")
	fmt.Println(list)
}
