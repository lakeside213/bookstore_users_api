package users

import (
	"fmt"
	"microservice_tut/users_api/datasources/mysql/usersdb"
	"microservice_tut/users_api/utils/errors"
	"microservice_tut/users_api/utils/mysql_utils"
)

const (
	queryInsertUser       = "INSERT into users(first_name, last_name,email, created_at, status, password) VALUES(?,?,?,?,?, ?)"
	queryGetUser          = "SELECT id, first_name, last_name, email, status, created_at FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id = ?"
	queryDeleteUser       = "DELETE FROM users WHERE id=?"
	queryFindUserByStatus = "SELECT  id, first_name, last_name, email, created_at, status FROM users WHERE status=?;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.UsersDB.Prepare(queryInsertUser)

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt, user.Status, user.Password)

	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)

	}
	userID, err := insertResult.LastInsertId()

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.ID = userID
	return nil
}

func (user *User) Get() *errors.RestErr {
	if err := usersdb.UsersDB.Ping(); err != nil {
		panic(err)
	}
	stmt, err := usersdb.UsersDB.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Status); err != nil {
		fmt.Println(err)
		text := fmt.Sprintf("error when trying to get user %d : %s", user.ID, err.Error())
		return errors.NewInternalServerError(text)
	}

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.UsersDB.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.UsersDB.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.UsersDB.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
