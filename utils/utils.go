package utils

import (
	"go-api/models"

	"golang.org/x/crypto/bcrypt"
)

func UserExists(mail string) bool {
	var exists bool

	models.DB.Model(models.User{}).
		Select("count(*) > 0").
		Where("email = ?", mail).
		Find(&exists)

	return exists

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
