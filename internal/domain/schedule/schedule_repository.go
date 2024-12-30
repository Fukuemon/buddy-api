package schedule

import (
	"context"

	"github.com/Fukuemon/go-pkg/query"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	FindByFacilityID(ctx context.Context, facility_id string, filters []query.Filter, sort query.SortOption) ([]*Schedule, error)
	FindByID(ctx context.Context, id string) (*Schedule, error)
	Create(ctx context.Context, tx *gorm.DB, schedule *Schedule) error
}
