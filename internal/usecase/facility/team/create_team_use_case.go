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
	Name       string `json:"name"`
	FacilityID string `json:"facility_id"`
}

type CreateUseCaseOutputDto struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FacilityID string `json:"facility_id"`
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
