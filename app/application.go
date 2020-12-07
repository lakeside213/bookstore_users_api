package app

import (
	"microservice_tut/users_api/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapUrls()
	logger.Info("starting app")
	router.Run(":8080")
}
