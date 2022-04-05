package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MoulikKhamaru/bookstore_users-api/domain/users"
	"github.com/MoulikKhamaru/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Impliment this !")
}

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println("before Unmarshal", user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//Do something
		return
	}
	if error := json.Unmarshal(bytes, user); error != nil {
		fmt.Println(error.Error())
		return
	}
	result, saveerror := services.CreateUser(user)
	if saveerror != nil {
		//TODO: handles error
		return
	}
	fmt.Println("after unmarshal", user)
	c.String(http.StatusNotImplemented, "Implement createuser!")

}
