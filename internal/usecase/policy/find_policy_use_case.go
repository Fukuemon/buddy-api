package policy

import (
	policyDomain "api-buddy/domain/policy"
	"context"
	"time"
)

type FindPolicyUseCase struct {
	policyRepository policyDomain.PolicyRepository
}

func NewFindPolicyUseCase(policyRepository policyDomain.PolicyRepository) *FindPolicyUseCase {
	return &FindPolicyUseCase{
		policyRepository: policyRepository,
	}
}

type FindUseCaseOutputDto struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (uc *FindPolicyUseCase) Run(ctx context.Context, input string) (*FindUseCaseOutputDto, error) {
	policy, err := uc.policyRepository.FindByID(ctx, input)
	if err != nil {
		return nil, err
	}

	return &FindUseCaseOutputDto{
		ID:        policy.ID,
		Name:      policy.Name,
		CreatedAt: policy.CreatedAt,
		UpdatedAt: policy.UpdatedAt,
	}, nil
}
