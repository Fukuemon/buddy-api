package visit_category

import "context"

type VisitCategoryRepository interface {
	Create(ctx context.Context, visitCategory *VisitCategory) error
	FindAll(ctx context.Context) ([]*VisitCategory, error)
	FindByID(ctx context.Context, id string) (*VisitCategory, error)
}
