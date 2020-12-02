package services

import (
	"microservice_tut/users_api/domain/users"
	"microservice_tut/users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(ID int64) (*users.User, *errors.RestErr) {
	res := &users.User{ID: ID}
	if err := res.Get(); err != nil {
		return nil, err
	}
	return res, nil
}
