package app

import (
	"microservice_tut/users_api/controllers/ping"
	"microservice_tut/users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:id", users.GetUser)
	// router.GET("/users/search", controllers.FindUser)
	router.POST("/users", users.CreateUser)

}
