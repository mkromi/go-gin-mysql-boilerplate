package Routes
import (
	"github.com/gin-gonic/gin"
	"go-gin-mysql-boilerplate/Controllers"
)

func SetupRouter() *gin.Engine {
	route := gin.Default()

	auth := route.Group("/auth")
	{
		auth.POST("signin", Controllers.AuthSignin)
	}

	route.GET("users", Controllers.UserFetchAll)
	route.GET("users/:id", Controllers.UserFetchSingle)
	route.POST("users", Controllers.UserCreate)
	route.PUT("users/:id", Controllers.UserUpdate)
	route.DELETE("users/:id", Controllers.UserDelete)

	return route
}