package team

import (
	"api-buddy/presentation/settings"
	"api-buddy/usecase/facility/team"

	"github.com/gin-gonic/gin"
)

type handler struct {
	createTeamUseCase *team.CreateTeamUseCase
	findTeamUseCase   *team.FindTeamUseCase
	fetchTeamsUseCase *team.FetchTeamsUseCase
}

func NewHandler(createTeamUseCase *team.CreateTeamUseCase, findTeamUseCase *team.FindTeamUseCase, fetchTeamsUseCase *team.FetchTeamsUseCase) *handler {
	return &handler{
		createTeamUseCase: createTeamUseCase,
		findTeamUseCase:   findTeamUseCase,
		fetchTeamsUseCase: fetchTeamsUseCase,
	}
}

// Create godoc
// @Summary      施設に紐づくチームを作成する
// @Tags         Team
// @Accept       json
// @Produce      json
// @Param        request body      CreateTeamRequest  true  "Create Team Request"
// @Success      201      {object} TeamResponse
// @Failure      400      {object} ErrorResponse
// @Failure      500      {object} ErrorResponse
// @Router       /facilities/{facility_id}/teams [post]
func (h handler) CreateByFacilityId(ctx *gin.Context) {
	facilityID := ctx.Param("facility_id")
	var params CreateTeamRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		settings.ReturnBadRequest(ctx, err)
		return
	}

	input := team.CreateUseCaseInputDto{
		Name:       params.Name,
		FacilityID: facilityID,
	}
	output, err := h.createTeamUseCase.Run(ctx, input)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	response := CreateTeamResponse{
		ID:   output.ID,
		Name: output.Name,
	}

	settings.ReturnStatusCreated(ctx, response)
}

// FindByID godoc
// @Summary      単一のチーム取得する
// @Tags         Team
// @Accept       json
// @Produce      json
// @Param        team_id path string true "Team ID"
// @Success      200      {object} TeamResponse
// @Failure      400      {object} ErrorResponse
// @Failure      500      {object} ErrorResponse
// @Router       /teams/{team_id} [get]
func (h handler) FindByID(ctx *gin.Context) {
	teamID := ctx.Param("team_id")

	output, err := h.findTeamUseCase.Run(ctx, teamID)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	response := TeamResponse{
		ID:        output.ID,
		Name:      output.Name,
		CreatedAt: output.CreatedAt,
		UpdatedAt: output.UpdatedAt,
	}

	settings.ReturnStatusOK(ctx, response)
}

// FetchTeams godoc
// @Summary      施設IDに紐づくチームを取得する
// @Tags         Team
// @Accept       json
// @Produce      json
// @Param        facility_id query string false "Facility ID"
// @Success      200      {object} []TeamResponse
// @Failure      400      {object} ErrorResponse
// @Failure      500      {object} ErrorResponse
// @Router       /facilities/{facility_id}/teams [get]
func (h handler) FetchByFacilityId(ctx *gin.Context) {
	facilityID := ctx.Param("facility_id")

	output, err := h.fetchTeamsUseCase.Run(ctx, facilityID)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	response := make(TeamListResponse, 0, len(output))
	for _, team := range output {
		response = append(response, TeamResponse{
			ID:        team.ID,
			Name:      team.Name,
			CreatedAt: team.CreatedAt,
			UpdatedAt: team.UpdatedAt,
		})
	}

	settings.ReturnStatusOK(ctx, response)
}
