package address

import (
	addressDomain "api-buddy/domain/address"
	"context"
	"time"
)

type FindAddressUseCase struct {
	addressRepository addressDomain.AddressRepository
}

func NewFindAddressUseCase(addressRepository addressDomain.AddressRepository) *FindAddressUseCase {
	return &FindAddressUseCase{
		addressRepository: addressRepository,
	}
}

type FindUseCaseOutputDto struct {
	ID           string
	ZipCode      string
	Prefecture   string
	City         string
	AddressLine1 string
	AddressLine2 string
	Latitude     float64
	Longitude    float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (uc *FindAddressUseCase) Run(ctx context.Context, input string) (*FindUseCaseOutputDto, error) {
	address, err := uc.addressRepository.FindByID(ctx, input)
	if err != nil {
		return nil, err
	}

	return &FindUseCaseOutputDto{
		ID:           address.ID,
		ZipCode:      address.ZipCode,
		Prefecture:   address.Prefecture,
		City:         address.City,
		AddressLine1: address.AddressLine1,
		AddressLine2: address.AddressLine2,
		Latitude:     address.Latitude,
		Longitude:    address.Longitude,
		CreatedAt:    address.CreatedAt,
		UpdatedAt:    address.UpdatedAt,
	}, nil
}
