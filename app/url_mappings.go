package app

import (
	"github.com/MoulikKhamaru/bookstore_users-api/controllers/ping"
	"github.com/MoulikKhamaru/bookstore_users-api/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", user.CreateUser)

	router.GET("/users:id", user.GetUser)
}
