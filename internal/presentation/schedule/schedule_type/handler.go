package schedule_type

import (
	errorDomain "api-buddy/domain/error"
	_ "api-buddy/presentation/common"
	"api-buddy/presentation/settings"
	scheduleTypeUse "api-buddy/usecase/schedule/schedule_type"

	pathValidator "github.com/Fukuemon/go-pkg/validator/gin"

	"github.com/gin-gonic/gin"
)

type handler struct {
	fetchScheduleTypesUseCase *scheduleTypeUse.FetchScheduleTypesUseCase
}

func NewHandler(fetchScheduleTypesUseCase *scheduleTypeUse.FetchScheduleTypesUseCase) *handler {
	return &handler{
		fetchScheduleTypesUseCase: fetchScheduleTypesUseCase,
	}
}

// FetchScheduleTypes godoc
// @Summary 予定種別一覧を取得する
// @Tags ScheduleType
// @Accept json
// @Produce json
// @Success 200 {object} ScheduleTypeListResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /facilities/{facility_id}/schedules/schedule_types [get]
func (h *handler) FetchScheduleTypes(ctx *gin.Context) {
	facilityId := pathValidator.Param(ctx, "facility_id", "required", "ulid")

	err := facilityId.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}
	output, err := h.fetchScheduleTypesUseCase.Run(ctx, facilityId.ParamValue)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := make(ScheduleTypeListResponse, 0, len(output))
	for _, o := range output {
		response = append(response, ScheduleTypeResponse{
			ID:   o.ID,
			Name: o.Name,
		})
	}
	settings.ReturnStatusOK(ctx, response)
}
