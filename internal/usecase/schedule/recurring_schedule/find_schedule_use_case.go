package recurring_schedule

import (
	"api-buddy/domain/common"
	recurringRuleDomain "api-buddy/domain/schedule/recurring_rule"
	recurringScheduleDomain "api-buddy/domain/schedule/recurring_schedule"
	scheduleTypeDomain "api-buddy/domain/schedule/schedule_type"
	scheduleUse "api-buddy/usecase/schedule"
	"context"
)

type FindRecurringScheduleUseCase struct {
	recurringScheduleRepository recurringScheduleDomain.RecurringScheduleRepository
}

func NewFindRecurringScheduleUseCase(recurringScheduleRepository recurringScheduleDomain.RecurringScheduleRepository) *FindRecurringScheduleUseCase {
	return &FindRecurringScheduleUseCase{
		recurringScheduleRepository: recurringScheduleRepository,
	}
}

type FindRecurringScheduleUseCaseOutputDto struct {
	ID             string
	ScheduleType   *scheduleTypeDomain.ScheduleTypeEnum
	RecurringRule  *recurringRuleDomain.RecurringRule
	Date           common.Date
	StartTime      common.Time
	EndTime        common.Time
	IsOverTimeWork bool
	StaffID        string
	StaffName      string
	VisitInfo      *scheduleUse.VisitInfoModel
	Title          *string
	Description    *string
	ExclusionDates *common.JSONSlice[int]
}

func (uc *FindRecurringScheduleUseCase) Run(ctx context.Context, scheduleID string) (*FindRecurringScheduleUseCaseOutputDto, error) {
	recurringSchedule, err := uc.recurringScheduleRepository.FindByID(ctx, scheduleID)
	if err != nil {
		return nil, err
	}

	outputDto := &FindRecurringScheduleUseCaseOutputDto{
		ID:             recurringSchedule.ID,
		ScheduleType:   &recurringSchedule.ScheduleType.Name,
		RecurringRule:  recurringSchedule.RecurringRule,
		Date:           recurringSchedule.Date,
		StartTime:      recurringSchedule.StartTime,
		EndTime:        recurringSchedule.EndTime,
		IsOverTimeWork: recurringSchedule.IsOverTimeWork,
		StaffID:        recurringSchedule.Staff.ID,
		StaffName:      recurringSchedule.Staff.Username,
		VisitInfo:      scheduleUse.BuildVisitInfoModel(recurringSchedule.VisitInfo),
		Title:          recurringSchedule.Title,
		Description:    recurringSchedule.Description,
		ExclusionDates: recurringSchedule.RecurringExclusionDates,
	}

	return outputDto, nil
}
