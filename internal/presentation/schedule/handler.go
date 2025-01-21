package schedule

import (
	"api-buddy/domain/common"
	errorDomain "api-buddy/domain/error"
	visitInfoDomain "api-buddy/domain/visit_info"
	routeDomain "api-buddy/domain/visit_info/route"
	_ "api-buddy/presentation/common"
	"api-buddy/presentation/settings"
	"api-buddy/usecase/schedule"
	recurringSchedule "api-buddy/usecase/schedule/recurring_schedule"

	"github.com/Fukuemon/go-pkg/validator"
	pathValidator "github.com/Fukuemon/go-pkg/validator/gin"
	"github.com/gin-gonic/gin"
)

type handler struct {
	createScheduleUseCase          *schedule.CreateScheduleUseCase
	createRecurringScheduleUseCase *recurringSchedule.CreateRecurringScheduleUseCase
	fetchScheduleUseCase           *schedule.FetchScheduleUseCase
	findScheduleUseCase            *schedule.FindScheduleUseCase
	findREcurringScheduleUseCase   *recurringSchedule.FindRecurringScheduleUseCase
}

func NewHandler(
	createScheduleUseCase *schedule.CreateScheduleUseCase,
	createRecurringScheduleUseCase *recurringSchedule.CreateRecurringScheduleUseCase,
	fetchScheduleUseCase *schedule.FetchScheduleUseCase,
	findScheduleUseCase *schedule.FindScheduleUseCase,
	findRecurringScheduleUseCase *recurringSchedule.FindRecurringScheduleUseCase,
) *handler {
	return &handler{
		createScheduleUseCase:          createScheduleUseCase,
		createRecurringScheduleUseCase: createRecurringScheduleUseCase,
		fetchScheduleUseCase:           fetchScheduleUseCase,
		findScheduleUseCase:            findScheduleUseCase,
		findREcurringScheduleUseCase:   findRecurringScheduleUseCase,
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

	var visitInfo *visitInfoDomain.VisitInfoModel
	if params.VisitInfo != nil {
		visitInfo = &visitInfoDomain.VisitInfoModel{
			PatientID:     params.VisitInfo.PatientID,
			AssignStaffID: params.VisitInfo.AssignStaffID,
			CompanionID:   params.VisitInfo.CompanionID,
			Route: func() *routeDomain.RouteModel {
				if params.VisitInfo.Route != nil {
					return &routeDomain.RouteModel{
						TravelTime:    params.VisitInfo.Route.TravelTime,
						FromAddressID: params.VisitInfo.Route.FromAddressID,
						DestinationID: params.VisitInfo.Route.DestinationID,
					}
				}
				return nil
			}(),
			ServiceCodeID:    params.VisitInfo.ServiceCodeID,
			VisitCategoryIDs: params.VisitInfo.VisitCategoryIDs,
		}
	}

	input := schedule.CreateUseCaseInputDto{
		ScheduleTypeID:      params.ScheduleTypeID,
		Date:                params.Date,
		StartTime:           params.StartTime,
		EndTime:             params.EndTime,
		StaffID:             params.StaffID,
		FacilityID:          facilityID,
		Title:               params.Title,
		Description:         params.Description,
		VisitInfo:           visitInfo,
		RecurringScheduleID: params.RecurringScheduleID,
	}

	output, err := h.createScheduleUseCase.Run(ctx, input)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := CreateScheduleResponse{
		ID:           output.ID,
		ScheduleType: string(output.ScheduleType.Name),
		Date:         output.Date,
		StartTime:    output.StartTime,
		EndTime:      output.EndTime,
		StaffName:    output.Staff.Username,
		VisitInfo: func() *VisitInfoResponseModel {
			if output.VisitInfo == nil {
				return nil
			}
			return &VisitInfoResponseModel{
				ID:                output.VisitInfo.ID,
				PatientName:       output.VisitInfo.Patient.Name,
				AssignedStaffName: output.VisitInfo.AssignedStaff.Username,
				CompanionName: func() string {
					if output.VisitInfo.Companion != nil {
						return output.VisitInfo.Companion.Username
					}
					return ""
				}(),
				Route: func() *RouteResponseModel {
					if output.VisitInfo.Route != nil {
						return &RouteResponseModel{
							TravelTime:    output.VisitInfo.Route.TravelTime,
							AddressID:     output.VisitInfo.Route.AddressID,
							DestinationID: output.VisitInfo.Route.DestinationID,
						}
					}
					return nil
				}(),
				ServiceCode: output.VisitInfo.ServiceCode.Code,
				VisitCategories: func() []VisitCategoryResponseModel {
					var categories []VisitCategoryResponseModel
					if output.VisitInfo.VisitCategories != nil {
						for _, category := range output.VisitInfo.VisitCategories {
							categories = append(categories, VisitCategoryResponseModel{
								ID:   category.ID,
								Name: category.Name,
							})
						}
					}
					return categories
				}(),
			}
		}(),
		Title: func() string {
			if output.Title != nil {
				return *output.Title
			}
			return "" // デフォルト値
		}(),
		Description: func() string {
			if output.Description != nil {
				return *output.Description
			}
			return "" // デフォルト値
		}(),
		RecurringScheduleID: func() string {
			if output.RecurringScheduleID == nil {
				return ""
			} else {
				return *output.RecurringScheduleID
			}
		}(),
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

	var visitInfo *visitInfoDomain.VisitInfoModel
	if params.VisitInfo != nil {
		visitInfo = &visitInfoDomain.VisitInfoModel{
			PatientID:     params.VisitInfo.PatientID,
			AssignStaffID: params.VisitInfo.AssignStaffID,
			CompanionID:   params.VisitInfo.CompanionID,
			Route: func() *routeDomain.RouteModel {
				if params.VisitInfo.Route != nil {
					return &routeDomain.RouteModel{
						TravelTime:    params.VisitInfo.Route.TravelTime,
						FromAddressID: params.VisitInfo.Route.FromAddressID,
						DestinationID: params.VisitInfo.Route.DestinationID,
					}
				}
				return nil
			}(),
			ServiceCodeID:    params.VisitInfo.ServiceCodeID,
			VisitCategoryIDs: params.VisitInfo.VisitCategoryIDs,
		}
	}

	input := recurringSchedule.CreateRecurringScheduleUseCaseInputDto{
		ScheduleTypeID: params.ScheduleTypeID,
		RecurringRule: recurringSchedule.RecurringRuleModel{
			Frequency:   params.RecurringRuleModel.Frequency,
			DayOfWeek:   params.RecurringRuleModel.DayOfWeek,
			DayOfMonth:  params.RecurringRuleModel.DayOfMonth,
			WeekOfMonth: params.RecurringRuleModel.WeekOfMonth,
			EndDate:     params.RecurringRuleModel.EndDate,
		},
		Date:        params.Date,
		StartTime:   params.StartTime,
		EndTime:     params.EndTime,
		StaffID:     params.StaffID,
		FacilityID:  facilityID,
		Title:       params.Title,
		Description: params.Description,
		VisitInfo:   visitInfo,
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
		VisitInfo: func() *VisitInfoResponseModel {
			if output.VisitInfo == nil {
				return nil
			}
			return &VisitInfoResponseModel{
				ID:                output.VisitInfo.ID,
				PatientName:       output.VisitInfo.Patient.Name,
				AssignedStaffName: output.VisitInfo.AssignedStaff.Username,
				CompanionName: func() string {
					if output.VisitInfo.Companion != nil {
						return output.VisitInfo.Companion.Username
					}
					return ""
				}(),
				Route: func() *RouteResponseModel {
					if output.VisitInfo.Route != nil {
						return &RouteResponseModel{
							TravelTime:    output.VisitInfo.Route.TravelTime,
							AddressID:     output.VisitInfo.Route.AddressID,
							DestinationID: output.VisitInfo.Route.DestinationID,
						}
					}
					return nil
				}(),
				ServiceCode: output.VisitInfo.ServiceCode.Code,
				VisitCategories: func() []VisitCategoryResponseModel {
					var categories []VisitCategoryResponseModel
					if output.VisitInfo.VisitCategories != nil {
						for _, category := range output.VisitInfo.VisitCategories {
							categories = append(categories, VisitCategoryResponseModel{
								ID:   category.ID,
								Name: category.Name,
							})
						}
					}
					return categories
				}(),
			}
		}(),
		Title: func() string {
			if output.Title != nil {
				return *output.Title
			}
			return "" // デフォルト値
		}(),
		Description: func() string {
			if output.Description != nil {
				return *output.Description
			}
			return "" // デフォルト値
		}(),
		RecurringRule: &RecurringRuleResponseModel{
			Frequency: output.RecurringRule.Frequency,
			DaysOfWeek: func() int {
				if output.RecurringRule.DayOfWeek != nil {
					return *output.RecurringRule.DayOfWeek
				}
				return 0
			}(),
			DayOfMonth: func() int {
				if output.RecurringRule.DayOfMonth != nil {
					return *output.RecurringRule.DayOfMonth
				}
				return 0
			}(),
			WeekOfMonth: func() int {
				if output.RecurringRule.WeekOfMonth != nil {
					return *output.RecurringRule.WeekOfMonth
				}
				return 0
			}(),
			EndDate: func() *common.Date {
				if output.RecurringRule.EndDate != nil {
					return output.RecurringRule.EndDate
				}
				return nil
			}(),
		},
	}

	settings.ReturnStatusCreated(ctx, response)
}

// GetSchedule swagger
// @Summary 施設に紐づく予定と繰り返し予定を全て取得する
// @Tags Schedule
// @Accept json
// @Produce json
// @Param facility_id path string true "施設ID"
// @Success 200 {object} ScheduleListResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /facilities/{facility_id}/schedules [get]
func (h *handler) FetchByFacilityId(ctx *gin.Context) {
	facilityID := pathValidator.Param(ctx, "facility_id", "required", "ulid")
	err := facilityID.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	scheduleType := ctx.Query("schedule_type")
	sortField := ctx.Query("sort_field")
	sortOrder := ctx.Query("sort_order")

	input := schedule.FetchScheduleUseCaseInputDto{
		ScheduleType: scheduleType,
		SortField:    sortField,
		SortOrder:    sortOrder,
	}

	output, err := h.fetchScheduleUseCase.Run(ctx, facilityID.ParamValue, input)
	if err != nil {
		ctx.Error(err)
		return
	}

	schedules := ScheduleListResponse{
		Schedules: func() []*ScheduleResponse {
			var schedules []*ScheduleResponse
			for _, schedule := range output.Schedules {
				schedules = append(schedules, &ScheduleResponse{
					ID:           schedule.ID,
					ScheduleType: string(schedule.ScheduleType),
					Date:         schedule.Date,
					StartTime:    schedule.StartTime,
					EndTime:      schedule.EndTime,
					StaffID:      schedule.StaffID,
					StaffName:    schedule.StaffName,
					VisitInfo: func() *VisitInfoResponseModel {
						if schedule.VisitInfo == nil {
							return nil
						}
						return &VisitInfoResponseModel{
							PatientName:       schedule.VisitInfo.Patient,
							AssignedStaffName: schedule.VisitInfo.AssignedStaff,
							CompanionName: func() string {
								if schedule.VisitInfo.Companion != nil {
									return *schedule.VisitInfo.Companion
								}
								return ""
							}(),
							Route: func() *RouteResponseModel {
								if schedule.VisitInfo.Route != nil {
									return &RouteResponseModel{
										TravelTime:    schedule.VisitInfo.Route.TravelTime,
										AddressID:     schedule.VisitInfo.Route.Address,
										DestinationID: schedule.VisitInfo.Route.Destination,
									}
								}
								return nil
							}(),
							ServiceCode: schedule.VisitInfo.ServiceCode,
							VisitCategories: func() []VisitCategoryResponseModel {
								var categories []VisitCategoryResponseModel
								if schedule.VisitInfo.VisitCategories != nil {
									for _, category := range schedule.VisitInfo.VisitCategories {
										categories = append(categories, VisitCategoryResponseModel{
											ID:   category.ID,
											Name: category.Name,
										})
									}
								}
								return categories
							}(),
						}
					}(),
					Title: func() string {
						if schedule.Title != nil {
							return *schedule.Title
						}
						return ""
					}(),
					Description: func() string {
						if schedule.Description != nil {
							return *schedule.Description
						}
						return ""
					}(),
				})
			}
			return schedules
		}(),
		RecurringSchedules: func() []*RecurringScheduleResponse {
			var recurringSchedules []*RecurringScheduleResponse
			for _, recurringSchedule := range output.RecurringSchedules {
				recurringSchedules = append(recurringSchedules, &RecurringScheduleResponse{
					ID: recurringSchedule.ID,
					RecurringRule: func() *RecurringRuleResponseModel {
						if recurringSchedule.RecurringRule != nil {
							return &RecurringRuleResponseModel{
								Frequency: recurringSchedule.RecurringRule.Frequency,
								DaysOfWeek: func() int {
									if recurringSchedule.RecurringRule.DayOfWeek != nil {
										return *recurringSchedule.RecurringRule.DayOfWeek
									}
									return 0
								}(),
								DayOfMonth: func() int {
									if recurringSchedule.RecurringRule.DayOfMonth != nil {
										return *recurringSchedule.RecurringRule.DayOfMonth
									}
									return 0
								}(),
								WeekOfMonth: func() int {
									if recurringSchedule.RecurringRule.WeekOfMonth != nil {
										return *recurringSchedule.RecurringRule.WeekOfMonth
									}
									return 0
								}(),
								EndDate: func() *common.Date {
									if recurringSchedule.RecurringRule.EndDate != nil {
										return recurringSchedule.RecurringRule.EndDate
									}
									return nil
								}(),
							}
						}
						return nil
					}(),
					ScheduleType: string(recurringSchedule.ScheduleType),
					Date:         recurringSchedule.Date,
					StartTime:    recurringSchedule.StartTime,
					EndTime:      recurringSchedule.EndTime,
					StaffID:      recurringSchedule.StaffID,
					StaffName:    recurringSchedule.StaffName,
					VisitInfo: func() *VisitInfoResponseModel {
						if recurringSchedule.VisitInfo == nil {
							return nil
						}
						return &VisitInfoResponseModel{
							PatientName:       recurringSchedule.VisitInfo.Patient,
							AssignedStaffName: recurringSchedule.VisitInfo.AssignedStaff,
							CompanionName: func() string {
								if recurringSchedule.VisitInfo.Companion != nil {
									return *recurringSchedule.VisitInfo.Companion
								}
								return ""
							}(),
							Route: func() *RouteResponseModel {
								if recurringSchedule.VisitInfo.Route != nil {
									return &RouteResponseModel{
										TravelTime:    recurringSchedule.VisitInfo.Route.TravelTime,
										AddressID:     recurringSchedule.VisitInfo.Route.Address,
										DestinationID: recurringSchedule.VisitInfo.Route.Destination,
									}
								}
								return nil
							}(),
							ServiceCode: recurringSchedule.VisitInfo.ServiceCode,
							VisitCategories: func() []VisitCategoryResponseModel {
								var categories []VisitCategoryResponseModel
								if recurringSchedule.VisitInfo.VisitCategories != nil {
									for _, category := range recurringSchedule.VisitInfo.VisitCategories {
										categories = append(categories, VisitCategoryResponseModel{
											ID:   category.ID,
											Name: category.Name,
										})
									}
								}
								return categories
							}(),
						}
					}(),
					Title: func() string {
						if recurringSchedule.Title != nil {
							return *recurringSchedule.Title
						}
						return ""
					}(),
					Description: func() string {
						if recurringSchedule.Description != nil {
							return *recurringSchedule.Description
						}
						return ""
					}(),
					ExclusionDates: func() []int {
						if recurringSchedule.ExclusionDates == nil {
							return nil
						}
						return *recurringSchedule.ExclusionDates
					}(),
				})
			}
			return recurringSchedules
		}(),
	}

	settings.ReturnStatusOK(ctx, schedules)
}

// FindSchedule godoc
// @Summary 単一の予定を取得する
// @Tags Schedule
// @Accept json
// @Produce json
// @Param facility_id path string true "施設ID"
// @Param schedule_id path string true "予定ID"
// @Success 200 {object} ScheduleResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /schedules/{schedule_id} [get]
func (h *handler) FindScheduleByID(ctx *gin.Context) {
	scheduleID := pathValidator.Param(ctx, "schedule_id", "required", "ulid")
	err := scheduleID.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	output, err := h.findScheduleUseCase.Run(ctx, scheduleID.ParamValue)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := ScheduleResponse{
		ID:             output.ID,
		ScheduleType:   string(*output.ScheduleType),
		Date:           output.Date,
		StartTime:      output.StartTime,
		EndTime:        output.EndTime,
		IsOverTimeWork: output.IsOverTimeWork,
		StaffID:        output.StaffID,
		StaffName:      output.StaffName,
		VisitInfo: func() *VisitInfoResponseModel {
			if output.VisitInfo == nil {
				return nil
			}
			return &VisitInfoResponseModel{
				PatientName:       output.VisitInfo.Patient,
				AssignedStaffName: output.VisitInfo.AssignedStaff,
				CompanionName: func() string {
					if output.VisitInfo.Companion != nil {
						return *output.VisitInfo.Companion
					}
					return ""
				}(),
				Route: func() *RouteResponseModel {
					if output.VisitInfo.Route != nil {
						return &RouteResponseModel{
							TravelTime:    output.VisitInfo.Route.TravelTime,
							AddressID:     output.VisitInfo.Route.Address,
							DestinationID: output.VisitInfo.Route.Destination,
						}
					}
					return nil
				}(),
				ServiceCode: output.VisitInfo.ServiceCode,
				VisitCategories: func() []VisitCategoryResponseModel {
					var categories []VisitCategoryResponseModel
					if output.VisitInfo.VisitCategories != nil {
						for _, category := range output.VisitInfo.VisitCategories {
							categories = append(categories, VisitCategoryResponseModel{
								ID:   category.ID,
								Name: category.Name,
							})
						}
					}
					return categories
				}(),
			}
		}(),
		Title:       *output.Title,
		Description: *output.Description,
		CancelReason: func() string {
			if output.CancelReason != nil {
				return *output.CancelReason
			}
			return ""
		}(),
	}

	settings.ReturnStatusOK(ctx, response)
}

// FindRecurringSchedule godoc
// @Summary 単一の繰り返し予定を取得する
// @Tags Schedule
// @Accept json
// @Produce json
// @Param facility_id path string true "施設ID"
// @Param recurring_schedule_id path string true "繰り返し予定ID"
// @Success 200 {object} RecurringScheduleResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /schedules/recurring/{recurring_schedule_id} [get]
func (h *handler) FindRecurringScheduleByID(ctx *gin.Context) {
	recurringScheduleID := pathValidator.Param(ctx, "recurring_schedule_id", "required", "ulid")
	err := recurringScheduleID.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	output, err := h.findREcurringScheduleUseCase.Run(ctx, recurringScheduleID.ParamValue)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := RecurringScheduleResponse{
		ID: output.ID,
		RecurringRule: &RecurringRuleResponseModel{
			Frequency: output.RecurringRule.Frequency,
			DaysOfWeek: func() int {
				if output.RecurringRule.DayOfWeek != nil {
					return *output.RecurringRule.DayOfWeek
				}
				return 0
			}(),
			DayOfMonth: func() int {
				if output.RecurringRule.DayOfMonth != nil {
					return *output.RecurringRule.DayOfMonth
				}
				return 0
			}(),
			WeekOfMonth: func() int {
				if output.RecurringRule.WeekOfMonth != nil {
					return *output.RecurringRule.WeekOfMonth
				}
				return 0
			}(),
			EndDate: func() *common.Date {
				if output.RecurringRule.EndDate != nil {
					return output.RecurringRule.EndDate
				}
				return nil
			}(),
		},
		ScheduleType:   string(*output.ScheduleType),
		Date:           output.Date,
		StartTime:      output.StartTime,
		EndTime:        output.EndTime,
		IsOverTimeWork: output.IsOverTimeWork,
		StaffID:        output.StaffID,
		StaffName:      output.StaffName,
		VisitInfo: func() *VisitInfoResponseModel {
			if output.VisitInfo == nil {
				return nil
			}
			return &VisitInfoResponseModel{
				PatientName:       output.VisitInfo.Patient,
				AssignedStaffName: output.VisitInfo.AssignedStaff,
				CompanionName: func() string {
					if output.VisitInfo.Companion != nil {
						return *output.VisitInfo.Companion
					}
					return ""
				}(),
				Route: func() *RouteResponseModel {
					if output.VisitInfo.Route != nil {
						return &RouteResponseModel{
							TravelTime:    output.VisitInfo.Route.TravelTime,
							AddressID:     output.VisitInfo.Route.Address,
							DestinationID: output.VisitInfo.Route.Destination,
						}
					}
					return nil
				}(),
				ServiceCode: output.VisitInfo.ServiceCode,
				VisitCategories: func() []VisitCategoryResponseModel {
					var categories []VisitCategoryResponseModel
					if output.VisitInfo.VisitCategories != nil {
						for _, category := range output.VisitInfo.VisitCategories {
							categories = append(categories, VisitCategoryResponseModel{
								ID:   category.ID,
								Name: category.Name,
							})
						}
					}
					return categories
				}(),
			}
		}(),
		Title: func() string {
			if output.Title != nil {
				return *output.Title
			}
			return ""
		}(),
		Description: func() string {
			if output.Description != nil {
				return *output.Description
			}
			return ""
		}(),
		ExclusionDates: func() []int {
			if output.ExclusionDates == nil {
				return nil
			}
			return *output.ExclusionDates
		}(),
	}

	settings.ReturnStatusOK(ctx, response)
}

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
