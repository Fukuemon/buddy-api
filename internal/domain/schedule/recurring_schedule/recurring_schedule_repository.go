package recurring_schedule

import "github.com/Fukuemon/go-pkg/query"

type RecurringScheduleRepository interface {
	FindByFacilityID(facility_id string, filters []query.Filter, sort query.SortOption) ([]*RecurringSchedule, error)
	FindByID(id string) (*RecurringSchedule, error)
	Create(recurring_schedule *RecurringSchedule) error
}
