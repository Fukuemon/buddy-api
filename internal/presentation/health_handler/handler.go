package health_handler

import (
	"api-buddy/presentation/settings"

	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *gin.Context) {
	response := HealthResponse{
		Status: "ok",
	}

	settings.ReturnStatusOK(ctx, response)
}
