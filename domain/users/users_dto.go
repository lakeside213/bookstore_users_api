package users

import (
	"microservice_tut/users_api/utils/errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
	Password  string `json:"password"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.LastName = strings.TrimSpace(user.LastName)
	// if user.Password == "" {
	// 	return errors.NewBadRequestError("no password")
	// }
	return nil
}
