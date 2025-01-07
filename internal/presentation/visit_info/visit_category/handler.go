package visit_category

import (
	_ "api-buddy/presentation/common"
	"api-buddy/presentation/settings"
	visitCategoryUse "api-buddy/usecase/visit_info/visit_category"

	"github.com/gin-gonic/gin"
)

type handler struct {
	fetchVisitCategoriesUseCase *visitCategoryUse.FetchVisitCategoriesUseCase
}

func NewHandler(fetchVisitCategoriesUseCase *visitCategoryUse.FetchVisitCategoriesUseCase) *handler {
	return &handler{
		fetchVisitCategoriesUseCase: fetchVisitCategoriesUseCase,
	}
}

// FetchVisitCategories godoc
// @Summary 訪問カテゴリ一覧を取得する
// @Tags VisitCategory
// @Accept json
// @Produce json
// @Success 200 {object} VisitCategoryListResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /visit_infos/visit_categories [get]
func (h *handler) FetchVisitCategories(ctx *gin.Context) {
	output, err := h.fetchVisitCategoriesUseCase.Run(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := make(VisitCategoryListResponse, 0, len(output))
	for _, o := range output {
		response = append(response, VisitCategoryResponse{
			ID:   o.ID,
			Name: o.Name,
		})
	}
	settings.ReturnStatusOK(ctx, response)
}
