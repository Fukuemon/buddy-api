package department

import (
	departmentDomain "api-buddy/domain/facility/department"
	"context"
)

type CreateDepartmentUseCase struct {
	departmentRepository departmentDomain.DepartmentRepository
}

func NewCreateDepartmentUseCase(departmentRepository departmentDomain.DepartmentRepository) *CreateDepartmentUseCase {
	return &CreateDepartmentUseCase{
		departmentRepository: departmentRepository,
	}
}

type CreateUseCaseInputDto struct {
	Name       string
	FacilityID string
}

type CreateUseCaseOutputDto struct {
	ID   string
	Name string
}

func (uc *CreateDepartmentUseCase) Run(ctx context.Context, input CreateUseCaseInputDto) (*CreateUseCaseOutputDto, error) {
	department, err := departmentDomain.NewDepartment(input.Name, input.FacilityID)
	if err != nil {
		return nil, err
	}

	err = uc.departmentRepository.Create(ctx, department)
	if err != nil {
		return nil, err
	}

	return &CreateUseCaseOutputDto{
		ID:   department.ID,
		Name: department.Name,
	}, nil
}
