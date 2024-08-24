package health_handler

import (
	"api-buddy/presentation/settings"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary      ヘルスチェック用エンドポイント
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  HealthResponse
// @Router       /health [get]
func HealthCheck(ctx *gin.Context) {
	response := HealthResponse{
		Status: "ok",
	}

	settings.ReturnStatusOK(ctx, response)
}
