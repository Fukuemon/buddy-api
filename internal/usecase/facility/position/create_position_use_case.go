package position

import (
	positionDomain "api-buddy/domain/facility/position"
	policyDomain "api-buddy/domain/policy"
	"context"
)

type CreatePositionUseCase struct {
	positionRepository positionDomain.PositionRepository
	policyRepository   policyDomain.PolicyRepository
}

func NewCreatePositionUseCase(positionRepository positionDomain.PositionRepository, policyRepository policyDomain.PolicyRepository) *CreatePositionUseCase {
	return &CreatePositionUseCase{
		positionRepository: positionRepository,
		policyRepository:   policyRepository,
	}
}

type CreateUseCaseInputDto struct {
	Name       string   `json:"name"`
	FacilityID string   `json:"facility_id"`
	PolicyIDs  []string `json:"policy_ids"`
}

type CreateUseCaseOutputDto struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	FacilityID string      `json:"facility_id"`
	Policies   []PolicyDto `json:"policies"`
}

func (uc *CreatePositionUseCase) Create(ctx context.Context, input CreateUseCaseInputDto) (*CreateUseCaseOutputDto, error) {
	policies, err := uc.policyRepository.FindByIDs(ctx, input.PolicyIDs)
	if err != nil {
		return nil, err
	}

	position, err := positionDomain.NewPosition(input.Name, input.FacilityID, policies)
	if err != nil {
		return nil, err
	}

	err = uc.positionRepository.Create(ctx, position)
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

	return &CreateUseCaseOutputDto{
		ID:         position.ID,
		Name:       position.Name,
		FacilityID: position.FacilityID,
		Policies:   policiesDto,
	}, nil
}
