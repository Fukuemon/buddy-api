package recurring_schedule

import (
	"context"

	"github.com/Fukuemon/go-pkg/query"
)

type RecurringScheduleRepository interface {
	FindByFacilityID(ctx context.Context, facility_id string, filters []query.Filter, sort query.SortOption) ([]*RecurringSchedule, error)
	FindByID(ctx context.Context, id string) (*RecurringSchedule, error)
	Create(ctx context.Context, recurring_schedule *RecurringSchedule) error
}
