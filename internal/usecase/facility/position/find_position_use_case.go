package position

import (
	positionDomain "api-buddy/domain/facility/position"
	"context"
	"time"
)

type FindPositionUseCase struct {
	positionRepository positionDomain.PositionRepository
}

func NewFindPositionUseCase(positionRepository positionDomain.PositionRepository) *FindPositionUseCase {
	return &FindPositionUseCase{
		positionRepository: positionRepository,
	}
}

type FindUseCaseOutputDto struct {
	ID         string
	Name       string
	FacilityID string
	Policies   []PolicyDto
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PolicyDto struct {
	ID   string
	Name string
}

func (uc *FindPositionUseCase) Run(ctx context.Context, input string) (*FindUseCaseOutputDto, error) {
	position, err := uc.positionRepository.FindByID(ctx, input)
	if err != nil {
		return nil, err
	}

	policiesDto := make([]PolicyDto, 0, len(position.Policies))
	for _, policy := range position.Policies {
		policiesDto = append(policiesDto, PolicyDto{
			ID:   policy.ID,
			Name: policy.Name,
		})
	}

	return &FindUseCaseOutputDto{
		ID:         position.ID,
		Name:       position.Name,
		FacilityID: position.FacilityID,
		Policies:   policiesDto,
	}, nil
}
