package services

import (
	"microservice_tut/users_api/domain/users"
	"microservice_tut/users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
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

func UpdateUser(isPatch bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPatch {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(ID int64) *errors.RestErr {
	user := &users.User{ID: ID}
	return user.Delete()
}

func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)

}
