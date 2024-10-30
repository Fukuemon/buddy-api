package address

import (
	addressDomain "api-buddy/domain/address"
	"context"
	"time"
)

type CreateAddressUseCase struct {
	addressRepository addressDomain.AddressRepository
}

func NewCreateAddressUseCase(addressRepository addressDomain.AddressRepository) *CreateAddressUseCase {
	return &CreateAddressUseCase{
		addressRepository: addressRepository,
	}
}

type CreateUseCaseInputDto struct {
	ZipCode      string
	Prefecture   string
	City         string
	AddressLine1 string
	AddressLine2 string
	Latitude     float64
	Longitude    float64
}

type CreateUseCaseOutputDto struct {
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

func (uc *CreateAddressUseCase) Run(ctx context.Context, input CreateUseCaseInputDto) (*CreateUseCaseOutputDto, error) {
	address, err := addressDomain.NewAddress(
		input.ZipCode,
		input.Prefecture,
		input.City,
		input.AddressLine1,
		input.AddressLine2,
		input.Latitude,
		input.Longitude,
	)
	if err != nil {
		return nil, err
	}

	err = uc.addressRepository.Create(ctx, address)
	if err != nil {
		return nil, err
	}

	return &CreateUseCaseOutputDto{
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
