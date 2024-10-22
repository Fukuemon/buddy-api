package team

import (
	teamDomain "api-buddy/domain/facility/team"
	"context"
)

type CreateTeamUseCase struct {
	teamRepository teamDomain.TeamRepository
}

func NewCreateTeamUseCase(teamRepository teamDomain.TeamRepository) *CreateTeamUseCase {
	return &CreateTeamUseCase{
		teamRepository: teamRepository,
	}
}

type CreateUseCaseInputDto struct {
	Name       string
	FacilityID string
}

type CreateUseCaseOutputDto struct {
	ID         string
	Name       string
	FacilityID string
}

func (uc *CreateTeamUseCase) Run(ctx context.Context, input CreateUseCaseInputDto) (*CreateUseCaseOutputDto, error) {
	team, err := teamDomain.NewTeam(input.Name, input.FacilityID)
	if err != nil {
		return nil, err
	}

	err = uc.teamRepository.Create(ctx, team)
	if err != nil {
		return nil, err
	}

	return &CreateUseCaseOutputDto{
		ID:         team.ID,
		Name:       team.Name,
		FacilityID: team.FacilityID,
	}, nil
}
