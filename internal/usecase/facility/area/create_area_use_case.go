package area

import (
	addressDomain "api-buddy/domain/address"
	areaDomain "api-buddy/domain/facility/area"
	"context"
)

type CreateAreaUseCase struct {
	areaRepository    areaDomain.AreaRepository
	addressRepository addressDomain.AddressRepository
}

func NewCreateAreaUseCase(areaRepository areaDomain.AreaRepository, addressRepository addressDomain.AddressRepository) *CreateAreaUseCase {
	return &CreateAreaUseCase{
		areaRepository:    areaRepository,
		addressRepository: addressRepository,
	}
}

type CreateAreaUseCaseInputDto struct {
	Name       string
	FacilityID string
	AddressIDs []string
}

type CreateAreaUseCaseOutputDto struct {
	ID         string
	Name       string
	FacilityID string
	AddressIDs []AddressDto
}

type AddressDto struct {
	ID           string
	Prefecture   string
	City         string
	AddressLine1 string
	AddressLine2 string
	Latitude     float64
	Longitude    float64
}

func (uc *CreateAreaUseCase) Run(ctx context.Context, input CreateAreaUseCaseInputDto) (*CreateAreaUseCaseOutputDto, error) {
	addresses, err := uc.addressRepository.FindByIDs(ctx, input.AddressIDs)
	if err != nil {
		return nil, err
	}

	area, err := areaDomain.NewArea(
		input.Name,
		input.FacilityID,
		addresses,
	)
	if err != nil {
		return nil, err
	}

	err = uc.areaRepository.Create(ctx, area)
	if err != nil {
		return nil, err
	}

	return &CreateAreaUseCaseOutputDto{
		ID:         area.ID,
		Name:       area.Name,
		FacilityID: area.FacilityID,
		AddressIDs: toAddressDto(addresses),
	}, nil
}

func toAddressDto(addresses []*addressDomain.Address) []AddressDto {
	var addressDtos []AddressDto
	for _, address := range addresses {
		addressDtos = append(addressDtos, AddressDto{
			ID:           address.ID,
			Prefecture:   address.Prefecture,
			City:         address.City,
			AddressLine1: address.AddressLine1,
			AddressLine2: address.AddressLine2,
			Latitude:     address.Latitude,
			Longitude:    address.Longitude,
		})
	}
	return addressDtos
}
