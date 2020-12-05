package users

import (
	"microservice_tut/users_api/domain/users"
	"microservice_tut/users_api/services"
	"microservice_tut/users_api/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserID(userID string) (int64, *errors.RestErr) {
	ID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}
	return ID, nil
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "you got work to do")
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		rErr := errors.NewBadRequestError("invalid json body")
		c.JSON(rErr.Status, rErr)
		return
	}
	res, e := services.CreateUser(user)
	if e != nil {

		c.JSON(e.Status, e)
		return
	}

	c.JSON(http.StatusCreated, res)
}
func GetUser(c *gin.Context) {
	ID, err := getUserID(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(ID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
func UpdateUser(c *gin.Context) {
	ID, err := getUserID(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		rErr := errors.NewBadRequestError("invalid json body")
		c.JSON(rErr.Status, rErr)
		return
	}
	user.ID = ID
	isPatch := c.Request.Method == http.MethodPatch
	res, e := services.UpdateUser(isPatch, user)
	if e != nil {

		c.JSON(e.Status, e)
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteUser(c *gin.Context) {
	ID, err := getUserID(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if err := services.DeleteUser(ID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Search(status)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users)
}
