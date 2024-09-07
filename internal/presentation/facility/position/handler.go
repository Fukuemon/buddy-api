package position

import (
	"api-buddy/presentation/settings"
	"api-buddy/usecase/facility/position"

	"github.com/gin-gonic/gin"
)

type handler struct {
	createPositionUseCase *position.CreatePositionUseCase
	findPositionUseCase   *position.FindPositionUseCase
	fetchPositionsUseCase *position.FetchPositionsUseCase
}

func NewHandler(createPositionUseCase *position.CreatePositionUseCase, findPositionUseCase *position.FindPositionUseCase, fetchPositionsUseCase *position.FetchPositionsUseCase) *handler {
	return &handler{
		createPositionUseCase: createPositionUseCase,
		findPositionUseCase:   findPositionUseCase,
		fetchPositionsUseCase: fetchPositionsUseCase,
	}
}

// Create godoc
// @Summary      施設に紐づく役職を作成する
// @Tags         Position
// @Accept       json
// @Produce      json
// @Param        request body      CreatePositionRequest  true  "Create Position Request"
// @Success      201      {object} PositionResponse
// @Failure      400      {object} ErrorResponse
// @Failure      500      {object} ErrorResponse
// @Router       /facilities/{facility_id}/positions [post]
func (h handler) CreateByFacilityId(ctx *gin.Context) {
	facilityID := ctx.Param("facility_id")
	var params CreatePositionRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		settings.ReturnBadRequest(ctx, err)
		return
	}

	input := position.CreateUseCaseInputDto{
		Name:       params.Name,
		FacilityID: facilityID,
		PolicyIDs:  params.PolicyIDs,
	}

	output, err := h.createPositionUseCase.Create(ctx, input)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	response := CreatePositionResponse{
		ID:       output.ID,
		Name:     output.Name,
		Policies: output.Policies,
	}

	settings.ReturnStatusCreated(ctx, response)
}

// FindById godoc
// @Summary      単一の役職を取得する
// @Tags         Position
// @Accept       json
// @Produce      json
// @Param        position_id path string true "Position ID"
// @Success      200      {object} PositionResponse
// @Failure      400      {object} ErrorResponse
// @Failure      500      {object} ErrorResponse
// @Router       /positions/{position_id} [get]
func (h handler) FindById(ctx *gin.Context) {
	positionID := ctx.Param("id")

	output, err := h.findPositionUseCase.Run(ctx, positionID)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	response := PositionResponse{
		ID:         output.ID,
		Name:       output.Name,
		FacilityID: output.FacilityID,
		Policies:   output.Policies,
		CreatedAt:  output.CreatedAt,
		UpdatedAt:  output.UpdatedAt,
		DeletedAt:  output.DeletedAt,
	}

	settings.ReturnStatusOK(ctx, response)
}

// FetchByFacilityId godoc
// @Summary      施設IDに紐づく役職を取得する
// @Tags         Position
// @Accept       json
// @Produce      json
// @Param        facility_id path string true "Facility ID"
// @Success      200      {object} PositionResponse
// @Failure      400      {object} ErrorResponse
// @Failure      500      {object} ErrorResponse
// @Router       /facilities/{facility_id}/positions [get]
func (h handler) FetchByFacilityId(ctx *gin.Context) {
	facilityID := ctx.Param("facility_id")

	output, err := h.fetchPositionsUseCase.Run(ctx, facilityID)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	response := make(PositionListResponse, 0, len(output))

	for _, position := range output {
		response = append(response, PositionResponse{
			ID:         position.ID,
			Name:       position.Name,
			FacilityID: position.FacilityID,
			Policies:   position.Policies,
			CreatedAt:  position.CreatedAt,
			UpdatedAt:  position.UpdatedAt,
			DeletedAt:  position.DeletedAt,
		})
	}

	settings.ReturnStatusOK(ctx, response)
}
