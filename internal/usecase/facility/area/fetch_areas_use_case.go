package area

import (
	areaDomain "api-buddy/domain/facility/area"
	"context"
)

type FetchAreaUseCase struct {
	areaRepository areaDomain.AreaRepository
}

func NewFetchAreaUseCase(areaRepository areaDomain.AreaRepository) *FetchAreaUseCase {
	return &FetchAreaUseCase{
		areaRepository: areaRepository,
	}
}

type FetchAreaUseCaseOutputDto struct {
	ID         string
	Name       string
	FacilityID string
	Addresses  []AddressDto
}

func (uc *FetchAreaUseCase) Run(ctx context.Context, facility_id string) ([]FetchAreaUseCaseOutputDto, error) {
	areas, err := uc.areaRepository.FindByFacilityID(ctx, facility_id)
	if err != nil {
		return nil, err
	}

	output := make([]FetchAreaUseCaseOutputDto, 0, len(areas))
	for _, area := range areas {
		addressesDto := make([]AddressDto, 0, len(area.Addresses))
		for _, address := range area.Addresses {
			addressesDto = append(addressesDto, AddressDto{
				ID:           address.ID,
				Prefecture:   address.Prefecture,
				City:         address.City,
				AddressLine1: address.AddressLine1,
				AddressLine2: address.AddressLine2,
				Latitude:     address.Latitude,
				Longitude:    address.Longitude,
			})
		}

		output = append(output, FetchAreaUseCaseOutputDto{
			ID:         area.ID,
			Name:       area.Name,
			FacilityID: area.FacilityID,
			Addresses:  addressesDto,
		})
	}

	return output, nil
}
