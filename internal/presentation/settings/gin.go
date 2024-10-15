package settings

import (
	"errors"
	"net/http"

	errorDomain "api-buddy/domain/error"

	"github.com/gin-gonic/gin"
)

func NewGinEngine() *gin.Engine {
	router := gin.Default()

	return router
}

func ReturnStatusOK[T any](ctx *gin.Context, body T) {
	ctx.JSON(http.StatusOK, body)
}

func ReturnStatusCreated[T any](ctx *gin.Context, body T) {
	ctx.JSON(http.StatusCreated, body)
}

func ReturnStatusNoContent(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func ReturnStatusBadRequest(ctx *gin.Context, err error) {
	returnAbortWith(ctx, http.StatusBadRequest, err)
}

func ReturnBadRequest(ctx *gin.Context, err error) {
	ReturnStatusBadRequest(ctx, err)
}

func ReturnStatusUnauthorized(ctx *gin.Context, err error) {
	returnAbortWith(ctx, http.StatusUnauthorized, err)
}

func ReturnUnauthorized(ctx *gin.Context, err error) {
	ReturnStatusUnauthorized(ctx, err)
}

func ReturnStatusForbidden(ctx *gin.Context, err error) {
	returnAbortWith(ctx, http.StatusForbidden, err)
}

func ReturnForbidden(ctx *gin.Context, err error) {
	ReturnStatusForbidden(ctx, err)
}

func ReturnStatusNotFound(ctx *gin.Context, err error) {
	returnAbortWith(ctx, http.StatusNotFound, err)
}

func ReturnNotFound(ctx *gin.Context, err error) {
	ReturnStatusNotFound(ctx, err)
}

func ReturnStatusInternalServerError(ctx *gin.Context, err error) {
	returnAbortWith(ctx, http.StatusInternalServerError, err)
}

func ReturnError(ctx *gin.Context, err error) {
	ctx.Error(err)
}
func returnAbortWith(ctx *gin.Context, code int, err error) {
	var domainErr *errorDomain.Error
	if errors.As(err, &domainErr) { // trueの場合、domainErrにerrの内容が格納される
		msg := domainErr.Error()

		ctx.AbortWithStatusJSON(code, gin.H{
			"code":        code,
			"description": domainErr.Description(), // 事前定義したエラー情報を持つ(domain/error参照)
			"msg":         msg,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"code": code,
		"msg":  err.Error(),
	})
}

// エラー内容によって適切なステータスコードを返す
func HandleErrorResponse(ctx *gin.Context, err error) {
	var domainErr *errorDomain.Error

	if errors.As(err, &domainErr) { // trueの場合、domainErrにerrの内容が格納される
		switch domainErr.Description() {
		case errorDomain.InvalidInputErr.Description():
			ReturnBadRequest(ctx, err)
		case errorDomain.NotFoundErr.Description():
			ReturnNotFound(ctx, err)
		case errorDomain.GeneralDBError.Description():
			ReturnStatusInternalServerError(ctx, err)
		case errorDomain.UnAuthorizedErr.Description():
			ReturnUnauthorized(ctx, err)
		case errorDomain.ForbiddenErr.Description():
			ReturnForbidden(ctx, err)
		default:
			ReturnStatusInternalServerError(ctx, err)
		}
	} else {
		ReturnStatusInternalServerError(ctx, err)
	}
}
