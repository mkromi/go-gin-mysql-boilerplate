package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-mysql-boilerplate/Models"
	"go-gin-mysql-boilerplate/Models/Schema"
	"go-gin-mysql-boilerplate/Services"
	"go-gin-mysql-boilerplate/Validations"
	"strconv"
	"strings"
)

func AuthSignin(c *gin.Context) {
	var request Validations.AuthSignin
	if requestErr := c.ShouldBind(&request); requestErr != nil {
		errRes := strings.Split(requestErr.Error(), ": ")
		Services.ValidationError(c, "These fields are required!", errRes)
		return
	}

	var emailUser Schema.User
	emailUserErr := Models.UserFetchWithEmail(&emailUser, request.Email)
	if emailUserErr != nil {
		Services.NotAcceptable(c, "You do not have an account with this email!", nil)
		return
	}

	if emailUser.Password != Services.MD5Hash(request.Password) {
		Services.NotAcceptable(c, "Invalid email or password!", nil)
		return
	}

	clientName, _, _ := c.Request.BasicAuth()
	tokens := Services.GenerateTokens(strconv.Itoa(int(emailUser.Id)),  clientName)

	Services.Success(c, "Welcome!", gin.H{"userInfo": emailUser, "tokens": tokens})
}

func AuthSignout(c *gin.Context) {
	res := Services.Signout(strings.Split(c.Request.Header.Get("Authorization"), " ")[1])
	if res != true {
		Services.BadRequest(c, "Something went wrong!", nil)
		return
	}
	Services.Deleted(c, "Logout Successful!", nil)
	return
}
