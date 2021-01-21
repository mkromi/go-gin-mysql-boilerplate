package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-mysql-boilerplate/Models"
	"go-gin-mysql-boilerplate/Models/Schema"
	"go-gin-mysql-boilerplate/Services"
	"strings"
)

func UserCreate(c *gin.Context) {
	var user Schema.User
	if requestErr := c.ShouldBind(&user); requestErr != nil {
		errRes := strings.Split(requestErr.Error(), ": ")
		Services.ValidationError(c, "These fields are required!", errRes)
		return
	}

	saveErr := Models.UserCreate(&user)
	if saveErr != nil {
		Services.NotAcceptable(c, "Something went wrong!", saveErr)
	} else {
		Services.Success(c, "Registration complete!", user)
	}
}
