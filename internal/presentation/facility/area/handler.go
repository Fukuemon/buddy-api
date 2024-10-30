package area

import (
	errorDomain "api-buddy/domain/error"
	_ "api-buddy/presentation/common"
	"api-buddy/presentation/settings"
	"api-buddy/usecase/facility/area"

	"github.com/Fukuemon/go-pkg/validator"
	pathValidator "github.com/Fukuemon/go-pkg/validator/gin"
	"github.com/gin-gonic/gin"
)

type handler struct {
	createAreaUseCase *area.CreateAreaUseCase
	findAreaUseCase   *area.FindAreaUseCase
	fetchAreasUseCase *area.FetchAreaUseCase
}

func NewHandler(createAreaUseCase *area.CreateAreaUseCase, findAreaUseCase *area.FindAreaUseCase, fetchAreasUseCase *area.FetchAreaUseCase) *handler {
	return &handler{
		createAreaUseCase: createAreaUseCase,
		findAreaUseCase:   findAreaUseCase,
		fetchAreasUseCase: fetchAreasUseCase,
	}
}

// Create godoc
// @Summary      施設に紐づくエリアを作成する
// @Tags         Area
// @Accept       json
// @Produce      json
// @Param        request body      CreateAreaRequest  true  "Create Area Request"
// @Success      201      {object} AreaResponse
// @Failure      400      {object} common.ErrorResponse
// @Failure      403      {object} common.ErrorResponse
// @Failure      500      {object} common.ErrorResponse
// @Router       /facilities/{facility_id}/areas [post]
func (h handler) Create(ctx *gin.Context) {
	facilityId := pathValidator.Param(ctx, "facility_id", "required", "ulid")
	var params CreateAreaRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	err := facilityId.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	if err := validator.StructValidation(params); err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	input := area.CreateAreaUseCaseInputDto{
		Name:       params.Name,
		FacilityID: facilityId.ParamValue,
		AddressIDs: params.AddressIDs,
	}

	output, err := h.createAreaUseCase.Run(ctx, input)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	response := CreateAreaResponse{
		ID:         output.ID,
		Name:       output.Name,
		FacilityID: output.FacilityID,
	}

	settings.ReturnStatusCreated(ctx, response)
}

// FetchByFacilityId godoc
// @Summary      施設IDに紐づくエリアを取得する
// @Tags         Area
// @Accept       json
// @Produce      json
// @Param        facility_id path string true "Facility ID"
// @Success      200      {object} AreaResponse
// @Failure      400      {object} common.ErrorResponse
// @Failure      403      {object} common.ErrorResponse
// @Failure      500      {object} common.ErrorResponse
// @Router       /facilities/{facility_id}/areas [get]
func (h handler) FetchByFacilityId(ctx *gin.Context) {
	facilityId := pathValidator.Param(ctx, "facility_id", "required", "ulid")

	err := facilityId.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	output, err := h.fetchAreasUseCase.Run(ctx, facilityId.ParamValue)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := make(AreaListResponse, 0, len(output))
	for _, area := range output {
		addresses := make([]AddressModel, 0, len(area.Addresses))
		for _, address := range area.Addresses {
			addresses = append(addresses, AddressModel{
				ID:           address.ID,
				Prefecture:   address.Prefecture,
				City:         address.City,
				AddressLine1: address.AddressLine1,
				AddressLine2: address.AddressLine2,
				Latitude:     address.Latitude,
				Longitude:    address.Longitude,
			})
		}
		response = append(response, AreaResponse{
			ID:         area.ID,
			Name:       area.Name,
			FacilityID: area.FacilityID,
			Addresses:  addresses,
		})
	}

	settings.ReturnStatusOK(ctx, response)
}

// FindById godoc
// @Summary      単一のエリアを取得する
// @Tags         Area
// @Accept       json
// @Produce      json
// @Param        area_id path string true "Area ID"
// @Success      200      {object} AreaResponse
// @Failure      400      {object} common.ErrorResponse
// @Failure      403      {object} common.ErrorResponse
// @Failure      404      {object} common.ErrorResponse
// @Failure      500      {object} common.ErrorResponse
// @Router       /areas/{area_id} [get]
func (h handler) FindById(ctx *gin.Context) {
	areaId := pathValidator.Param(ctx, "area_id", "required", "ulid")

	err := areaId.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	output, err := h.findAreaUseCase.Run(ctx, areaId.ParamValue)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	addresses := make([]AddressModel, 0, len(output.Addresses))
	for _, address := range output.Addresses {
		addresses = append(addresses, AddressModel{
			ID:           address.ID,
			Prefecture:   address.Prefecture,
			City:         address.City,
			AddressLine1: address.AddressLine1,
			AddressLine2: address.AddressLine2,
			Latitude:     address.Latitude,
			Longitude:    address.Longitude,
		})
	}

	response := AreaResponse{
		ID:         output.ID,
		Name:       output.Name,
		FacilityID: output.FacilityID,
		Addresses:  addresses,
	}

	settings.ReturnStatusOK(ctx, response)
}
