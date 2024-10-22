package user

import (
	errorDomain "api-buddy/domain/error"
	_ "api-buddy/presentation/common"
	"api-buddy/presentation/settings"
	"api-buddy/usecase/user"

	pathValidator "github.com/Fukuemon/go-pkg/validator/gin"
	"github.com/gin-gonic/gin"
)

type handler struct {
	findUserUseCase *user.FindUserUseCase
}

func NewHandler(findUserUseCase *user.FindUserUseCase) *handler {
	return &handler{
		findUserUseCase: findUserUseCase,
	}
}

// FindByUserId godoc
// @Summary      単一のユーザーを取得する
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user_id path string true "User ID"
// @Success      200      {object} UserResponse
// @Failure      400      {object} common.ErrorResponse
// @Failure      403      {object} common.ErrorResponse
// @Failure      404      {object} common.ErrorResponse
// @Failure      500      {object} common.ErrorResponse
// @Router       /users/{user_id} [get]
func (h *handler) FindByUserId(ctx *gin.Context) {
	userId := pathValidator.Param(ctx, "user_id", "required", "ulid")

	err := userId.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}
	output, err := h.findUserUseCase.Run(ctx, userId.ParamValue)
	if err != nil {
		ctx.Error(err)
	}

	policies := []PolicyModel{}
	for _, policy := range output.Policies {
		policies = append(policies, PolicyModel{
			ID:   policy.ID,
			Name: policy.Name,
		})
	}

	response := UserResponse{
		ID:          output.ID,
		Username:    output.Username,
		Position:    output.Position.Name,
		Team:        output.Team.Name,
		Facility:    output.Facility.Name,
		Department:  output.Department.Name,
		Policies:    policies,
		Email:       output.Email,
		PhoneNumber: output.PhoneNumber,
		CreatedAt:   output.CreatedAt,
		UpdatedAt:   output.UpdatedAt,
	}

	settings.ReturnStatusOK(ctx, response)
}
