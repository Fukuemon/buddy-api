package user

import (
	userDomain "api-buddy/domain/user"
	"context"

	"github.com/Fukuemon/go-pkg/query"
)

type FetchUsersUseCase struct {
	userRepository userDomain.UserRepository
}

func NewFetchUsersUseCase(userRepository userDomain.UserRepository) *FetchUsersUseCase {
	return &FetchUsersUseCase{
		userRepository: userRepository,
	}
}

// Output DTO
type FetchUsersUseCaseOutputDto struct {
	ID         string
	Username   string
	Position   string
	Team       string
	Department string
	Area       string
}

// Input DTO
type FetchUsersUseCaseInputDto struct {
	Username   string
	Position   string
	Team       string
	Department string
	Area       string
	SortField  string
	SortOrder  string
}

func (uc *FetchUsersUseCase) Run(ctx context.Context, facility_id string, input FetchUsersUseCaseInputDto) ([]FetchUsersUseCaseOutputDto, error) {
	// フィルタリング条件を定義
	var filters []query.Filter

	if input.Username != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "username", Value: input.Username, RelationMapping: userDomain.UserRelationMappings})
	}
	if input.Position != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "position", Value: input.Position, RelationMapping: userDomain.UserRelationMappings})
	}
	if input.Team != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "team", Value: input.Team, RelationMapping: userDomain.UserRelationMappings})
	}
	if input.Department != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "department", Value: input.Department, RelationMapping: userDomain.UserRelationMappings})
	}
	if input.Area != "" {
		filters = append(filters, &query.ByFieldFilter{Field: "area", Value: input.Area, RelationMapping: userDomain.UserRelationMappings})
	}

	sortOption := query.SortOption{
		Field: input.SortField,
		Order: input.SortOrder,
	}

	users, err := uc.userRepository.FindByFacilityID(ctx, facility_id, filters, sortOption)
	if err != nil {
		return nil, err
	}

	output := make([]FetchUsersUseCaseOutputDto, 0, len(users))
	for _, user := range users {
		output = append(output, FetchUsersUseCaseOutputDto{
			ID:         user.ID,
			Username:   user.Username,
			Position:   user.Position.Name,
			Team:       user.Team.Name,
			Department: user.Department.Name,
			Area:       user.Area.Name,
		})
	}

	return output, nil
}
