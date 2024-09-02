package policy

import "context"

type PolicyRepository interface {
	Create(ctx context.Context, policy *Policy) error
	FindByID(ctx context.Context, id string) (*Policy, error)
}
