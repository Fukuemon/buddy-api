package team

import "context"

type TeamRepository interface {
	Create(ctx context.Context, team *Team) error
	FindByID(ctx context.Context, id string) (*Team, error)
	FindByFacilityID(ctx context.Context, facilityID string) ([]*Team, error)
	FindAll(ctx context.Context) ([]*Team, error)
}
