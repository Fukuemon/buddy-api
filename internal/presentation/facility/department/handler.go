package department

import (
	errorDomain "api-buddy/domain/error"
	"api-buddy/presentation/settings"
	"api-buddy/usecase/facility/department"

	pathValidator "github.com/Fukuemon/go-pkg/validator/gin"
	"github.com/gin-gonic/gin"
)

type handler struct {
	findDepartmentUseCase   *department.FindDepartmentUseCase
	fetchDepartmentsUseCase *department.FetchDepartmentsUseCase
}

func NewHandler(findDepartmentUseCase *department.FindDepartmentUseCase, fetchDepartmentsUseCase *department.FetchDepartmentsUseCase) *handler {
	return &handler{
		findDepartmentUseCase:   findDepartmentUseCase,
		fetchDepartmentsUseCase: fetchDepartmentsUseCase,
	}
}

// FindById godoc
// @Summary      単一の部署を取得する
// @Tags         Department
// @Accept       json
// @Produce      json
// @Param        department_id path string true "Department ID"
// @Success      200      {object} DepartmentResponse
// @Failure      400      {object} ErrorResponse
// @Failure      404      {object} ErrorResponse
// @Failure      500      {object} ErrorResponse
// @Router       /departments/{department_id} [get]
func (h handler) FindById(ctx *gin.Context) {
	departmentId := pathValidator.Param(ctx, "department_id", "required", "ulid")

	err := departmentId.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	output, err := h.findDepartmentUseCase.Run(ctx, departmentId.ParamValue)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	response := DepartmentResponse{
		ID:         output.ID,
		Name:       output.Name,
		FacilityID: output.FacilityID,
		CreateAt:   output.CreatedAt,
		UpdateAt:   output.UpdatedAt,
	}

	settings.ReturnStatusOK(ctx, response)
}

// FetchByFacilityId godoc
// @Summary      施設IDに紐づく部署を取得する
// @Tags         Department
// @Accept       json
// @Produce      json
// @Param        facility_id path string true "Facility ID"
// @Success      200      {object} DepartmentResponse
// @Failure      400      {object} ErrorResponse
// @Failure      500      {object} ErrorResponse
// @Router       /facilities/{facility_id}/departments [get]
func (h handler) FetchByFacilityId(ctx *gin.Context) {
	facilityId := pathValidator.Param(ctx, "facility_id", "required", "ulid")

	err := facilityId.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	output, err := h.fetchDepartmentsUseCase.Run(ctx, facilityId.ParamValue)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := FetchDepartmentsResponse{
		Departments: make([]DepartmentResponse, 0, len(output)),
	}
	for _, department := range output {
		response.Departments = append(response.Departments, DepartmentResponse{
			ID:         department.ID,
			Name:       department.Name,
			FacilityID: department.FacilityID,
			CreateAt:   department.CreatedAt,
			UpdateAt:   department.UpdatedAt,
		})
	}

	settings.ReturnStatusOK(ctx, response)
}
