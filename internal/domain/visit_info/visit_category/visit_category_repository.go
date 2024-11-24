package visit_category

import "context"

type VisitCategoryRepository interface {
	FindAll(ctx context.Context) ([]*VisitCategory, error)
	FindByID(ctx context.Context, id string) (*VisitCategory, error)
}
