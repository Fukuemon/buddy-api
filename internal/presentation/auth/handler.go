package auth

import (
	errorDomain "api-buddy/domain/error"
	userDomain "api-buddy/domain/user"
	"api-buddy/infrastructure/aws/cognito"
	_ "api-buddy/presentation/common"
	"api-buddy/presentation/settings"
	"api-buddy/usecase/user"

	"github.com/Fukuemon/go-pkg/validator"
	"github.com/gin-gonic/gin"
)

type handler struct {
	createUserUseCase *user.CreateUserUseCase
}

func NewHandler(createUserUseCase *user.CreateUserUseCase) *handler {
	return &handler{
		createUserUseCase: createUserUseCase,
	}
}

// SignUp godoc
// @Summary      サインアップ用エンドポイント
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      SignUpRequest  true  "Sign Up Request"
// @Success      201      {object}  SignUpResponse
// @Failure      400      {object}  common.ErrorResponse
// @Failure      500      {object}  common.ErrorResponse
// @Router       /auth/signup [post]
func (h handler) SignUp(ctx *gin.Context) {
	var params SignUpRequest

	if err := validator.StructValidation(params); err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	option := &userDomain.Option{
		Email:       params.Email,
		PhoneNumber: params.PhoneNumber,
	}

	if (params.Email == nil || *params.Email == "") && (params.PhoneNumber == nil || *params.PhoneNumber == "") {
		settings.ReturnBadRequest(ctx, errorDomain.NewError("メールアドレスか電話番号のどちらかは必須です"))
		return
	}

	input := user.CreateUserUseCaseInputDto{
		Username:     params.Username,
		Password:     params.Password,
		FacilityID:   params.FacilityID,
		DepartmentID: params.DepartmentID,
		PositionID:   params.PositionID,
		TeamID:       params.TeamID,
		Option:       option,
	}

	output, err := h.createUserUseCase.Run(ctx, input)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := SignUpResponse{
		ID:             output.ID,
		UserName:       output.Username,
		Email:          output.Email,
		PhoneNumber:    output.PhoneNumber,
		FacilityName:   output.Facility.Name,
		DepartmentName: output.Department.Name,
		PositionName:   output.Position.Name,
		TeamName:       output.Team.Name,
		Policies:       output.Policies,
	}

	settings.ReturnStatusCreated(ctx, response)

}

// SignIn godoc
// @Summary      サインイン用エンドポイント
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      SignInRequest  true  "Sign In Request"
// @Success      201      {object}  SignInResponse
// @Failure      400      {object}  common.ErrorResponse
// @Failure      500      {object}  common.ErrorResponse
// @Router       /auth/signin [post]
func (h handler) SignIn(ctx *gin.Context) {
	var req SignInRequest

	if err := validator.StructValidation(req); err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	authOutput, err := cognito.Actions.SignIn(req.Username, req.Password)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := SignInResponse{
		AccessToken: *authOutput.AuthenticationResult.AccessToken,
		IdToken:     *authOutput.AuthenticationResult.IdToken,
	}

	settings.ReturnStatusCreated(ctx, response)
}
