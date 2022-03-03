package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type List struct {
	gorm.Model
	Name      string
	ListItems []ListItem
}

type ListItem struct {
	gorm.Model
	Name   string
	ListID uint
}

func SetupDatabase() {
	godotenv.Load(".env")

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		// fmt.Println("ERROR while connecting to database :", err)
		log.Fatal()
	}

	DB = db
}

func MakeTables(allNew bool) {

	if allNew {
		DB.Migrator().DropTable(&List{})
		DB.Migrator().DropTable("lists")
		DB.Migrator().DropTable(&ListItem{})
		DB.Migrator().DropTable("list_items")
	}

	DB.AutoMigrate(&List{})
	DB.AutoMigrate(&ListItem{})
	DB.AutoMigrate(&User{})

}
