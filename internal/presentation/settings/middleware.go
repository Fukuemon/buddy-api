package settings

import (
	errDomain "api-buddy/domain/error"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			if errors.As(err.Err, new(*errDomain.Error)) {
				HandleErrorResponse(c, err.Err)
				log.Println(err.Err)
			} else {
				ReturnStatusInternalServerError(c, err.Err)
			}
		}
	}
}
