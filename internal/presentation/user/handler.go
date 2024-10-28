package user

import (
	errorDomain "api-buddy/domain/error"
	_ "api-buddy/presentation/common"
	"api-buddy/presentation/settings"
	"api-buddy/usecase/user"
	"strings"

	pathValidator "github.com/Fukuemon/go-pkg/validator/gin"
	"github.com/gin-gonic/gin"
)

type handler struct {
	findUserUseCase   *user.FindUserUseCase
	fetchUsersUseCase *user.FetchUsersUseCase
}

func NewHandler(findUserUseCase *user.FindUserUseCase, fetchUsersUseCase *user.FetchUsersUseCase) *handler {
	return &handler{
		findUserUseCase:   findUserUseCase,
		fetchUsersUseCase: fetchUsersUseCase,
	}
}

// FindByUserId godoc
// @Summary      単一のユーザーを取得する
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user_id path string true "User ID"
// @Success      200      {object} UserDetailResponse
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

	response := UserDetailResponse{
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

// FindUsers godoc
// @Summary      施設IDに紐づくユーザーを取得する
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        facility_id path string true "Facility ID"
// @Param        username query string false "Username"
// @Param        position query string false "Position"
// @Param        department query string false "Department"
// @Param        team query string false "Team"
// @Param        sort_field query string false "Sort Field"
// @Param        sort_order query string false "Sort Order (asc or desc)"
// @Success      200      {array} UserResponse
// @Failure      400      {object} common.ErrorResponse
// @Failure      500      {object} common.ErrorResponse
// @Router       /facilities/{facility_id}/users [get]
func (h *handler) FetchByFacilityId(ctx *gin.Context) {
	facilityId := pathValidator.Param(ctx, "facility_id", "required", "ulid")
	err := facilityId.ParamValidate()
	if err != nil {
		ctx.Error(errorDomain.ValidationError(err))
		return
	}

	// クエリパラメータを取得
	username := ctx.Query("username")
	position := ctx.Query("position")
	department := ctx.Query("department")
	team := ctx.Query("team")
	sortField := ctx.Query("sort_field")
	sortOrder := strings.ToLower(ctx.Query("sort_order"))

	// フィルタリングとソートのための DTO を作成
	input := user.FetchUsersUseCaseInputDto{
		Username:   username,
		Position:   position,
		Department: department,
		Team:       team,
		SortField:  sortField,
		SortOrder:  sortOrder,
	}

	// ユースケースを実行してユーザー一覧を取得
	output, err := h.fetchUsersUseCase.Run(ctx, facilityId.ParamValue, input)
	if err != nil {
		ctx.Error(err)
		return
	}

	// 出力を整形してレスポンスを作成
	response := UserListResponse{}
	for _, user := range output {
		response.Users = append(response.Users, UserResponse{
			ID:         user.ID,
			Username:   user.Username,
			Position:   user.Position,
			Team:       user.Team,
			Department: user.Department,
		})
	}

	// レスポンスを返す
	settings.ReturnStatusOK(ctx, response)
}
