package schedule_cancel

import "context"

type ScheduleCancelRepository interface {
	FindByID(ctx context.Context, id string) (*ScheduleCancel, error)
	Create(ctx context.Context, schedule_cancel *ScheduleCancel) error
}
