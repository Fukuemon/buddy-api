package schedule

import (
	"context"

	"github.com/Fukuemon/go-pkg/query"
)

type ScheduleRepository interface {
	FindByFacilityID(ctx context.Context, facility_id string, filters []query.Filter, sort query.SortOption) ([]*Schedule, error)
	FindByID(ctx context.Context, id string) (*Schedule, error)
	Create(ctx context.Context, schedule *Schedule) error
}
