package patient

import (
	addressDomain "api-buddy/domain/address"
	patientDomain "api-buddy/domain/patient"
	"context"

	"github.com/Fukuemon/go-pkg/query"
)

type FetchPatientUseCase struct {
	patientRepository patientDomain.PatientRepository
}

func NewFetchPatientUseCase(patientRepository patientDomain.PatientRepository) *FetchPatientUseCase {
	return &FetchPatientUseCase{
		patientRepository: patientRepository,
	}
}

// Output DTO
type FetchPatientUseCaseOutputDto struct {
	ID              string
	Name            string
	PreferredTime   *patientDomain.PreferredTime
	PreferredGender *patientDomain.PreferredGender
	AssignedStaff   string
	Address         *addressDomain.Address
	Area            string
}

// Input DTO
type FetchPatientUseCaseInputDto struct {
	Name            string
	PreferredTime   *patientDomain.PreferredTime
	PreferredGender *patientDomain.PreferredGender
	AssignedStaff   string
	ZipCode         string
	Area            string
	SortField       string
	SortOrder       string
}

func (uc *FetchPatientUseCase) Run(ctx context.Context, facility_id string, input FetchPatientUseCaseInputDto) ([]FetchPatientUseCaseOutputDto, error) {
	// フィルタリング条件を定義
	var filters []query.Filter

	if input.Name != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "name", Value: input.Name, RelationMapping: patientDomain.PatientRelationMappings})
	}
	if input.PreferredTime != nil {
		filters = append(filters, &query.ByFieldFilter{Field: "preferred_time", Value: input.PreferredTime.String(), RelationMapping: patientDomain.PatientRelationMappings})
	}
	if input.PreferredGender != nil {
		filters = append(filters, &query.ByFieldFilter{Field: "preferred_gender", Value: input.PreferredGender.String(), RelationMapping: patientDomain.PatientRelationMappings})
	}

	if input.AssignedStaff != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "assigned_staff", Value: input.AssignedStaff, RelationMapping: patientDomain.PatientRelationMappings})
	}

	if input.ZipCode != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "address", Value: input.ZipCode, RelationMapping: patientDomain.PatientRelationMappings})
	}

	if input.Area != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "area", Value: input.Area, RelationMapping: patientDomain.PatientRelationMappings})
	}

	sortOption := query.SortOption{
		Field: input.SortField,
		Order: input.SortOrder,
	}

	patients, err := uc.patientRepository.FindByFacilityID(ctx, facility_id, filters, sortOption)
	if err != nil {
		return nil, err
	}

	output := make([]FetchPatientUseCaseOutputDto, 0, len(patients))
	for _, patient := range patients {
		output = append(output, FetchPatientUseCaseOutputDto{
			ID:              patient.ID,
			Name:            patient.Name,
			PreferredTime:   patient.PreferredTime,
			PreferredGender: patient.PreferredGender,
			AssignedStaff:   patient.AssignedStaff.Username,
			Address:         patient.Address,
			Area:            patient.Area.Name,
		})
	}

	return output, nil
}
