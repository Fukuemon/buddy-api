package schedule

import (
	"api-buddy/domain/common"
	errorDomain "api-buddy/domain/error"
	facilityDomain "api-buddy/domain/facility"
	recurringScheduleDomain "api-buddy/domain/schedule/recurring_schedule"
	scheduleCancelDomain "api-buddy/domain/schedule/schedule_cancel"
	scheduleTypeDomain "api-buddy/domain/schedule/schedule_type"
	userDomain "api-buddy/domain/user"
	visitInfoDomain "api-buddy/domain/visit_info"

	"github.com/Fukuemon/go-pkg/query"
	"github.com/Fukuemon/go-pkg/ulid"
)

type Schedule struct {
	ID                  string                           `gorm:"primaryKey"`
	ScheduleType        *scheduleTypeDomain.ScheduleType `gorm:"foreignKey:ScheduleTypeID"`
	ScheduleTypeID      string
	Date                common.Date
	StartTime           common.Time
	EndTime             common.Time
	IsOverTimeWork      bool
	Staff               *userDomain.User `gorm:"foreignKey:StaffID"`
	StaffID             string
	Facility            *facilityDomain.Facility `gorm:"foreignKey:FacilityID"`
	FacilityID          string
	Title               string
	VisitInfo           *visitInfoDomain.VisitInfo `gorm:"foreignKey:VisitInfoID"`
	VisitInfoID         *string
	RecurringSchedule   *recurringScheduleDomain.RecurringSchedule `gorm:"foreignKey:RecurringScheduleID"`
	RecurringScheduleID *string
	Description         string
	ScheduleCancel      *scheduleCancelDomain.ScheduleCancel `gorm:"foreignKey:ScheduleCancelID"`
	ScheduleCancelID    *string
	common.CommonModel
}

type ScheduleOption func(*Schedule) error

func WithTitle(title string) ScheduleOption {
	return func(s *Schedule) error {
		if title == "" {
			err := errorDomain.NewError("タイトルが含まれていません")
			return errorDomain.WrapError(errorDomain.InvalidInputErr, err)
		}
		s.Title = title
		return nil
	}
}

func WithVisitInfo(visitInfo *visitInfoDomain.VisitInfo) ScheduleOption {
	return func(s *Schedule) error {
		if visitInfo == nil {
			err := errorDomain.NewError("訪問情報が含まれていません")
			return errorDomain.WrapError(errorDomain.InvalidInputErr, err)
		}
		s.VisitInfo = visitInfo
		s.VisitInfoID = &visitInfo.ID
		return nil
	}
}

func WithRecurringSchedule(
	recurringSchedule *recurringScheduleDomain.RecurringSchedule,
) ScheduleOption {
	return func(s *Schedule) error {
		if recurringSchedule == nil {
			err := errorDomain.NewError("繰り返し予定の情報が含まれていません")
			return errorDomain.WrapError(errorDomain.InvalidInputErr, err)
		}
		s.RecurringSchedule = recurringSchedule
		s.RecurringScheduleID = &recurringSchedule.ID
		return nil
	}
}

func WithDescription(description string) ScheduleOption {
	return func(s *Schedule) error {
		s.Description = description
		return nil
	}
}

func WithScheduleCancel(scheduleCancel *scheduleCancelDomain.ScheduleCancel) ScheduleOption {
	return func(s *Schedule) error {
		if scheduleCancel == nil {
			err := errorDomain.NewError("キャンセル情報が含まれていません")
			return errorDomain.WrapError(errorDomain.InvalidInputErr, err)
		}
		s.ScheduleCancel = scheduleCancel
		s.ScheduleCancelID = &scheduleCancel.ID
		return nil
	}
}

// NewSchedule creates a new schedule instance
func NewSchedule(
	scheduleType *scheduleTypeDomain.ScheduleType,
	date common.Date,
	startTime common.Time,
	endTime common.Time,
	staff *userDomain.User,
	facility *facilityDomain.Facility,
	options ...ScheduleOption,
) (*Schedule, error) {
	// Validate start and end times
	if endTime.Before(startTime.Time) {
		err := errorDomain.NewError("終了時間が開始時間より前です")
		return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
	}
	if endTime.Equal(startTime.Time) {
		err := errorDomain.NewError("終了時間が開始時間と同じです")
		return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
	}

	schedule := &Schedule{
		ID:                  ulid.NewULID(),
		ScheduleType:        scheduleType,
		ScheduleTypeID:      scheduleType.ID,
		Date:                date,
		StartTime:           startTime,
		EndTime:             endTime,
		IsOverTimeWork:      startTime.Hour() >= 17,
		Staff:               staff,
		StaffID:             staff.ID,
		Facility:            facility,
		FacilityID:          facility.ID,
		RecurringSchedule:   nil,
		RecurringScheduleID: nil,
		VisitInfo:           nil,
		VisitInfoID:         nil,
		Description:         "",
		ScheduleCancel:      nil,
		ScheduleCancelID:    nil,
	}

	// Apply options
	for _, option := range options {
		if err := option(schedule); err != nil {
			return nil, err
		}
	}

	// 予定種別が通常の場合、タイトルは必須
	if schedule.ScheduleType.Name == scheduleTypeDomain.Normal {
		if schedule.Title == "" {
			err := errorDomain.NewError("通常の予定の場合、タイトルは必須です")
			return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
		}
	}

	// 予定種別が訪問の場合、訪問情報は必須
	if schedule.ScheduleType.Name == scheduleTypeDomain.Visit {
		if schedule.VisitInfo == nil {
			err := errorDomain.NewError("訪問の予定の場合、訪問情報は必須です")
			return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
		}
	}

	common.InitializeCommonModel(&schedule.CommonModel)
	return schedule, nil
}

var ScheduleRelationMappings = map[string]query.RelationMapping{
	"staff": {
		TableName:   "users",
		JoinKey:     "users.id = schedules.staff_id",
		FilterField: "users.name",
	},
	"facility": {
		TableName:   "facilities",
		JoinKey:     "facilities.id = schedules.facility_id",
		FilterField: "facilities.name",
	},
	"schedule_type": {
		TableName:   "schedule_types",
		JoinKey:     "schedule_types.id = schedules.schedule_type_id",
		FilterField: "schedule_types.name",
	},
	"visit_info": {
		TableName:   "visit_infos",
		JoinKey:     "visit_infos.id = schedules.visit_info_id",
		FilterField: "visit_infos.name",
	},
	"recurring_schedule": {
		TableName:   "recurring_schedules",
		JoinKey:     "recurring_schedules.id = schedules.recurring_schedule_id",
		FilterField: "recurring_schedules.id",
	},
	"schedule_cancel": {
		TableName:   "schedule_cancels",
		JoinKey:     "schedule_cancels.id = schedules.schedule_cancel_id",
		FilterField: "schedule_cancels.id",
	},
}
