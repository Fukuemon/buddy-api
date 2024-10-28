package user

import (
	"context"

	"github.com/Fukuemon/go-pkg/query"
)

type UserRepository interface {
	FindByFacilityID(ctx context.Context, facility_id string, filters []query.Filter, sort query.SortOption) ([]*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) error
}
