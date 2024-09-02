package department

import "context"

type DepartmentRepository interface {
	Create(ctx context.Context, department *Department) error
	FindByID(ctx context.Context, id string) (*Department, error)
	FindByFacilityID(ctx context.Context, facilityID string) ([]*Department, error)
}
