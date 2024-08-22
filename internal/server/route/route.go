package route

import (
	"api-buddy/presentation/health_handler"
	"api-buddy/presentation/settings"

	"github.com/gin-gonic/gin"
)

func InitRoute(api *gin.Engine) {
	api.Use(settings.ErrorHandler())
	v1 := api.Group("/v1")
	v1.GET("/health", health_handler.HealthCheck)
}
