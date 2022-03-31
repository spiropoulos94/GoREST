package utils

import (
	"go-api/models"
)

func UserExists(mail string) bool {
	var exists bool

	models.DB.Model(models.User{}).
		Select("count(*) > 0").
		Where("email = ?", mail).
		Find(&exists)

	return exists

}
