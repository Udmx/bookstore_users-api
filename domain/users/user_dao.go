package users

//data access object

import (
	"github.com/udmx/bookstore_users-api/datasources/mysql/users_db"
	"github.com/udmx/bookstore_users-api/utils/date_utils"
	"github.com/udmx/bookstore_users-api/utils/errors"
	"github.com/udmx/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name,last_name,email,date_created) VALUES(?,?,?,?);"
	queryGetUser    = "SELECT id,first_name,last_name,email,date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id) //Return single row
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, _err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if _err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
