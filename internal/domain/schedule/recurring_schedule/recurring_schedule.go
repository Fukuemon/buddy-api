package recurring_schedule

import (
	"api-buddy/domain/common"
	errorDomain "api-buddy/domain/error"
	facilityDomain "api-buddy/domain/facility"
	recurringRuleDomain "api-buddy/domain/schedule/recurring_rule"
	scheduleTypeDomain "api-buddy/domain/schedule/schedule_type"
	userDomain "api-buddy/domain/user"
	visitInfoDomain "api-buddy/domain/visit_info"

	"github.com/Fukuemon/go-pkg/query"
	"github.com/Fukuemon/go-pkg/ulid"
)

type RecurringSchedule struct {
	ID                      string                             `gorm:"primaryKey"`
	RecurringRule           *recurringRuleDomain.RecurringRule `gorm:"foreignKey:RecurringRuleID"`
	RecurringRuleID         string
	ScheduleType            *scheduleTypeDomain.ScheduleType `gorm:"foreignKey:ScheduleTypeID"`
	ScheduleTypeID          string
	Date                    common.Date
	StartTime               common.Time
	EndTime                 common.Time
	IsOverTimeWork          bool
	Staff                   *userDomain.User
	StaffID                 string `gorm:"foreignKey::StaffID"`
	Facility                *facilityDomain.Facility
	FacilityID              string
	VisitInfoID             *string
	VisitInfo               *visitInfoDomain.VisitInfo `gorm:"foreignKey:VisitInfoID"`
	Title                   *string
	Description             *string
	RecurringExclusionDates *common.JSONSlice[int]
	common.CommonModel
}

type RecurringScheduleOption func(*RecurringSchedule) error

func WithTitle(title *string) RecurringScheduleOption {
	return func(s *RecurringSchedule) error {
		s.Title = title
		return nil
	}
}

func WithDescription(description *string) RecurringScheduleOption {
	return func(s *RecurringSchedule) error {
		s.Description = description
		return nil
	}
}

func WithVisitInfo(visitInfo *visitInfoDomain.VisitInfo) RecurringScheduleOption {
	return func(s *RecurringSchedule) error {
		if visitInfo == nil {
			return errorDomain.NewError("訪問情報が含まれていません")
		}
		s.VisitInfo = visitInfo
		s.VisitInfoID = &visitInfo.ID
		return nil
	}
}

func WithRecurringExclusionDates(recurring_exclusion_dates common.JSONSlice[int]) RecurringScheduleOption {
	return func(s *RecurringSchedule) error {
		if len(recurring_exclusion_dates) == 0 {
			return errorDomain.NewError("除外日が含まれていません")
		}
		s.RecurringExclusionDates = &recurring_exclusion_dates
		return nil
	}
}

func NewRecurringSchedule(
	recurringRule *recurringRuleDomain.RecurringRule,
	scheduleType *scheduleTypeDomain.ScheduleType,
	date common.Date,
	startTime common.Time,
	endTime common.Time,
	staff *userDomain.User,
	facility *facilityDomain.Facility,
	options ...RecurringScheduleOption,
) (*RecurringSchedule, error) {
	return newRecurringSchedule(
		ulid.NewULID(),
		recurringRule,
		scheduleType,
		date,
		startTime,
		endTime,
		staff,
		facility,
		options...,
	)
}

func newRecurringSchedule(
	id string,
	recurringRule *recurringRuleDomain.RecurringRule,
	scheduleType *scheduleTypeDomain.ScheduleType,
	date common.Date,
	startTime common.Time,
	endTime common.Time,
	staff *userDomain.User,
	facility *facilityDomain.Facility,
	options ...RecurringScheduleOption,
) (*RecurringSchedule, error) {
	if endTime.Before(startTime.Time) {
		return nil, errorDomain.NewError("終了時間が開始時間より前です")
	}
	if endTime.Equal(startTime.Time) {
		return nil, errorDomain.NewError("終了時間が開始時間と同じです")
	}

	recurringSchedule := &RecurringSchedule{
		ID:                      id,
		RecurringRule:           recurringRule,
		RecurringRuleID:         recurringRule.ID,
		ScheduleType:            scheduleType,
		ScheduleTypeID:          scheduleType.ID,
		Date:                    date,
		StartTime:               startTime,
		EndTime:                 endTime,
		IsOverTimeWork:          startTime.Hour() >= 17,
		Staff:                   staff,
		StaffID:                 staff.ID,
		Facility:                facility,
		FacilityID:              facility.ID,
		VisitInfo:               nil,
		VisitInfoID:             nil,
		Title:                   nil,
		Description:             nil,
		RecurringExclusionDates: nil,
	}

	for _, option := range options {
		if err := option(recurringSchedule); err != nil {
			return nil, err
		}
	}

	// 通常予定の場合、タイトルは必須
	// 予定種別が通常の場合、タイトルは必須
	if recurringSchedule.ScheduleType.Name == scheduleTypeDomain.Normal {
		if recurringSchedule.Title == nil {
			return nil, errorDomain.NewError("通常の予定の場合、タイトルは必須です")
		}
	}

	return recurringSchedule, nil
}

var RecurringScheduleRelationMappings = map[string]query.RelationMapping{
	"recurring_rule": {
		TableName:   "recurring_rules",
		JoinKey:     "recurring_rules.id = recurring_schedules.recurring_rule_id",
		FilterField: "recurring_rule.name",
	},
	"schedule_type": {
		TableName:   "schedule_types",
		JoinKey:     "schedule_types.id = recurring_schedules.schedule_type_id",
		FilterField: "schedule_type.name",
	},
	"staff": {
		TableName:   "users",
		JoinKey:     "users.id = recurring_schedules.staff_id",
		FilterField: "users.name",
	},
	"facility": {
		TableName:   "facilities",
		JoinKey:     "facilities.id = recurring_schedules.facility_id",
		FilterField: "facilities.name",
	},
	"visit_info": {
		TableName:   "visit_infos",
		JoinKey:     "visit_infos.id = recurring_schedules.visit_info_id",
		FilterField: "visit_info.name",
	},
}
