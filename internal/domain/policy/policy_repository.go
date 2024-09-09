package policy

import "context"

type PolicyRepository interface {
	Create(ctx context.Context, policy *Policy) error
	FindByID(ctx context.Context, id string) (*Policy, error)
	FindByIDs(ctx context.Context, ids []string) ([]*Policy, error)
	FindByPositionID(ctx context.Context, positionID string) ([]*Policy, error)
	FindAll(ctx context.Context) ([]*Policy, error)
}
