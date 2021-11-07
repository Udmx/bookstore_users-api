package users

import (
	"github.com/gin-gonic/gin"
	"github.com/udmx/bookstore_users-api/domain/users"
	"github.com/udmx/bookstore_users-api/services"
	"github.com/udmx/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	//TODO: Handle error
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	fmt.Println(err.Error())
	//	//TODO: Handle json error
	//	return
	//}

	// Replace code:
	if err := c.ShouldBindJSON(&user); err != nil {
		//fmt.Println(err.Error())
		// Handle json error => return bad request to the caller.
		//restErr := errors.RestErr{
		//	Message: "invalid json body",
		//	Status:  http.StatusBadRequest,
		//	Error:   "bad_request",
		//}
		restErr:=errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//Handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
