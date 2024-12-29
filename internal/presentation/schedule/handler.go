package schedule

import (
	errorDomain "api-buddy/domain/error"
	visitInfoDomain "api-buddy/domain/visit_info"
	routeDomain "api-buddy/domain/visit_info/route"
	_ "api-buddy/presentation/common"
	"api-buddy/presentation/settings"
	"api-buddy/usecase/schedule"
	recurringSchedule "api-buddy/usecase/schedule/recurring_schedule"

	"github.com/Fukuemon/go-pkg/validator"
	"github.com/gin-gonic/gin"
)

type handler struct {
	createScheduleUseCase          *schedule.CreateScheduleUseCase
	createRecurringScheduleUseCase *recurringSchedule.CreateRecurringScheduleUseCase
}

func NewHandler(
	createScheduleUseCase *schedule.CreateScheduleUseCase,
	createRecurringScheduleUseCase *recurringSchedule.CreateRecurringScheduleUseCase,
) *handler {
	return &handler{
		createScheduleUseCase:          createScheduleUseCase,
		createRecurringScheduleUseCase: createRecurringScheduleUseCase,
	}
}

// CreateSchedule godoc
// @Summary 予定を作成する
// @Tags Schedule
// @Accept json
// @Produce json
// @Param request body CreateScheduleRequest true "予定作成リクエスト"
// @Param facility_id path string true "施設ID"
// @Success 201 {object} CreateScheduleResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /facilities/{facility_id}/schedules [post]
func (h *handler) CreateSchedule(ctx *gin.Context) {
	facilityID := ctx.Param("facility_id")
	var params CreateScheduleRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	if err := validator.StructValidation(params); err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	input := schedule.CreateUseCaseInputDto{
		ScheduleTypeID: params.ScheduleTypeID,
		Date:           params.Date,
		StartTime:      params.StartTime,
		EndTime:        params.EndTime,
		StaffID:        params.StaffID,
		FacilityID:     facilityID,
		Title:          params.Title,
		Description:    params.Description,
		VisitInfo: &visitInfoDomain.VisitInfoModel{
			PatientID:     params.VisitInfo.PatientID,
			AssignStaffID: params.VisitInfo.AssignStaffID,
			CompanionID:   params.VisitInfo.CompanionID,
			Route: &routeDomain.RouteModel{
				TravelTime:    params.VisitInfo.Route.TravelTime,
				FromAddressID: params.VisitInfo.Route.FromAddressID,
				DestinationID: params.VisitInfo.Route.DestinationID,
			},
			ServiceCodeID:    params.VisitInfo.ServiceCodeID,
			VisitCategoryIDs: params.VisitInfo.VisitCategoryIDs,
		},
	}

	output, err := h.createScheduleUseCase.Run(ctx, input)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := CreateScheduleResponse{
		ID:             output.ID,
		ScheduleTypeID: output.ScheduleTypeID,
		Date:           output.Date,
		StartTime:      output.StartTime,
		EndTime:        output.EndTime,
		StaffID:        output.StaffID,
		VisitInfo: &VisitInfoResponseModel{
			ID:              output.VisitInfo.ID,
			PatientID:       output.VisitInfo.PatientID,
			AssignedStaffID: output.VisitInfo.AssignedStaffID,
			CompanionID:     output.VisitInfo.CompanionID,
			Route: &RouteResponseModel{
				TravelTime:    output.VisitInfo.Route.TravelTime,
				AddressID:     output.VisitInfo.Route.AddressID,
				DestinationID: output.VisitInfo.Route.DestinationID,
			},
			ServiceCodeID: output.VisitInfo.ServiceCodeID,
			VisitCategories: func() []VisitCategoryResponseModel {
				var categories []VisitCategoryResponseModel
				for _, category := range output.VisitInfo.VisitCategories {
					categories = append(categories, VisitCategoryResponseModel{
						ID:   category.ID,
						Name: category.Name,
					})
				}
				return categories
			}(),
		},
		Title:       output.Title,
		Description: output.Description,
	}

	settings.ReturnStatusCreated(ctx, response)

}

// CreateRecurringSchedule godoc
// @Summary 繰り返し予定を作成する
// @Tags Schedule
// @Accept json
// @Produce json
// @Param request body CreateRecurringScheduleRequest true "予定作成リクエスト"
// @Param facility_id path string true "施設ID"
// @Success 201 {object} CreateRecurringScheduleResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /facilities/{facility_id}/schedules/recurring [post]
func (h *handler) CreateRecurringSchedule(ctx *gin.Context) {
	facilityID := ctx.Param("facility_id")
	var params CreateRecurringScheduleRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	if err := validator.StructValidation(params); err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	input := recurringSchedule.CreateRecurringScheduleUseCaseInputDto{
		ScheduleTypeID: params.ScheduleTypeID,
		RecurringRule: recurringSchedule.RecurringRuleModel{
			Frequency:   params.RecurringRuleModel.Frequency,
			DaysOfWeek:  params.RecurringRuleModel.DaysOfWeek,
			DayOfMonth:  params.RecurringRuleModel.DayOfMonth,
			WeekOfMonth: params.RecurringRuleModel.WeekOfMonth,
			StartDate:   params.RecurringRuleModel.StartDate,
			EndDate:     params.RecurringRuleModel.EndDate,
		},
		Date:        params.Date,
		StartTime:   params.StartTime,
		EndTime:     params.EndTime,
		StaffID:     params.StaffID,
		FacilityID:  facilityID,
		Title:       params.Title,
		Description: params.Description,
		VisitInfo: &visitInfoDomain.VisitInfoModel{
			PatientID:     params.VisitInfo.PatientID,
			AssignStaffID: params.VisitInfo.AssignStaffID,
			CompanionID:   params.VisitInfo.CompanionID,
			Route: &routeDomain.RouteModel{
				TravelTime:    params.VisitInfo.Route.TravelTime,
				FromAddressID: params.VisitInfo.Route.FromAddressID,
				DestinationID: params.VisitInfo.Route.DestinationID,
			},
			ServiceCodeID:    params.VisitInfo.ServiceCodeID,
			VisitCategoryIDs: params.VisitInfo.VisitCategoryIDs,
		},
	}

	output, err := h.createRecurringScheduleUseCase.Run(ctx, input)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := CreateRecurringScheduleResponse{
		ID:             output.ID,
		ScheduleTypeID: output.ScheduleTypeID,
		Date:           output.Date,
		StartTime:      output.StartTime,
		EndTime:        output.EndTime,
		StaffID:        output.StaffID,
		VisitInfo: &VisitInfoResponseModel{
			ID:              output.VisitInfo.ID,
			PatientID:       output.VisitInfo.PatientID,
			AssignedStaffID: output.VisitInfo.AssignedStaffID,
			CompanionID:     output.VisitInfo.CompanionID,
			Route: &RouteResponseModel{
				TravelTime:    output.VisitInfo.Route.TravelTime,
				AddressID:     output.VisitInfo.Route.AddressID,
				DestinationID: output.VisitInfo.Route.DestinationID,
			},
			ServiceCodeID: output.VisitInfo.ServiceCodeID,
			VisitCategories: func() []VisitCategoryResponseModel {
				var categories []VisitCategoryResponseModel
				for _, category := range output.VisitInfo.VisitCategories {
					categories = append(categories, VisitCategoryResponseModel{
						ID:   category.ID,
						Name: category.Name,
					})
				}
				return categories
			}(),
		},
		Title:       output.Title,
		Description: output.Description,
		RecurringRule: &RecurringRuleResponseModel{
			Frequency:   output.RecurringRule.Frequency,
			DaysOfWeek:  output.RecurringRule.DaysOfWeek,
			DayOfMonth:  output.RecurringRule.DayOfMonth,
			WeekOfMonth: output.RecurringRule.WeekOfMonth,
			StartDate:   output.RecurringRule.StartDate,
			EndDate:     output.RecurringRule.EndDate,
		},
	}
	settings.ReturnStatusCreated(ctx, response)
}

// GetSchedule godoc
// @Summary 予定を取得する
// @Tags Schedule
// @Accept json

// CreateChangeRecurringSchedule godoc
// @Summary 繰り返し予定を変更する
// @Tags Schedule
// @Accept json
// @Produce json
// @Param request body CreateChangeRecurringScheduleRequest true "予定作成リクエスト"
// @Param facility_id path string true "施設ID"
// @Param recurring_schedule_id path string true "繰り返し予定ID"
// @Success 201 {object} CreateChangeRecurringScheduleResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /facilities/{facility_id}/schedules/{recurring_schedule_id} [put]
