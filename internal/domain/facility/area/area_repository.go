package area

import "context"

type AreaRepository interface {
	Create(ctx context.Context, area *Area) error
	FindByID(ctx context.Context, id string) (*Area, error)
	FindByFacilityID(ctx context.Context, facilityID string) ([]*Area, error)
}
