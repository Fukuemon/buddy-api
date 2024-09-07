package department

import (
	departmentDomain "api-buddy/domain/facility/department"
	"context"
	"time"
)

type FindDepartmentUseCase struct {
	departmentRepository departmentDomain.DepartmentRepository
}

func NewFindDepartmentUseCase(departmentRepository departmentDomain.DepartmentRepository) *FindDepartmentUseCase {
	return &FindDepartmentUseCase{
		departmentRepository: departmentRepository,
	}
}

type FindUseCaseOutputDto struct {
	ID         string
	Name       string
	FacilityID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func (uc *FindDepartmentUseCase) Run(ctx context.Context, input string) (*FindUseCaseOutputDto, error) {
	department, err := uc.departmentRepository.FindByID(ctx, input)
	if err != nil {
		return nil, err
	}

	return &FindUseCaseOutputDto{
		ID:         department.ID,
		Name:       department.Name,
		FacilityID: department.FacilityID,
		CreatedAt:  department.CreatedAt,
		UpdatedAt:  department.UpdatedAt,
		DeletedAt:  department.DeletedAt,
	}, nil
}
