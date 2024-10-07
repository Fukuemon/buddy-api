package position

import (
	positionDomain "api-buddy/domain/facility/position"
	"context"
	"time"
)

type FetchPositionsUseCase struct {
	positionRepository positionDomain.PositionRepository
}

func NewFetchPositionsUseCase(positionRepository positionDomain.PositionRepository) *FetchPositionsUseCase {
	return &FetchPositionsUseCase{
		positionRepository: positionRepository,
	}
}

type FetchPositionsUseCaseOutputDto struct {
	ID         string
	Name       string
	FacilityID string
	Policies   []PolicyDto
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (uc *FetchPositionsUseCase) Run(ctx context.Context, input string) ([]FetchPositionsUseCaseOutputDto, error) {
	positions, err := uc.positionRepository.FindByFacilityID(ctx, input)
	if err != nil {
		return nil, err
	}

	output := make([]FetchPositionsUseCaseOutputDto, 0, len(positions))
	for _, position := range positions {
		policiesDto := make([]PolicyDto, 0, len(position.Policies))
		for _, policy := range position.Policies {
			policiesDto = append(policiesDto, PolicyDto{
				ID:   policy.ID,
				Name: policy.Name,
			})
		}

		output = append(output, FetchPositionsUseCaseOutputDto{
			ID:         position.ID,
			Name:       position.Name,
			FacilityID: position.FacilityID,
			Policies:   policiesDto,
		})
	}

	return output, nil
}
