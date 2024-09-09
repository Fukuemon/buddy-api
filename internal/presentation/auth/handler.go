package auth

import (
	errorDomain "api-buddy/domain/error"
	userDomain "api-buddy/domain/user"
	"api-buddy/infrastructure/aws/cognito"
	"api-buddy/presentation/settings"
	"api-buddy/usecase/user"
	"net/http"

	"github.com/Fukuemon/go-pkg/ulid"
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
// @Failure      400      {object}  ErrorResponse
// @Router       /auth/signup [post]
func (h handler) SignUp(ctx *gin.Context) {
	var params SignUpRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// facility, position, department, teamのIDがulidかどうかのチェック
	if !ulid.IsValid(params.FacilityID) {
		settings.ReturnBadRequest(ctx, errorDomain.NewError("施設IDが不正です"))
		return
	}

	if !ulid.IsValid(params.DepartmentID) {
		settings.ReturnBadRequest(ctx, errorDomain.NewError("部署IDが不正です"))
		return
	}

	if !ulid.IsValid(params.PositionID) {
		settings.ReturnBadRequest(ctx, errorDomain.NewError("役職IDが不正です"))
		return
	}

	if !ulid.IsValid(params.TeamID) {
		settings.ReturnBadRequest(ctx, errorDomain.NewError("チームIDが不正です"))
		return
	}

	option := &userDomain.Option{
		Email:       params.Email,
		PhoneNumber: params.PhoneNumber,
	}

	if params.Email == nil && params.PhoneNumber == nil {
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
		settings.ReturnStatusInternalServerError(ctx, err)
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
// @Failure      400      {object}  ErrorResponse
// @Router       /auth/signin [post]
func (h handler) SignIn(ctx *gin.Context) {
	var req SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	authOutput, err := cognito.Actions.SignIn(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	response := SignInResponse{
		AccessToken: *authOutput.AuthenticationResult.AccessToken,
		IdToken:     *authOutput.AuthenticationResult.IdToken,
	}

	ctx.JSON(http.StatusOK, response)
}
