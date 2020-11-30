package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetUser gets a user
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "you got work to do")
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "you got work to do")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "you got work to do")
}
