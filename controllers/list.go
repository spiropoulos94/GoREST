package controllers

import (
	"encoding/json"
	"go-api/models"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateList(c *gin.Context) {
	var list models.List

	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(jsonData, &list)

	list.Name = strings.Trim(list.Name, " ")

	// user, _ := c.Get("user")

	// c.JSON(200, gin.H{
	// 	"user": user,
	// })

	// c.Abort()

	if list.Name == "" {
		c.JSON(400, gin.H{
			"message": "List name is needed",
		})
		return
	}

	// if len(list.Items) == 0 {
	// 	c.JSON(400, gin.H{
	// 		"message": "Cannot Create empty list",
	// 	})
	// 	return
	// }

	result := models.DB.Create(&list)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "List created",
		"data":    list,
	})
}

func FindList(c *gin.Context) {

	id := c.Param("id")

	var list models.List
	result := models.DB.Preload("Items").First(&list, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "List found",
		"data":    list,
	})
}

func GetLists(c *gin.Context) {
	var lists []models.List
	result := models.DB.Preload("Items").Find(&lists)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "lists successfully retrieved",
		"data":    lists,
	})
}

func DeleteList(c *gin.Context) {
	id := c.Param("id")

	queryResult := models.DB.First(&models.List{}, id)

	if queryResult.Error != nil {
		c.JSON(404, gin.H{
			"message": queryResult.Error.Error(),
		})
		return
	}

	result := models.DB.Unscoped().Delete(&models.List{}, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": queryResult.Error.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "List deleted!",
		})
	}
}

func UpdateList(c *gin.Context) {
	// Check if record exists
	id := c.Param("id")
	var list models.List

	queryResult := models.DB.First(&list, id)

	if queryResult.Error != nil {
		c.JSON(404, gin.H{
			"message": queryResult.Error.Error(),
		})
		return
	}

	var newList models.List

	if err := c.ShouldBindJSON(&newList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Updates list properties

	if result := models.DB.Model(&list).Where("id = ?", id).Updates(newList); result.Error != nil {
		c.JSON(404, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	// Deletes current list items to replace them with new ones

	if result := models.DB.Unscoped().Delete(&models.Item{}, "list_id LIKE ?", id); result.Error != nil {
		c.JSON(404, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	//Replaces new items

	if err := models.DB.Model(&list).Association("Items").Append(newList.Items); err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "List Updated",
		"data":    list,
	})

}
