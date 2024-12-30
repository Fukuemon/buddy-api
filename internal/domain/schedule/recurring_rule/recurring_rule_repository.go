package recurring_rule

import (
	"context"

	"gorm.io/gorm"
)

type RecurringRuleRepository interface {
	FindByFacilityID(ctx context.Context, facility_id string) ([]*RecurringRule, error)
	FindByID(ctx context.Context, id string) (*RecurringRule, error)
	Create(ctx context.Context, tx *gorm.DB, recurring_rule *RecurringRule) error
}
