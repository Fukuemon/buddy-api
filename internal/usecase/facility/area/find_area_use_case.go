package area

import (
	areaDomain "api-buddy/domain/facility/area"
	"context"
)

type FindAreaUseCase struct {
	AreaRepository areaDomain.AreaRepository
}

func NewFindAreaUseCase(areaRepository areaDomain.AreaRepository) *FindAreaUseCase {
	return &FindAreaUseCase{
		AreaRepository: areaRepository,
	}
}

type FindAreaUseCaseOutputDto struct {
	ID         string
	Name       string
	FacilityID string
	Addresses  []AddressDto
	CreatedAt  string
	UpdatedAt  string
}

func (uc *FindAreaUseCase) Run(ctx context.Context, input string) (*FindAreaUseCaseOutputDto, error) {
	area, err := uc.AreaRepository.FindByID(ctx, input)
	if err != nil {
		return nil, err
	}

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

	return &FindAreaUseCaseOutputDto{
		ID:         area.ID,
		Name:       area.Name,
		FacilityID: area.FacilityID,
		Addresses:  addressesDto,
		CreatedAt:  area.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  area.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
