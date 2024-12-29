package visit_category

import "api-buddy/domain/visit_info/visit_category"

type VisitCategoryResponse struct {
	ID   string                           `json:"id"`
	Name visit_category.VisitCategoryType `json:"name"`
}

type VisitCategoryListResponse []VisitCategoryResponse
