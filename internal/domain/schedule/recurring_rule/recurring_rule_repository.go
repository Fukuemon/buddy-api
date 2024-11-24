package recurring_rule

import (
	"context"
)

type RecurringRuleRepository interface {
	FindByFacilityID(ctx context.Context, facility_id string) ([]*RecurringRule, error)
	FindByID(ctx context.Context, id string) (*RecurringRule, error)
	Create(ctx context.Context, recurring_rule *RecurringRule) error
}
