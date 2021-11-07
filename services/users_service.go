package services

import (
	"github.com/udmx/bookstore_users-api/domain/users"
	"github.com/udmx/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
