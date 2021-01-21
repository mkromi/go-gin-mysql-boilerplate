package Routes
import (
	"github.com/gin-gonic/gin"
	"go-gin-mysql-boilerplate/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("users", Controllers.UserFetchAll)
		api.GET("users/:id", Controllers.UserFetchSingle)
		api.POST("users", Controllers.UserCreate)
		api.PUT("users/:id", Controllers.UserUpdate)
		api.DELETE("users/:id", Controllers.UserDelete)
	}
	return r
}