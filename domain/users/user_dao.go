package users

//data access object

import (
	"fmt"
	"github.com/udmx/bookstore_users-api/datasources/mysql/users_db"
	"github.com/udmx/bookstore_users-api/utils/date_utils"
	"github.com/udmx/bookstore_users-api/utils/errors"
	"strings"
)

const (
	indexUniqueEmail = "users_email_uindex"
	errorNoRows      = " no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created) VALUES(?,?,?,?);"
	queryGetUser     = "SELECT id,first_name,last_name,email,date_created FROM users WHERE id=?;"
)

//var (
//	usersDB = make(map[int64]*User)
//)

func (user *User) Get() *errors.RestErr {
	//if err := users_db.Client.Ping(); err != nil {
	//	panic(err) //try to connect
	//}
	//result := usersDB[user.Id]
	//if result == nil {
	//	return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	//}
	//user.Id = result.Id
	//user.FirstName = result.FirstName
	//user.LastName = result.LastName
	//user.Email = result.Email
	//user.DateCreated = result.DateCreated
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//results,_:=stmt.Query(user.Id) // Return rows
	//if err != nil {
	//	return errors.NewInternalServerError(err.Error())
	//}
	//defer results.Close()

	result := stmt.QueryRow(user.Id) //Return single row
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
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

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	//current := usersDB[user.Id]
	//if current != nil {
	//	if current.Email == user.Email {
	//		return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
	//	}
	//	return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	//}
	//
	////now:=time.Now().UTC()
	////user.DateCreated = now.Format()
	//user.DateCreated = date_utils.GetNowString()
	//usersDB[user.Id] = user
	user.Id = userId
	return nil
}
