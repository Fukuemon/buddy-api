package team

import (
	teamDomain "api-buddy/domain/facility/team"
	"context"
	"time"
)

type FetchTeamsUseCase struct {
	teamRepository teamDomain.TeamRepository
}

func NewFetchTeamsUseCase(teamRepository teamDomain.TeamRepository) *FetchTeamsUseCase {
	return &FetchTeamsUseCase{
		teamRepository: teamRepository,
	}
}

type FetchTeamsUseCaseOutputDto struct {
	ID         string
	Name       string
	FacilityID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (uc *FetchTeamsUseCase) Run(ctx context.Context, input string) ([]FetchTeamsUseCaseOutputDto, error) {
	var teams []*teamDomain.Team
	var err error

	if input != "" {
		teams, err = uc.teamRepository.FindByFacilityID(ctx, input)
	} else {
		teams, err = uc.teamRepository.FindAll(ctx)
	}

	if err != nil {
		return nil, err
	}

	output := make([]FetchTeamsUseCaseOutputDto, 0, len(teams))
	for _, team := range teams {
		output = append(output, FetchTeamsUseCaseOutputDto{
			ID:         team.ID,
			Name:       team.Name,
			FacilityID: team.FacilityID,
			CreatedAt:  team.CreatedAt,
			UpdatedAt:  team.UpdatedAt,
		})

	}

	return output, nil
}
