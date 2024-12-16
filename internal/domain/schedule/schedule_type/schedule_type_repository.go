package schedule_type

import "context"

type ScheduleTypeRepository interface {
	FindAll(ctx context.Context, facility_id string) ([]*ScheduleType, error)
	FindByID(ctx context.Context, id string) (*ScheduleType, error)
}
