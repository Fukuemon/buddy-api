package schedule

import (
	errorDomain "api-buddy/domain/error"
	facilityDomain "api-buddy/domain/facility"
	recurringScheduleDomain "api-buddy/domain/schedule/recurring_schedule"
	scheduleCancelDomain "api-buddy/domain/schedule/schedule_cancel"
	scheduleTypeDomain "api-buddy/domain/schedule/schedule_type"
	userDomain "api-buddy/domain/user"
	visitInfoDomain "api-buddy/domain/visit_info"
	"time"

	"github.com/Fukuemon/go-pkg/query"
	"github.com/Fukuemon/go-pkg/ulid"
)

type Option struct {
	VisitInfo             *visitInfoDomain.VisitInfo
	VisitInfoID           *string
	Title                 *string
	RecurringScheduleID   *string
	RecurringSchedule     *recurringScheduleDomain.RecurringSchedule
	BeforeChangeDate      *time.Time
	BeforeChangeStartTime *time.Time
	Description           *string
	ScheduleCancel        *scheduleCancelDomain.ScheduleCancel
	ScheduleCancelID      *string
}

type Schedule struct {
	ID                    string
	ScheduleType          *scheduleTypeDomain.ScheduleType
	ScheduleTypeID        string `gorm:"foreignKey:ID;references:ScheduleTypeID"`
	Date                  time.Time
	StartTime             time.Time
	EndTime               time.Time
	IsOverTimeWork        bool
	Staff                 *userDomain.User
	StaffID               string `gorm:"foreignKey:ID;references:StaffID"`
	Facility              *facilityDomain.Facility
	FacilityID            string `gorm:"foreignKey:ID;references:FacilityID"`
	Title                 string
	VisitInfo             *visitInfoDomain.VisitInfo
	VisitInfoID           string `gorm:"foreignKey:ID;references:VisitInfoID"`
	RecurringSchedule     *recurringScheduleDomain.RecurringSchedule
	RecurringScheduleID   string `gorm:"foreignKey:ID;references:RecurringScheduleID"`
	BeforeChangeDate      *time.Time
	BeforeChangeStartTime *time.Time
	Description           string
	ScheduleCancel        *scheduleCancelDomain.ScheduleCancel
	ScheduleCancelID      string `gorm:"foreignKey:ID;references:ScheduleCancelID"`
}

func NewSchedule(
	schedule_type *scheduleTypeDomain.ScheduleType,
	date time.Time,
	start_time time.Time,
	end_time time.Time,
	is_over_time_work bool,
	staff *userDomain.User,
	facility *facilityDomain.Facility,
	options *Option,
) (*Schedule, error) {
	return newSchedule(
		ulid.NewULID(),
		schedule_type,
		date,
		start_time,
		end_time,
		is_over_time_work,
		staff,
		facility,
		options,
	)
}

func newSchedule(
	id string,
	schedule_type *scheduleTypeDomain.ScheduleType,
	date time.Time,
	start_time time.Time,
	end_time time.Time,
	is_over_time_work bool,
	staff *userDomain.User,
	facility *facilityDomain.Facility,
	options *Option,
) (*Schedule, error) {
	schedule := &Schedule{
		ID:             id,
		ScheduleType:   schedule_type,
		Date:           date,
		StartTime:      start_time,
		EndTime:        end_time,
		IsOverTimeWork: is_over_time_work,
		Staff:          staff,
		StaffID:        staff.ID,
		Facility:       facility,
		FacilityID:     facility.ID,
	}

	if options != nil {
		// 通常予定の場合
		if schedule_type.Name == scheduleTypeDomain.Normal {
			if options.Title == nil {
				return nil, errorDomain.NewError("タイトルが含まれていません")
			}
			schedule.Title = *options.Title
		}

		// 訪問予定の場合
		if schedule_type.Name == scheduleTypeDomain.Visit {

			if options.VisitInfoID == nil {
				return nil, errorDomain.NewError("訪問情報が含まれていません")
			}
			schedule.VisitInfoID = *options.VisitInfoID
			schedule.VisitInfo = options.VisitInfo
		}

		// 繰り返し予定からの変更だった場合
		if options.RecurringScheduleID != nil {
			schedule.RecurringScheduleID = *options.RecurringScheduleID
			if options.BeforeChangeDate == nil {
				return nil, errorDomain.NewError("変更前の日付が含まれていません")
			}
			if options.BeforeChangeStartTime == nil {
				return nil, errorDomain.NewError("変更前の開始時間が含まれていません")
			}
			schedule.BeforeChangeDate = options.BeforeChangeDate
			schedule.BeforeChangeStartTime = options.BeforeChangeStartTime
			schedule.RecurringSchedule = options.RecurringSchedule
		}

		// 補足情報
		if options.Description != nil {
			schedule.Description = *options.Description
		}

		// 予定キャンセル情報
		if options.ScheduleCancelID != nil {
			schedule.ScheduleCancelID = *options.ScheduleCancelID
			schedule.ScheduleCancel = options.ScheduleCancel
		}
	}
	// 開始時間が17:00以降の場合
	if schedule.StartTime.Hour() >= 17 {
		schedule.IsOverTimeWork = true
	}

	// 終了時間が開始時間より前の場合
	if schedule.EndTime.Before(schedule.StartTime) {
		return nil, errorDomain.NewError("終了時間が開始時間より前です")
	}

	// 終了時間が開始時間と同じ場合
	if schedule.EndTime.Equal(schedule.StartTime) {
		return nil, errorDomain.NewError("終了時間が開始時間と同じです")
	}

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
