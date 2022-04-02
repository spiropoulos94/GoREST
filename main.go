package main

// Kane refactor se models kai conrollers opws edw : https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/

import (
	"go-api/models"
)

func main() {

	makeNewTablesOnInit := true

	models.SetupDatabase()
	models.MakeTables(makeNewTablesOnInit)

	routerStart()

}
