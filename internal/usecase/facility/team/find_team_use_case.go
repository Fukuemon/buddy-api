package team

import (
	teamDomain "api-buddy/domain/facility/team"
	"context"
	"time"
)

type FindTeamUseCase struct {
	teamRepository teamDomain.TeamRepository
}

func NewFindTeamUseCase(teamRepository teamDomain.TeamRepository) *FindTeamUseCase {
	return &FindTeamUseCase{
		teamRepository: teamRepository,
	}
}

type FindTeamUseCaseOutputDto struct {
	ID         string
	Name       string
	FacilityID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func (uc *FindTeamUseCase) Run(ctx context.Context, input string) (*FindTeamUseCaseOutputDto, error) {
	team, err := uc.teamRepository.FindByID(ctx, input)
	if err != nil {
		return nil, err
	}

	return &FindTeamUseCaseOutputDto{
		ID:         team.ID,
		Name:       team.Name,
		FacilityID: team.FacilityID,
		CreatedAt:  team.CreatedAt,
		UpdatedAt:  team.UpdatedAt,
		DeletedAt:  team.DeletedAt,
	}, nil
}
