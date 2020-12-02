package users

import (
	"microservice_tut/users_api/utils/date"
	"microservice_tut/users_api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError("email already exists")
		}
		return errors.NewBadRequestError("user already exists")
	}
	user.CreatedAt = date.GetNowString()
	usersDB[user.ID] = user
	return nil
}

func (user *User) Get() *errors.RestErr {
	res := usersDB[user.ID]
	if res == nil {
		return errors.NewBadRequestError("user not found")
	}

	user.ID = res.ID
	user.FirstName = res.FirstName
	user.LastName = res.LastName
	user.Email = res.Email
	return nil
}
