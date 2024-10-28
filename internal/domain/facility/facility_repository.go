package facility

import (
	"context"
)

type FacilityRepository interface {
	Create(ctx context.Context, facility *Facility) error
	FindByID(ctx context.Context, id string) (*Facility, error)
}
