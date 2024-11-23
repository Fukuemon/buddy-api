package recurring_rule

import (
	"context"

	"github.com/Fukuemon/go-pkg/query"
)

type RecurringRuleRepository interface {
	FindByFacilityID(ctx context.Context, facility_id string, filters []query.Filter, sort query.SortOption) ([]*RecurringRule, error)
	FindByID(ctx context.Context, id string) (*RecurringRule, error)
	Create(ctx context.Context, recurring_rule *RecurringRule) error
}
