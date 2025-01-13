package schedule

import (
	"api-buddy/domain/common"
	scheduleDomain "api-buddy/domain/schedule"
	scheduleTypeDomain "api-buddy/domain/schedule/schedule_type"
	visitInfoDomain "api-buddy/domain/visit_info"
	"context"
)

type FindScheduleUseCase struct {
	scheduleRepository scheduleDomain.ScheduleRepository
}

func NewFindScheduleUseCase(scheduleRepository scheduleDomain.ScheduleRepository) *FindScheduleUseCase {
	return &FindScheduleUseCase{
		scheduleRepository: scheduleRepository,
	}
}

type FindScheduleUseCaseOutputDto struct {
	ID             string
	ScheduleType   *scheduleTypeDomain.ScheduleTypeEnum
	Date           common.Date
	StartTime      common.Time
	EndTime        common.Time
	IsOverTimeWork bool
	StaffName      string
	VisitInfo      *VisitInfoModel
	Title          *string
	Description    *string
	CancelReason   *string
}

func (uc *FindScheduleUseCase) Run(ctx context.Context, scheduleID string) (*FindScheduleUseCaseOutputDto, error) {
	schedule, err := uc.scheduleRepository.FindByID(ctx, scheduleID)
	if err != nil {
		return nil, err
	}

	outputDto := &FindScheduleUseCaseOutputDto{
		ID:             schedule.ID,
		ScheduleType:   &schedule.ScheduleType.Name,
		Date:           schedule.Date,
		StartTime:      schedule.StartTime,
		EndTime:        schedule.EndTime,
		IsOverTimeWork: schedule.IsOverTimeWork,
		StaffName:      schedule.Staff.Username,
		VisitInfo:      BuildVisitInfoModel(schedule.VisitInfo),
		Title:          &schedule.Title,
		Description:    &schedule.Description,
		CancelReason: func() *string {
			if schedule.ScheduleCancel != nil {
				return &schedule.ScheduleCancel.Reason
			}
			return nil
		}(),
	}

	return outputDto, nil
}

func BuildVisitInfoModel(visitInfo *visitInfoDomain.VisitInfo) *VisitInfoModel {
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

	return &VisitInfoModel{
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
