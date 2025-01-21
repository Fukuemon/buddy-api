package schedule

import (
	"api-buddy/domain/common"
	recurringScheduleDomain "api-buddy/domain/schedule/recurring_rule"
)

type CreateScheduleRequest struct {
	ScheduleTypeID      string                 `json:"schedule_type_id" validate:"required,ulid"`
	Date                common.Date            `json:"date" validate:"required" swaggertype:"string" format:"date"`
	StartTime           common.Time            `json:"start_time" validate:"required"`
	EndTime             common.Time            `json:"end_time" validate:"required"`
	StaffID             string                 `json:"staff_id" validate:"required,ulid"`
	VisitInfo           *VisitInfoRequestModel `json:"visit_info" validate:"omitempty"`
	Title               *string                `json:"title" validate:"omitempty"`
	Description         *string                `json:"description" validate:"omitempty"`
	RecurringScheduleID *string                `json:"recurring_schedule_id" validate:"omitempty,ulid"`
}

type CreateRecurringScheduleRequest struct {
	ScheduleTypeID     string                     `json:"schedule_type_id" validate:"required,ulid"`
	Date               common.Date                `json:"date" validate:"required" swaggertype:"string" format:"date"`
	StartTime          common.Time                `json:"start_time" validate:"required"`
	EndTime            common.Time                `json:"end_time" validate:"required"`
	StaffID            string                     `json:"staff_id" validate:"required,ulid"`
	VisitInfo          *VisitInfoRequestModel     `json:"visit_info" validate:"omitempty"`
	Title              *string                    `json:"title" validate:"omitempty"`
	Description        *string                    `json:"description" validate:"omitempty"`
	RecurringRuleModel *RecurringRuleRequestModel `json:"recurring_rule" validate:"omitempty"`
}

type CreateChangeRecurringScheduleRequest struct {
	ScheduleTypeID          string                 `json:"schedule_type_id" validate:"required,ulid"`
	Date                    common.Date            `json:"date" validate:"required" swaggertype:"string" format:"date"`
	StartTime               common.Time            `json:"start_time" validate:"required"`
	EndTime                 common.Time            `json:"end_time" validate:"required"`
	StaffID                 string                 `json:"staff_id" validate:"required,ulid"`
	VisitInfo               *VisitInfoRequestModel `json:"visit_info" validate:"omitempty"`
	Title                   *string                `json:"title" validate:"omitempty"`
	Description             *string                `json:"description" validate:"omitempty"`
	RecurringScheduleID     *string                `json:"recurring_schedule_id" validate:"recurring,ulid"`
	BeforeChangeDate        common.Date            `json:"before_change_date" validate:"required" swaggertype:"string" format:"date"`
	BeforeChangeStartTime   common.Time            `json:"before_change_start_time" validate:"required"`
	RecurringExclusionDates []int                  `json:"recurring_exclusion_dates" validate:"omitempty"`
}

type VisitInfoRequestModel struct {
	PatientID        string             `json:"patient_id" validate:"required,ulid"`
	AssignStaffID    string             `json:"assign_staff_id" validate:"required,ulid"`
	CompanionID      *string            `json:"companion_id" validate:"omitempty,ulid"`
	Route            *RouteRequestModel `json:"route" validate:"omitempty"`
	ServiceCodeID    string             `json:"service_code_id" validate:"required,ulid"`
	VisitCategoryIDs []*string          `json:"visit_category_ids" validate:"omitempty,dive,ulid"`
}

type RouteRequestModel struct {
	TravelTime    int    `json:"travel_time" validate:"omitempty"`
	FromAddressID string `json:"from_address_id" validate:"required,ulid"`
	DestinationID string `json:"destination_id" validate:"required,ulid"`
}

type RecurringRuleRequestModel struct {
	Frequency   recurringScheduleDomain.FrequencyEnum `json:"frequency" validate:"required"`
	DayOfWeek   *int                                  `json:"day_of_week" validate:"omitempty"`
	DayOfMonth  *int                                  `json:"day_of_month" validate:"omitempty"`
	WeekOfMonth *int                                  `json:"week_of_month" validate:"omitempty"`
	EndDate     *common.Date                          `json:"end_date" validate:"omitempty" swaggertype:"string" format:"date"`
}
