package users

import (
	"microservice_tut/users_api/domain/users"
	"microservice_tut/users_api/services"
	"microservice_tut/users_api/utils/errors"
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
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		rErr := errors.NewBadRequestError("invalid json body")
		c.JSON(rErr.Status, rErr)
		return
	}
	// fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(bytes))
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err)
	// }
	res, e := services.CreateUser(user)
	if e != nil {

		c.JSON(e.Status, e)
		return
	}

	c.JSON(http.StatusCreated, res)
}
