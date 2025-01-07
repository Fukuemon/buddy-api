package schedule_type

import "api-buddy/domain/schedule/schedule_type"

type ScheduleTypeResponse struct {
	ID   string                         `json:"id"`
	Name schedule_type.ScheduleTypeEnum `json:"name"`
}

type ScheduleTypeListResponse []ScheduleTypeResponse
