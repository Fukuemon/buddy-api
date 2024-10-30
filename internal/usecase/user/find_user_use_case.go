package user

import (
	facilityDomain "api-buddy/domain/facility"
	areaDomain "api-buddy/domain/facility/area"
	departmentDomain "api-buddy/domain/facility/department"
	positionDomain "api-buddy/domain/facility/position"
	teamDomain "api-buddy/domain/facility/team"
	policyDomain "api-buddy/domain/policy"
	userDomain "api-buddy/domain/user"
	"context"
	"time"
)

type FindUserUseCase struct {
	userRepository userDomain.UserRepository
}

func NewFindUserUseCase(userRepository userDomain.UserRepository) *FindUserUseCase {
	return &FindUserUseCase{
		userRepository: userRepository,
	}
}

type FindUserUseCaseOutputDto struct {
	ID          string
	Username    string
	Position    positionDomain.Position
	Team        teamDomain.Team
	Facility    facilityDomain.Facility
	Department  departmentDomain.Department
	Area        areaDomain.Area
	Policies    []*policyDomain.Policy
	Email       *string
	PhoneNumber *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (uc *FindUserUseCase) Run(ctx context.Context, input string) (*FindUserUseCaseOutputDto, error) {
	user, err := uc.userRepository.FindByID(ctx, input)
	if err != nil {
		return nil, err
	}

	return &FindUserUseCaseOutputDto{
		ID:          user.ID,
		Username:    user.Username,
		Position:    *user.Position,
		Team:        *user.Team,
		Facility:    *user.Facility,
		Department:  *user.Department,
		Area:        *user.Area,
		Policies:    user.Policies,
		Email:       &user.Email,
		PhoneNumber: &user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}
