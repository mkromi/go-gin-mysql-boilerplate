package Services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessResponse(c *gin.Context, status int, message interface{}, data interface {}) {
	c.AbortWithStatusJSON(status, gin.H{"message": message, "data": data})
}

func OkResponse(c *gin.Context, status int, message interface{}) {
	c.AbortWithStatusJSON(status, gin.H{"message": message})
}

func ErrorResponse(c *gin.Context, status int, message interface{}, error interface{}) {
	c.AbortWithStatusJSON(status, gin.H{"message": message, "errors": error})
}

func Success(c *gin.Context, message interface{}, data interface{}) {
	SuccessResponse(c, http.StatusOK, message, data)
}

func Created(c *gin.Context, message interface{}, data interface{}) {
	SuccessResponse(c, http.StatusCreated, message, data)
}

func Updated(c *gin.Context, message interface{}, data interface{}) {
	SuccessResponse(c, http.StatusAccepted, message, data)
}

func Deleted(c *gin.Context, message interface{}, data interface{}) {
	SuccessResponse(c, http.StatusAccepted, message, data)
}

func NoContent(c *gin.Context, message interface{}, data interface{}) {
	SuccessResponse(c, http.StatusNoContent, message, data)
}

func BadRequest(c *gin.Context, message interface{}, error interface{}) {
	ErrorResponse(c, http.StatusBadRequest, message, error)
}

func Unauthorized(c *gin.Context, message interface{}, error interface{}) {
	ErrorResponse(c, http.StatusUnauthorized, message, error)
}

func NotAcceptable(c *gin.Context, message interface{}, error interface{}) {
	ErrorResponse(c, http.StatusNotAcceptable, message, error)
}

func ValidationError(c *gin.Context, message interface{}, error interface{}) {
	ErrorResponse(c, http.StatusUnprocessableEntity, message, error)
}
