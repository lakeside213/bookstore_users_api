package main

import (
	"fmt"
	"microservice_tut/users_api/app"
	"microservice_tut/users_api/controllers"
)

func main() {
	fmt.Println("hey")
	app.StartApp()
	controllers.Hez()
}
