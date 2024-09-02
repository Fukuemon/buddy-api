package position

import "context"

type PositionRepository interface {
	Create(ctx context.Context, position *Position) error
	FindByID(ctx context.Context, id string) (*Position, error)
	FindByFacilityID(ctx context.Context, facilityID string) ([]*Position, error)
}
