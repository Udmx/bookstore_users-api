package app

import (
	"github.com/udmx/bookstore_users-api/controllers/ping"
	"github.com/udmx/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id",users.UpdateUser)
}
