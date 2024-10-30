package policy

import (
	errorDomain "api-buddy/domain/error"
	_ "api-buddy/presentation/common"
	"api-buddy/presentation/settings"
	"api-buddy/usecase/policy"

	"github.com/Fukuemon/go-pkg/validator"
	pathValidator "github.com/Fukuemon/go-pkg/validator/gin"
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
// @Failure      400      {object} common.ErrorResponse
// @Failure      500      {object} common.ErrorResponse
// @Router       /policies [post]
func (h handler) Create(ctx *gin.Context) {
	var params CreatePolicyRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	if err := validator.StructValidation(params); err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	input := policy.CreateUseCaseInputDto{
		Name: params.Name,
	}

	output, err := h.createPolicyUseCase.Run(ctx, &input)
	if err != nil {
		ctx.Error(err)
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
// @Failure      400      {object} common.ErrorResponse
// @Failure      403      {object} common.ErrorResponse
// @Failure      404      {object} common.ErrorResponse
// @Failure      500      {object} common.ErrorResponse
// @Router       /policies/{policy_id} [get]
func (h handler) FindById(ctx *gin.Context) {
	policyId := pathValidator.Param(ctx, "policy_id", "required", "ulid")

	err := policyId.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	output, err := h.findPolicyUseCase.Run(ctx, policyId.ParamValue)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := PolicyResponse{
		ID:        output.ID,
		Name:      output.Name,
		CreatedAt: output.CreatedAt,
		UpdatedAt: output.UpdatedAt,
	}

	settings.ReturnStatusOK(ctx, response)
}

// Fetch godoc
// @Summary      ポリシー一覧を取得する
// @Tags         Policy
// @Accept       json
// @Produce      json
// @Success      200      {object} PolicyListResponse
// @Failure      400      {object} common.ErrorResponse
// @Failure      403      {object} common.ErrorResponse
// @Failure      500      {object} common.ErrorResponse
// @Router       /policies [get]
func (h handler) Fetch(ctx *gin.Context) {
	output, err := h.fetchPoliciesUseCase.Run(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := make([]PolicyResponse, 0, len(output))
	for _, o := range output {
		response = append(response, PolicyResponse{
			ID:        o.ID,
			Name:      o.Name,
			CreatedAt: o.CreatedAt,
			UpdatedAt: o.UpdatedAt,
		})
	}

	settings.ReturnStatusOK(ctx, response)
}
