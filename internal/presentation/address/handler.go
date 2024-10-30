package address

import (
	errorDomain "api-buddy/domain/error"
	_ "api-buddy/presentation/common"
	"api-buddy/presentation/settings"
	"api-buddy/usecase/address"

	pathValidator "github.com/Fukuemon/go-pkg/validator/gin"
	"github.com/gin-gonic/gin"
)

type handler struct {
	findAddressUseCase    *address.FindAddressUseCase
	fetchAddressesUseCase *address.FetchAddressUseCase
	createAddressUseCase  *address.CreateAddressUseCase
}

func NewHandler(createAddressUseCase *address.CreateAddressUseCase, findAddressUseCase *address.FindAddressUseCase, fetchAddressesUseCase *address.FetchAddressUseCase) *handler {
	return &handler{
		createAddressUseCase:  createAddressUseCase,
		findAddressUseCase:    findAddressUseCase,
		fetchAddressesUseCase: fetchAddressesUseCase,
	}
}

// FindById godoc
// @Summary      単一の住所を取得する
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        address_id path string true "Address ID"
// @Success      200      {object} AddressDetailResponse
// @Failure      400      {object} common.ErrorResponse
// @Failure      403      {object} common.ErrorResponse
// @Failure      404      {object} common.ErrorResponse
// @Failure      500      {object} common.ErrorResponse
// @Router       /addresses/{address_id} [get]
func (h *handler) FindById(ctx *gin.Context) {
	addressId := pathValidator.Param(ctx, "address_id", "required", "ulid")

	err := addressId.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	output, err := h.findAddressUseCase.Run(ctx, addressId.ParamValue)
	if err != nil {
		ctx.Error(err)
	}

	response := AddressDetailResponse{
		ID:           output.ID,
		ZipCode:      output.ZipCode,
		Prefecture:   output.Prefecture,
		City:         output.City,
		AddressLine1: output.AddressLine1,
		AddressLine2: output.AddressLine2,
		Latitude:     output.Latitude,
		Longitude:    output.Longitude,
		CreatedAt:    output.CreatedAt,
		UpdatedAt:    output.UpdatedAt,
	}

	settings.ReturnStatusOK(ctx, response)
}

// FetchAddresses godoc
// @Summary      住所一覧を取得する
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        zip_code query string false "Zip Code"
// @Param        prefecture query string false "Prefecture"
// @Param        city query string false "City"
// @Param        address_line1 query string false "Address Line1"
// @Param        address_line2 query string false "Address Line2"
// @Param        latitude query string false "Latitude"
// @Param        longitude query string false "Longitude"
// @Success      200      {object} []AddressDetailResponse
// @Failure      400      {object} common.ErrorResponse
// @Failure      403      {object} common.ErrorResponse
// @Failure      404      {object} common.ErrorResponse
// @Failure      500      {object} common.ErrorResponse
// @Router       /addresses [get]
func (h *handler) Fetch(ctx *gin.Context) {
	zipCode := ctx.Query("zip_code")
	prefecture := ctx.Query("prefecture")
	city := ctx.Query("city")
	addressLine1 := ctx.Query("address_line1")
	addressLine2 := ctx.Query("address_line2")

	input := address.FetchUseCaseInputDto{
		ZipCode:      zipCode,
		Prefecture:   prefecture,
		City:         city,
		AddressLine1: addressLine1,
		AddressLine2: addressLine2,
	}

	output, err := h.fetchAddressesUseCase.Run(ctx, input)
	if err != nil {
		ctx.Error(err)
	}

	response := make(AddressListResponse, 0, len(output))
	for _, address := range output {
		response = append(response, AddressResponse{
			ID:           address.ID,
			ZipCode:      address.ZipCode,
			Prefecture:   address.Prefecture,
			City:         address.City,
			AddressLine1: address.AddressLine1,
			AddressLine2: address.AddressLine2,
			Latitude:     address.Latitude,
			Longitude:    address.Longitude,
		})
	}

	settings.ReturnStatusOK(ctx, response)
}

// CreateAddress godoc
// @Summary      住所を作成する
// @Tags         Address
// @Accept       json
// @Produce      json
// @Success      201      {object} CreateAddressResponse
// @Failure      400      {object} common.ErrorResponse
// @Failure      403      {object} common.ErrorResponse
// @Failure      404      {object} common.ErrorResponse
// @Failure      500      {object} common.ErrorResponse
// @Router       /addresses [post]
func (h *handler) Create(ctx *gin.Context) {
	var params CreateAddressRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	input := address.CreateUseCaseInputDto{
		ZipCode:      params.ZipCode,
		Prefecture:   params.Prefecture,
		City:         params.City,
		AddressLine1: params.AddressLine1,
		AddressLine2: params.AddressLine2,
		Latitude:     params.Latitude,
		Longitude:    params.Longitude,
	}

	address, err := h.createAddressUseCase.Run(ctx, input)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := CreateAddressResponse{
		ID:           address.ID,
		ZipCode:      address.ZipCode,
		Prefecture:   address.Prefecture,
		City:         address.City,
		AddressLine1: address.AddressLine1,
		AddressLine2: address.AddressLine2,
		Latitude:     address.Latitude,
		Longitude:    address.Longitude,
		CreatedAt:    address.CreatedAt,
		UpdatedAt:    address.UpdatedAt,
	}

	settings.ReturnStatusCreated(ctx, response)
}
