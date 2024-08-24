package route

import (
	"api-buddy/presentation/health_handler"
	"api-buddy/presentation/settings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoute(api *gin.Engine) {
	api.Use(settings.ErrorHandler())

	v1 := api.Group("/v1")
	// ヘルスチェック
	v1.GET("/health", health_handler.HealthCheck)

	// Swagger
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
