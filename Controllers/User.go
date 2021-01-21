package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-mysql-boilerplate/Models"
	"go-gin-mysql-boilerplate/Services"
	"go-gin-mysql-boilerplate/Validations"
	"strings"
)

func UserCreate(c *gin.Context) {
	var request Validations.UserCreate
	if requestErr := c.ShouldBind(&request); requestErr != nil {
		errRes := strings.Split(requestErr.Error(), ": ")
		Services.ValidationError(c, "These fields are required!", errRes)
		return
	}

	saveErr := Models.UserCreate(&request)
	if saveErr != nil {
		Services.NotAcceptable(c, "Something went wrong!", saveErr)
	} else {
		Services.Success(c, "Created", request)
	}
}

func UserUpdate(c *gin.Context) {
	userId := c.Param("id")

	var request Validations.UserUpdate
	if requestErr := c.ShouldBind(&request); requestErr != nil {
		errRes := strings.Split(requestErr.Error(), ": ")
		Services.ValidationError(c, "These fields are required!", errRes)
		return
	}

	updateErr := Models.UserUpdate(&request, userId)
	if updateErr != nil {
		Services.NotAcceptable(c, "Something went wrong!", updateErr)
	} else {
		Services.Success(c, "Updated", request)
	}
}
