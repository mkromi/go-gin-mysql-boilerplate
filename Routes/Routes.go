package Routes
import (
	"github.com/gin-gonic/gin"
	"go-gin-mysql-boilerplate/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.POST("user", Controllers.UserCreate)
	}
	return r
}