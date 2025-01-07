package service_code

import (
	_ "api-buddy/presentation/common"
	"api-buddy/presentation/settings"
	serviceCodeUse "api-buddy/usecase/visit_info/service_code"

	"github.com/gin-gonic/gin"
)

type handler struct {
	fetchServiceCodesUseCase *serviceCodeUse.FetchServiceCodesUseCase
}

func NewHandler(fetchServiceCodesUseCase *serviceCodeUse.FetchServiceCodesUseCase) *handler {
	return &handler{
		fetchServiceCodesUseCase: fetchServiceCodesUseCase,
	}
}

// FetchServiceCodes godoc
// @Summary サービスコード一覧を取得する
// @Tags ServiceCode
// @Accept json
// @Produce json
// @Success 200 {object} ServiceCodeListResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /visit_infos/service_codes [get]
func (h *handler) FetchServiceCodes(ctx *gin.Context) {
	output, err := h.fetchServiceCodesUseCase.Run(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := make(ServiceCodeListResponse, 0, len(output))
	for _, o := range output {
		response = append(response, ServiceCodeResponse{
			ID:                    o.ID,
			Code:                  o.Code,
			ServiceTimeRangeStart: o.ServiceTimeRangeStart,
			ServiceTimeRangeEnd:   o.ServiceTimeRangeEnd,
		})
	}

	settings.ReturnStatusOK(ctx, response)
}
