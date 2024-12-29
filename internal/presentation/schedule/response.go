package schedule

import (
	"api-buddy/domain/common"
	recurringRuleDomain "api-buddy/domain/schedule/recurring_rule"
	visitCategoryDomain "api-buddy/domain/visit_info/visit_category"
)

type CreateScheduleResponse struct {
	ID             string                  `json:"id"`
	ScheduleTypeID string                  `json:"schedule_type_id"`
	Date           common.Date             `json:"date"`
	StartTime      common.Time             `json:"start_time"`
	EndTime        common.Time             `json:"end_time"`
	StaffID        string                  `json:"staff_id"`
	VisitInfo      *VisitInfoResponseModel `json:"visit_info"`
	Title          string                  `json:"title"`
	Description    string                  `json:"description"`
}

type VisitInfoResponseModel struct {
	ID              string                       `json:"id"`
	PatientID       string                       `json:"patient_id"`
	AssignedStaffID string                       `json:"assigned_staff_id"`
	CompanionID     string                       `json:"companion_id"`
	Route           *RouteResponseModel          `json:"route"`
	ServiceCodeID   string                       `json:"service_code_id"`
	VisitCategories []VisitCategoryResponseModel `json:"visit_categories"`
}

type VisitCategoryResponseModel struct {
	ID   string                                `json:"id"`
	Name visitCategoryDomain.VisitCategoryType `json:"name"`
}

type RouteResponseModel struct {
	TravelTime    int    `json:"travel_time"`
	AddressID     string `json:"address_id"`
	DestinationID string `json:"destination_id"`
}

type CreateRecurringScheduleResponse struct {
	ID             string                      `json:"id"`
	ScheduleTypeID string                      `json:"schedule_type_id"`
	Date           common.Date                 `json:"date"`
	StartTime      common.Time                 `json:"start_time"`
	EndTime        common.Time                 `json:"end_time"`
	StaffID        string                      `json:"staff_id"`
	VisitInfo      *VisitInfoResponseModel     `json:"visit_info"`
	Title          string                      `json:"title"`
	Description    string                      `json:"description"`
	RecurringRule  *RecurringRuleResponseModel `json:"recurring_rule"`
}

type RecurringRuleResponseModel struct {
	Frequency   recurringRuleDomain.FrequencyEnum `json:"frequency"`
	DaysOfWeek  int                               `json:"days_of_week"`
	DayOfMonth  int                               `json:"day_of_month"`
	WeekOfMonth int                               `json:"week_of_month"`
	StartDate   common.Date                       `json:"start_date"`
	EndDate     common.Date                       `json:"end_date"`
}
