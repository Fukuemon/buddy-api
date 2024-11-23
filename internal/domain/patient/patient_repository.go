package patient

import (
	"context"

	"github.com/Fukuemon/go-pkg/query"
)

type PatientRepository interface {
	FindByFacilityID(ctx context.Context, facility_id string, filters []query.Filter, sort query.SortOption) ([]*Patient, error)
	FindByID(ctx context.Context, id string) (*Patient, error)
	Create(ctx context.Context, patient *Patient) error
}
