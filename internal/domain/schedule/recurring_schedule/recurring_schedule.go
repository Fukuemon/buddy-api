package recurring_schedule

import (
	errorDomain "api-buddy/domain/error"
	facilityDomain "api-buddy/domain/facility"
	recurringRuleDomain "api-buddy/domain/schedule/recurring_rule"
	scheduleTypeDomain "api-buddy/domain/schedule/schedule_type"
	userDomain "api-buddy/domain/user"
	visitInfoDomain "api-buddy/domain/visit_info"
	"time"
)

type Option struct {
	VisitInfo   *visitInfoDomain.VisitInfo
	VisitInfoID *string
	Title       *string
	Description *string
}

type RecurringSchedule struct {
	ID              string
	RecurringRule   *recurringRuleDomain.RecurringRule
	RecurringRuleID string
	ScheduleType    *scheduleTypeDomain.ScheduleType
	ScheduleTypeID  string `gorm:"foreignKey:ID;references:ScheduleTypeID"`
	Date            time.Time
	StartTime       time.Time
	EndTime         time.Time
	IsOverTimeWork  bool
	Staff           *userDomain.User
	StaffID         string `gorm:"foreignKey:ID;references:StaffID"`
	VisitInfo       *visitInfoDomain.VisitInfo
	Facility        *facilityDomain.Facility
	FacilityID      string
	VisitInfoID     string
	Title           string
	Description     string
}

func newRecurringSchedule(
	id string,
	recurring_rule *recurringRuleDomain.RecurringRule,
	schedule_type *scheduleTypeDomain.ScheduleType,
	date time.Time,
	start_time time.Time,
	end_time time.Time,
	is_over_time_work bool,
	staff *userDomain.User,
	facility *facilityDomain.Facility,
	options *Option,
) (*RecurringSchedule, error) {
	recurringSchedule := &RecurringSchedule{
		ID:              id,
		RecurringRule:   recurring_rule,
		RecurringRuleID: recurring_rule.ID,
		ScheduleType:    schedule_type,
		ScheduleTypeID:  schedule_type.ID,
		Date:            date,
		StartTime:       start_time,
		EndTime:         end_time,
		IsOverTimeWork:  is_over_time_work,
		Staff:           staff,
		StaffID:         staff.ID,
		Facility:        facility,
		FacilityID:      facility.ID,
	}

	if options != nil {
		if options != nil {
			// 通常予定の場合
			if schedule_type.Name == scheduleTypeDomain.ScheduleTypeNormal {
				if options.Title == nil {
					return nil, errorDomain.NewError("タイトルが含まれていません")
				}
				recurringSchedule.Title = *options.Title
			}

			// 訪問予定の場合
			if schedule_type.Name == scheduleTypeDomain.ScheduleTypeVisit {

				if options.VisitInfoID == nil {
					return nil, errorDomain.NewError("訪問情報が含まれていません")
				}
				recurringSchedule.VisitInfoID = *options.VisitInfoID
				recurringSchedule.VisitInfo = options.VisitInfo
			}

			// 補足情報
			if options.Description != nil {
				recurringSchedule.Description = *options.Description
			}
		}
		// 開始時間が17:00以降の場合
		if recurringSchedule.StartTime.Hour() >= 17 {
			recurringSchedule.IsOverTimeWork = true
		}
		// 終了時間が開始時間より前の場合
		if recurringSchedule.EndTime.Before(recurringSchedule.StartTime) {
			return nil, errorDomain.NewError("終了時間が開始時間より前です")
		}

		// 終了時間が開始時間と同じ場合
		if recurringSchedule.EndTime.Equal(recurringSchedule.StartTime) {
			return nil, errorDomain.NewError("終了時間が開始時間と同じです")
		}
	}

	return recurringSchedule, nil
}
