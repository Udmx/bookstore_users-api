package users

//data access object

import (
	"fmt"
	"github.com/udmx/bookstore_users-api/datasources/mysql/users_db"
	"github.com/udmx/bookstore_users-api/utils/date_utils"
	"github.com/udmx/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err) //try to connect
	}

	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	//now:=time.Now().UTC()
	//user.DateCreated = now.Format()
	user.DateCreated = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
