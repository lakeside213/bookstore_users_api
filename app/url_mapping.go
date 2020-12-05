package app

import (
	"microservice_tut/users_api/controllers/ping"
	"microservice_tut/users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/user/:id", users.GetUser)
	router.PUT("/user/:id", users.UpdateUser)
	router.PATCH("/user/:id", users.UpdateUser)
	router.DELETE("/user/:id", users.DeleteUser)
	router.GET("/internal/users/search", users.Search)
}
