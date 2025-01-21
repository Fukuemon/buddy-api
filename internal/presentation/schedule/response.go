package schedule

import (
	"api-buddy/domain/common"
	recurringRuleDomain "api-buddy/domain/schedule/recurring_rule"
	visitCategoryDomain "api-buddy/domain/visit_info/visit_category"
)

type CreateScheduleResponse struct {
	ID                  string                  `json:"id"`
	ScheduleType        string                  `json:"schedule_type"`
	Date                common.Date             `json:"date"`
	StartTime           common.Time             `json:"start_time"`
	EndTime             common.Time             `json:"end_time"`
	StaffName           string                  `json:"staff_name"`
	VisitInfo           *VisitInfoResponseModel `json:"visit_info"`
	Title               string                  `json:"title"`
	Description         string                  `json:"description"`
	RecurringScheduleID string                  `json:"recurring_schedule_id"`
}

type ScheduleResponse struct {
	ID             string                  `json:"id"`
	ScheduleType   string                  `json:"schedule_type"`
	Date           common.Date             `json:"date"`
	StartTime      common.Time             `json:"start_time"`
	EndTime        common.Time             `json:"end_time"`
	IsOverTimeWork bool                    `json:"is_over_time_work"`
	StaffID        string                  `json:"staff_id"`
	StaffName      string                  `json:"staff_name"`
	VisitInfo      *VisitInfoResponseModel `json:"visit_info"`
	Title          string                  `json:"title"`
	Description    string                  `json:"description"`
	CancelReason   string                  `json:"cancel_reason"`
}

type RecurringScheduleResponse struct {
	ID             string                      `json:"id"`
	RecurringRule  *RecurringRuleResponseModel `json:"recurring_rule"`
	ScheduleType   string                      `json:"schedule_type"`
	Date           common.Date                 `json:"date"`
	StartTime      common.Time                 `json:"start_time"`
	EndTime        common.Time                 `json:"end_time"`
	IsOverTimeWork bool                        `json:"is_over_time_work"`
	StaffID        string                      `json:"staff_id"`
	StaffName      string                      `json:"staff_name"`
	VisitInfo      *VisitInfoResponseModel     `json:"visit_info"`
	Title          string                      `json:"title"`
	Description    string                      `json:"description"`
	ExclusionDates common.JSONSlice[int]       `json:"exclusion_dates"`
}

type ScheduleListResponse struct {
	Schedules          []*ScheduleResponse          `json:"schedules"`
	RecurringSchedules []*RecurringScheduleResponse `json:"recurring_schedules"`
}

type VisitInfoResponseModel struct {
	ID                string                       `json:"id"`
	PatientName       string                       `json:"patient_name"`
	AssignedStaffName string                       `json:"assigned_staff_name"`
	CompanionName     string                       `json:"companion_name"`
	Route             *RouteResponseModel          `json:"route"`
	ServiceCode       string                       `json:"service_code"`
	VisitCategories   []VisitCategoryResponseModel `json:"visit_categories"`
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
	EndDate     common.Date                       `json:"end_date"`
}
