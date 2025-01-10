package schedule

import (
	"api-buddy/domain/common"
	scheduleDomain "api-buddy/domain/schedule"
	recurringRuleDomain "api-buddy/domain/schedule/recurring_rule"
	recurringScheduleDomain "api-buddy/domain/schedule/recurring_schedule"
	scheduleTypeDomain "api-buddy/domain/schedule/schedule_type"
	visitInfoDomain "api-buddy/domain/visit_info"
	visitCategoryDomain "api-buddy/domain/visit_info/visit_category"
	"context"
	"log"

	"github.com/Fukuemon/go-pkg/query"
)

type FetchScheduleUseCase struct {
	scheduleRepository          scheduleDomain.ScheduleRepository
	recurringScheduleRepository recurringScheduleDomain.RecurringScheduleRepository
}

func NewFetchScheduleUseCase(scheduleRepository scheduleDomain.ScheduleRepository, recurringScheduleRepository recurringScheduleDomain.RecurringScheduleRepository) *FetchScheduleUseCase {
	return &FetchScheduleUseCase{
		scheduleRepository:          scheduleRepository,
		recurringScheduleRepository: recurringScheduleRepository,
	}
}

// Output DTO
type FetchScheduleUseCaseOutputDto struct {
	Schedules          []*ScheduleModel
	RecurringSchedules []*RecurringScheduleModel
}

type ScheduleModel struct {
	ID             string
	ScheduleType   scheduleTypeDomain.ScheduleTypeEnum
	Date           common.Date
	StartTime      common.Time
	EndTime        common.Time
	IsOverTimeWork bool
	StaffName      string
	VisitInfo      *visitInfoModel
	Title          string
	Description    string
	CancelReason   string
}

type visitInfoModel struct {
	Patient         string
	AssignedStaff   string
	Companion       string
	Route           *routeModel
	ServiceCode     string
	VisitCategories []visitCategoryModel
}

type routeModel struct {
	TravelTime  int
	Address     string
	Destination string
}

type visitCategoryModel struct {
	ID   string
	Name visitCategoryDomain.VisitCategoryType
}

type RecurringScheduleModel struct {
	ID             string
	RecurringRule  *RecurringRuleModel
	ScheduleType   scheduleTypeDomain.ScheduleTypeEnum
	Date           common.Date
	StartTime      common.Time
	EndTime        common.Time
	IsOverTimeWork bool
	StaffName      string
	VisitInfo      *visitInfoModel
	Title          string
	Description    string
	ExclusionDates *common.JSONSlice[int]
}

type RecurringRuleModel struct {
	Frequency   recurringRuleDomain.FrequencyEnum
	DayOfWeek   int
	DayOfMonth  int
	WeekOfMonth int
	StartDate   common.Date
	EndDate     common.Date
}

// Input DTO
type FetchScheduleUseCaseInputDto struct {
	ScheduleType string
	SortField    string
	SortOrder    string
}

func (uc *FetchScheduleUseCase) Run(ctx context.Context, facilityID string, input FetchScheduleUseCaseInputDto) (*FetchScheduleUseCaseOutputDto, error) {
	outputDto := FetchScheduleUseCaseOutputDto{}
	filters := uc.buildFilters(input)
	sortOption := query.SortOption{
		Field: input.SortField,
		Order: input.SortOrder,
	}

	schedules, err := uc.scheduleRepository.FindByFacilityID(ctx, facilityID, filters, sortOption)
	if err != nil {
		return nil, err
	}

	scheduleList := make([]scheduleDomain.Schedule, len(schedules))
	for i, schedule := range schedules {
		scheduleList[i] = *schedule
	}
	outputDto.Schedules = uc.buildScheduleModels(scheduleList)

	recurringSchedules, err := uc.recurringScheduleRepository.FindByFacilityID(ctx, facilityID, filters, sortOption)
	if err != nil {
		return nil, err
	}
	log.Println("recurringSchedules", recurringSchedules)
	recurringScheduleList := make([]recurringScheduleDomain.RecurringSchedule, len(recurringSchedules))
	for i, recurringSchedule := range recurringSchedules {
		recurringScheduleList[i] = *recurringSchedule
	}
	outputDto.RecurringSchedules = uc.buildRecurringScheduleModels(recurringScheduleList)

	return &outputDto, nil
}

func (uc *FetchScheduleUseCase) buildFilters(input FetchScheduleUseCaseInputDto) []query.Filter {
	var filters []query.Filter
	if input.ScheduleType != "" {
		filters = append(filters, &query.ByFieldFilter{
			Field:           "schedule_type",
			Value:           input.ScheduleType,
			RelationMapping: scheduleDomain.ScheduleRelationMappings,
		})
	}
	return filters
}

func (uc *FetchScheduleUseCase) buildScheduleModels(schedules []scheduleDomain.Schedule) []*ScheduleModel {
	models := make([]*ScheduleModel, 0, len(schedules))
	for _, schedule := range schedules {
		models = append(models, &ScheduleModel{
			ID:             schedule.ID,
			ScheduleType:   schedule.ScheduleType.Name,
			Date:           schedule.Date,
			StartTime:      schedule.StartTime,
			EndTime:        schedule.EndTime,
			IsOverTimeWork: schedule.IsOverTimeWork,
			StaffName:      schedule.Staff.Username,
			VisitInfo:      uc.buildVisitInfoModel(schedule.VisitInfo),
			Title:          schedule.Title,
			Description:    schedule.Description,
			CancelReason: func() string {
				if schedule.ScheduleCancel != nil {
					return schedule.ScheduleCancel.Reason
				}
				return ""
			}(),
		})
	}
	return models
}

func (uc *FetchScheduleUseCase) buildRecurringScheduleModels(recurringSchedules []recurringScheduleDomain.RecurringSchedule) []*RecurringScheduleModel {
	models := make([]*RecurringScheduleModel, 0, len(recurringSchedules))
	for _, recurringSchedule := range recurringSchedules {
		models = append(models, &RecurringScheduleModel{
			ID:             recurringSchedule.ID,
			ScheduleType:   recurringSchedule.ScheduleType.Name,
			Date:           recurringSchedule.Date,
			StartTime:      recurringSchedule.StartTime,
			EndTime:        recurringSchedule.EndTime,
			IsOverTimeWork: recurringSchedule.IsOverTimeWork,
			StaffName:      recurringSchedule.Staff.Username,
			VisitInfo:      uc.buildVisitInfoModel(recurringSchedule.VisitInfo),
			Title:          recurringSchedule.Title,
			Description:    recurringSchedule.Description,
			ExclusionDates: recurringSchedule.RecurringExclusionDates,
			RecurringRule: func() *RecurringRuleModel {
				if recurringSchedule.RecurringRule != nil {
					return &RecurringRuleModel{
						Frequency:   recurringSchedule.RecurringRule.Frequency,
						DayOfWeek:   recurringSchedule.RecurringRule.DayOfWeek,
						DayOfMonth:  recurringSchedule.RecurringRule.DayOfMonth,
						WeekOfMonth: recurringSchedule.RecurringRule.WeekOfMonth,
						StartDate:   recurringSchedule.RecurringRule.StartDate,
						EndDate:     recurringSchedule.RecurringRule.EndDate,
					}
				}
				return nil
			}(),
		})
	}
	return models
}

func (uc *FetchScheduleUseCase) buildVisitInfoModel(visitInfo *visitInfoDomain.VisitInfo) *visitInfoModel {
	if visitInfo == nil {
		return nil // visitInfoがnilの場合、nilを返す
	}

	var route *routeModel
	if visitInfo.Route != nil {
		route = &routeModel{
			TravelTime: visitInfo.Route.TravelTime,
			Address: func() string {
				if visitInfo.Route.Address != nil {
					return visitInfo.Route.Address.JoinAddress()
				}
				return ""
			}(),
			Destination: func() string {
				if visitInfo.Route.Destination != nil {
					return visitInfo.Route.Destination.JoinAddress()
				}
				return ""
			}(),
		}
	}

	return &visitInfoModel{
		Patient:       visitInfo.Patient.Name,
		AssignedStaff: visitInfo.AssignedStaff.Username,
		Companion: func() string {
			if visitInfo.Companion != nil {
				return visitInfo.Companion.Username
			}
			return ""
		}(),
		Route: route,
		ServiceCode: func() string {
			if visitInfo.ServiceCode != nil {
				return visitInfo.ServiceCode.Code
			}
			return ""
		}(),
		VisitCategories: func() []visitCategoryModel {
			if visitInfo.VisitCategories == nil {
				return []visitCategoryModel{}
			}
			models := make([]visitCategoryModel, 0, len(visitInfo.VisitCategories))
			for _, visitCategory := range visitInfo.VisitCategories {
				models = append(models, visitCategoryModel{
					ID:   visitCategory.ID,
					Name: visitCategory.Name,
				})
			}
			return models
		}(),
	}
}
