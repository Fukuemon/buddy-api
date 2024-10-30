package user

import (
	facilityDomain "api-buddy/domain/facility"
	areaDomain "api-buddy/domain/facility/area"
	departmentDomain "api-buddy/domain/facility/department"
	positionDomain "api-buddy/domain/facility/position"
	teamDomain "api-buddy/domain/facility/team"
	policyDomain "api-buddy/domain/policy"
	userDomain "api-buddy/domain/user"
	"api-buddy/infrastructure/aws/cognito"
	"context"
)

type CreateUserUseCase struct {
	userRepository       userDomain.UserRepository
	facilityRepository   facilityDomain.FacilityRepository
	departmentRepository departmentDomain.DepartmentRepository
	positionRepository   positionDomain.PositionRepository
	teamRepository       teamDomain.TeamRepository
	areaRepository       areaDomain.AreaRepository
}

func NewCreateUserUseCase(
	userRepository userDomain.UserRepository,
	facilityRepository facilityDomain.FacilityRepository,
	departmentRepository departmentDomain.DepartmentRepository,
	positionRepository positionDomain.PositionRepository,
	teamRepository teamDomain.TeamRepository,
	areaRepository areaDomain.AreaRepository,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository:       userRepository,
		facilityRepository:   facilityRepository,
		departmentRepository: departmentRepository,
		positionRepository:   positionRepository,
		teamRepository:       teamRepository,
		areaRepository:       areaRepository,
	}
}

type CreateUserUseCaseInputDto struct {
	Username     string
	Password     string
	FacilityID   string
	DepartmentID string
	PositionID   string
	TeamID       string
	AreaID       string
	*userDomain.Option
}

type CreateUserUseCaseOutputDto struct {
	ID          string
	Username    string
	Facility    facilityDomain.Facility
	Department  departmentDomain.Department
	Position    positionDomain.Position
	Team        teamDomain.Team
	Area        areaDomain.Area
	Policies    []*policyDomain.Policy
	Email       *string
	PhoneNumber *string
}

func (uc *CreateUserUseCase) Run(ctx context.Context, input CreateUserUseCaseInputDto) (*CreateUserUseCaseOutputDto, error) {

	// 各Repositoryを使って、IDから各エンティティを取得する
	facility, err := uc.facilityRepository.FindByID(ctx, input.FacilityID)
	if err != nil {
		return nil, err
	}

	department, err := uc.departmentRepository.FindByID(ctx, input.DepartmentID)
	if err != nil {
		return nil, err
	}

	position, err := uc.positionRepository.FindByID(ctx, input.PositionID)
	if err != nil {
		return nil, err
	}

	team, err := uc.teamRepository.FindByID(ctx, input.TeamID)
	if err != nil {
		return nil, err
	}

	area, err := uc.areaRepository.FindByID(ctx, input.AreaID)
	if err != nil {
		return nil, err
	}

	// Userエンティティを生成する
	user, err := userDomain.NewUser(
		input.Username,
		position,
		team,
		facility,
		department,
		area,
		position.Policies,
		input.Option,
	)

	if err != nil {
		return nil, err
	}

	// Cognitoへのユーザー登録
	cognitoUserId, err := cognito.Actions.SignUp(&cognito.CognitoSignUpRequest{
		Username:    user.Username,
		Password:    input.Password,
		Email:       &user.Email,
		PhoneNumber: &user.PhoneNumber,
	})

	if err != nil || cognitoUserId == nil {
		return nil, err
	}

	err = uc.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &CreateUserUseCaseOutputDto{
		ID:          user.ID,
		Username:    user.Username,
		Facility:    *facility,
		Department:  *department,
		Position:    *position,
		Team:        *team,
		Area:        *area,
		Policies:    user.Policies,
		Email:       &user.Email,
		PhoneNumber: &user.PhoneNumber,
	}, nil

}
