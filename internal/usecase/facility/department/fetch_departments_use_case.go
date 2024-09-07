package department

import (
	departmentDomain "api-buddy/domain/facility/department"
	"context"
	"time"
)

type FetchDepartmentsUseCase struct {
	departmentRepository departmentDomain.DepartmentRepository
}

func NewFetchDepartmentsUseCase(departmentRepository departmentDomain.DepartmentRepository) *FetchDepartmentsUseCase {
	return &FetchDepartmentsUseCase{
		departmentRepository: departmentRepository,
	}
}

type FetchDepartmentsUseCaseOutputDto struct {
	ID         string
	Name       string
	FacilityID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func (uc *FetchDepartmentsUseCase) Run(ctx context.Context, input string) ([]FetchDepartmentsUseCaseOutputDto, error) {
	var departments []*departmentDomain.Department
	var err error

	if input != "" {
		departments, err = uc.departmentRepository.FindByFacilityID(ctx, input)
	} else {
		departments, err = uc.departmentRepository.FindAll(ctx)
	}

	if err != nil {
		return nil, err
	}

	output := make([]FetchDepartmentsUseCaseOutputDto, 0, len(departments))
	for _, department := range departments {
		output = append(output, FetchDepartmentsUseCaseOutputDto{
			ID:         department.ID,
			Name:       department.Name,
			FacilityID: department.FacilityID,
			CreatedAt:  department.CreatedAt,
			UpdatedAt:  department.UpdatedAt,
			DeletedAt:  department.DeletedAt,
		})
	}

	return output, nil
}
