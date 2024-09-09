package policy

import (
	policyDomain "api-buddy/domain/policy"
	"context"
)

type CreatePolicyUseCase struct {
	policyRepository policyDomain.PolicyRepository
}

func NewCreatePolicyUseCase(policyRepository policyDomain.PolicyRepository) *CreatePolicyUseCase {
	return &CreatePolicyUseCase{
		policyRepository: policyRepository,
	}
}

type CreateUseCaseInputDto struct {
	Name string
}

type CreateUseCaseOutputDto struct {
	ID   string
	Name string
}

func (uc *CreatePolicyUseCase) Run(ctx context.Context, input *CreateUseCaseInputDto) (*CreateUseCaseOutputDto, error) {
	policy, err := policyDomain.NewPolicy(input.Name)
	if err != nil {
		return nil, err
	}

	if err := uc.policyRepository.Create(ctx, policy); err != nil {
		return nil, err
	}

	return &CreateUseCaseOutputDto{
		ID:   policy.ID,
		Name: policy.Name,
	}, nil
}
