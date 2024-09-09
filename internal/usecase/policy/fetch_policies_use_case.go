package policy

import (
	policyDomain "api-buddy/domain/policy"
	"context"
	"time"
)

type FetchPoliciesUseCase struct {
	policyRepository policyDomain.PolicyRepository
}

func NewFetchPoliciesUseCase(policyRepository policyDomain.PolicyRepository) *FetchPoliciesUseCase {
	return &FetchPoliciesUseCase{
		policyRepository: policyRepository,
	}
}

type FetchUseCaseOutputDto struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (uc *FetchPoliciesUseCase) Run(ctx context.Context) ([]FetchUseCaseOutputDto, error) {
	policies, err := uc.policyRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	policiesDto := make([]FetchUseCaseOutputDto, 0, len(policies))
	for _, policy := range policies {
		policiesDto = append(policiesDto, FetchUseCaseOutputDto{
			ID:        policy.ID,
			Name:      policy.Name,
			CreatedAt: policy.CreatedAt,
			UpdatedAt: policy.UpdatedAt,
			DeletedAt: policy.DeletedAt,
		})
	}

	return policiesDto, nil
}
