package service_code

import (
	"api-buddy/domain/visit_info/service_code"
	"context"
)

type FetchServiceCodesUseCase struct {
	serviceCodeRepository service_code.ServiceCodeRepository
}

func NewFetchServiceCodesUseCase(serviceCodeRepository service_code.ServiceCodeRepository) *FetchServiceCodesUseCase {
	return &FetchServiceCodesUseCase{
		serviceCodeRepository: serviceCodeRepository,
	}
}

type FetchServiceCodesUseCaseOutputDto struct {
	ID                    string
	Code                  string
	ServiceTimeRangeStart int
	ServiceTimeRangeEnd   int
}

func (uc *FetchServiceCodesUseCase) Run(ctx context.Context) ([]FetchServiceCodesUseCaseOutputDto, error) {
	serviceCodes, err := uc.serviceCodeRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputDto []FetchServiceCodesUseCaseOutputDto
	for _, serviceCode := range serviceCodes {
		outputDto = append(outputDto, FetchServiceCodesUseCaseOutputDto{
			ID:                    serviceCode.ID,
			Code:                  serviceCode.Code,
			ServiceTimeRangeStart: serviceCode.ServiceTimeRangeStart,
			ServiceTimeRangeEnd:   serviceCode.ServiceTimeRangeEnd,
		})
	}

	return outputDto, nil
}
