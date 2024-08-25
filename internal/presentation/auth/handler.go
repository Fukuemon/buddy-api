package auth

import (
	"api-buddy/infrastructure/aws/cognito"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignIn godoc
// @Summary      サインイン用エンドポイント
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      SignInRequest  true  "Sign In Request"
// @Success      200      {object}  SignInResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /auth/signin [post]
func SignIn(ctx *gin.Context) {
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
