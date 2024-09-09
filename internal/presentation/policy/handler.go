package policy

import (
	"api-buddy/presentation/settings"
	"api-buddy/usecase/policy"

	"github.com/gin-gonic/gin"
)

type handler struct {
	createPolicyUseCase  *policy.CreatePolicyUseCase
	findPolicyUseCase    *policy.FindPolicyUseCase
	fetchPoliciesUseCase *policy.FetchPoliciesUseCase
}

func NewHandler(createPolicyUseCase *policy.CreatePolicyUseCase, findPolicyUseCase *policy.FindPolicyUseCase, fetchPoliciesUseCase *policy.FetchPoliciesUseCase) *handler {
	return &handler{
		createPolicyUseCase:  createPolicyUseCase,
		findPolicyUseCase:    findPolicyUseCase,
		fetchPoliciesUseCase: fetchPoliciesUseCase,
	}
}

// Create godoc
// @Summary      ポリシーを作成する
// @Tags         Policy
// @Accept       json
// @Produce      json
// @Param        request body      CreatePolicyRequest  true  "Create Policy Request"
// @Success      201      {object} CreatePolicyResponse
// @Failure      400      {object} ErrorResponse
// @Failure      500      {object} ErrorResponse
// @Router       /policies [post]
func (h handler) Create(ctx *gin.Context) {
	var params CreatePolicyRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		settings.ReturnBadRequest(ctx, err)
		return
	}

	input := policy.CreateUseCaseInputDto{
		Name: params.Name,
	}

	output, err := h.createPolicyUseCase.Run(ctx, &input)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	response := CreatePolicyResponse{
		ID:   output.ID,
		Name: output.Name,
	}

	settings.ReturnStatusCreated(ctx, response)
}

// FindById godoc
// @Summary      単一のポリシーを取得する
// @Tags         Policy
// @Accept       json
// @Produce      json
// @Param        policy_id path string true "Policy ID"
// @Success      200      {object} PolicyResponse
// @Failure      400      {object} ErrorResponse
// @Failure      404      {object} ErrorResponse
// @Failure      500      {object} ErrorResponse
// @Router       /policies/{policy_id} [get]
func (h handler) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	output, err := h.findPolicyUseCase.Run(ctx, id)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	response := PolicyResponse{
		ID:        output.ID,
		Name:      output.Name,
		CreatedAt: output.CreatedAt,
		UpdatedAt: output.UpdatedAt,
		DeletedAt: output.DeletedAt,
	}

	settings.ReturnStatusOK(ctx, response)
}

// Fetch godoc
// @Summary      ポリシー一覧を取得する
// @Tags         Policy
// @Accept       json
// @Produce      json
// @Success      200      {object} PolicyListResponse
// @Failure      400      {object} ErrorResponse
// @Failure      500      {object} ErrorResponse
// @Router       /policies [get]
func (h handler) Fetch(ctx *gin.Context) {
	output, err := h.fetchPoliciesUseCase.Run(ctx)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	response := PolicyListResponse{}
	for _, o := range output {
		response = append(response, PolicyResponse{
			ID:        o.ID,
			Name:      o.Name,
			CreatedAt: o.CreatedAt,
			UpdatedAt: o.UpdatedAt,
			DeletedAt: o.DeletedAt,
		})
	}

	settings.ReturnStatusOK(ctx, response)
}
