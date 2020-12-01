package services

import (
	"microservice_tut/users_api/domain/users"
	"microservice_tut/users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
