package department

import (
	"api-buddy/presentation/settings"
	"api-buddy/usecase/facility/department"

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
	id := ctx.Param("id")

	output, err := h.findDepartmentUseCase.Run(ctx, id)
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
		DeletedAt:  output.DeletedAt,
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
	facilityID := ctx.Param("facility_id")

	output, err := h.fetchDepartmentsUseCase.Run(ctx, facilityID)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
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
			DeletedAt:  department.DeletedAt,
		})
	}

	settings.ReturnStatusOK(ctx, response)
}
