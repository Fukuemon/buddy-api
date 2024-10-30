package address

import (
	addressDomain "api-buddy/domain/address"
	"context"

	"github.com/Fukuemon/go-pkg/query"
)

type FetchAddressUseCase struct {
	addressRepository addressDomain.AddressRepository
}

func NewFetchAddressUseCase(addressRepository addressDomain.AddressRepository) *FetchAddressUseCase {
	return &FetchAddressUseCase{
		addressRepository: addressRepository,
	}
}

type FetchUseCaseOutputDto struct {
	ID           string
	ZipCode      string
	Prefecture   string
	City         string
	AddressLine1 string
	AddressLine2 string
	Latitude     float64
	Longitude    float64
}

type FetchUseCaseInputDto struct {
	ZipCode      string
	Prefecture   string
	City         string
	AddressLine1 string
	AddressLine2 string
}

func (uc *FetchAddressUseCase) Run(ctx context.Context, input FetchUseCaseInputDto) ([]*FetchUseCaseOutputDto, error) {
	var filters []query.Filter

	if input.ZipCode != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "zip_code", Value: input.ZipCode})
	}

	if input.Prefecture != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "prefecture", Value: input.Prefecture})
	}

	if input.City != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "city", Value: input.City})
	}

	if input.AddressLine1 != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "address_line1", Value: input.AddressLine1})
	}

	if input.AddressLine2 != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "address_line2", Value: input.AddressLine2})
	}

	addresses, err := uc.addressRepository.Fetch(ctx, filters)
	if err != nil {
		return nil, err
	}

	var output []*FetchUseCaseOutputDto
	for _, address := range addresses {
		output = append(output, &FetchUseCaseOutputDto{
			ID:           address.ID,
			ZipCode:      address.ZipCode,
			Prefecture:   address.Prefecture,
			City:         address.City,
			AddressLine1: address.AddressLine1,
			AddressLine2: address.AddressLine2,
			Latitude:     address.Latitude,
			Longitude:    address.Longitude,
		})
	}

	return output, nil
}
