package address

import (
	"context"

	"github.com/Fukuemon/go-pkg/query"
)

type AddressRepository interface {
	FindByID(ctx context.Context, id string) (*Address, error)
	Create(ctx context.Context, address *Address) error
	FindByAreaID(ctx context.Context, areaID string) ([]*Address, error)
	FindByIDs(ctx context.Context, ids []string) ([]*Address, error)
	Fetch(ctx context.Context, filters []query.Filter) ([]*Address, error)
}
