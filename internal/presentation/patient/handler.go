package patient

import (
	errorDomain "api-buddy/domain/error"
	patientDomain "api-buddy/domain/patient"
	_ "api-buddy/presentation/common"
	"api-buddy/presentation/settings"
	"api-buddy/usecase/patient"
	"log"

	"github.com/Fukuemon/go-pkg/validator"
	pathValidator "github.com/Fukuemon/go-pkg/validator/gin"
	"github.com/gin-gonic/gin"
)

type handler struct {
	createPatientUseCase *patient.CreatePatientUseCase
	fetchPatientsUseCase *patient.FetchPatientUseCase
}

func NewHandler(createPatientUseCase *patient.CreatePatientUseCase, fetchPatientsUseCase *patient.FetchPatientUseCase) *handler {
	return &handler{
		createPatientUseCase: createPatientUseCase,
		fetchPatientsUseCase: fetchPatientsUseCase,
	}
}

// CreatePatient godoc
// @Summary      患者を作成する
// @Tags         Patient
// @Accept       json
// @Produce      json
// @Param        patient body CreatePatientRequest true "Patient"
// @Success      201      {object} CreatePatientResponse
// @Failure      400      {object} common.ErrorResponse
// @Failure      403      {object} common.ErrorResponse
// @Failure      404      {object} common.ErrorResponse
// @Failure      500      {object} common.ErrorResponse
// @Router       /facilities/{facility_id}/patients [post]
func (h *handler) CreatePatient(ctx *gin.Context) {
	facilityId := pathValidator.Param(ctx, "facility_id", "required", "ulid")
	var params CreatePatientRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	if err := validator.StructValidation(params); err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	input := patient.CreatePatientUseCaseInputDto{
		Name:            params.Name,
		PreferredTime:   params.PreferredTime,
		PreferredGender: params.PreferredGender,
		ServiceCodeID:   params.ServiceCodeID,
		AddressID:       params.AddressID,
		AreaID:          params.AreaID,
		AssignedStaffID: params.AssignedStaffID,
		FacilityID:      facilityId.ParamValue,
	}

	// ユースケースを実行して患者を作成
	output, err := h.createPatientUseCase.Run(ctx, input)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := CreatePatientResponse{
		ID:              output.ID,
		Name:            output.Name,
		PreferredTime:   output.PreferredTime,
		PreferredGender: output.PreferredGender,
		AssignedStaff:   output.AssignedStaff,
		Address:         output.Address,
		Area:            output.Area,
		Facility:        output.Facility,
	}

	settings.ReturnStatusCreated(ctx, response)

}

// FetchPatients swagger
// @Summary      施設IDに紐づく患者を取得する
// @Tags         Patient
// @Accept       json
// @Produce      json
// @Param        facility_id path string true "Facility ID"
// @Param        name query string false "Name"
// @Param        preferred_time query string false "Preferred Time"
// @Param 	     preferred_gender query string false "Preferred Gender"
// @Param        assigned_staff query string false "Assigned Staff"
// @Param        zip_code query string false "Zip Code"
// @Param        sort_field query string false "Sort Field"
// @Param        sort_order query string false "Sort Order (asc or desc)"
// @Success      200      {array} PatientListResponse
// @Failure      400      {object} common.ErrorResponse
// @Failure      500      {object} common.ErrorResponse
// @Router       /facilities/{facility_id}/patients [get]
func (h *handler) FetchByFacilityId(ctx *gin.Context) {
	facilityId := pathValidator.Param(ctx, "facility_id", "required", "ulid")
	err := facilityId.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	// クエリパラメータを取得
	name := ctx.Query("name")
	preferredTimeStr := ctx.Query("preferred_time")
	var preferredTime *patientDomain.PreferredTime
	preferredGenderStr := ctx.Query("preferred_gender")
	var preferredGender *patientDomain.PreferredGender
	assignedStaff := ctx.Query("assigned_staff")
	zipCode := ctx.Query("zip_code")
	area := ctx.Query("area")
	sortField := ctx.Query("sort_field")
	sortOrder := ctx.Query("sort_order")

	if preferredTimeStr != "" {
		tempPreferredTime := patientDomain.PreferredTime(preferredTimeStr) // 一時変数を作成
		if !tempPreferredTime.IsValid() {
			ctx.Error(errorDomain.WrapError(errorDomain.InvalidInputErr, errorDomain.NewError("希望する時刻の値が不正です")))
			return
		}
		preferredTime = &tempPreferredTime // ポインタを設定
	}

	if preferredGenderStr != "" {
		tempPreferredGender := patientDomain.PreferredGender(preferredGenderStr) // 一時変数を作成
		if !tempPreferredGender.IsValid() {
			ctx.Error(errorDomain.WrapError(errorDomain.InvalidInputErr, errorDomain.NewError("希望する性別の値が不正です")))
			return
		}
		preferredGender = &tempPreferredGender // ポインタを設定
	}

	log.Printf("preferredTime: %v, preferredGender: %v", preferredTime, preferredGender)

	// フィルタリングとソートのための DTO を作成
	input := patient.FetchPatientUseCaseInputDto{
		Name:            name,
		PreferredTime:   preferredTime,
		PreferredGender: preferredGender,
		AssignedStaff:   assignedStaff,
		ZipCode:         zipCode,
		Area:            area,
		SortField:       sortField,
		SortOrder:       sortOrder,
	}

	// ユースケースを実行して患者一覧を取得
	output, err := h.fetchPatientsUseCase.Run(ctx, facilityId.ParamValue, input)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := make(PatientListResponse, 0, len(output))
	for _, patient := range output {
		response = append(response, PatientResponse{
			ID:              patient.ID,
			Name:            patient.Name,
			PreferredTime:   patient.PreferredTime,
			PreferredGender: patient.PreferredGender,
			AssignedStaff:   patient.AssignedStaff,
			Address:         patient.Address.JoinAddress(),
			Area:            patient.Area,
		})
	}

	settings.ReturnStatusOK(ctx, response)
}
