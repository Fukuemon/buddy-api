package visit_info

import (
	"context"

	"github.com/Fukuemon/go-pkg/query"
)

type VisitInfoRepository interface {
	Create(ctx context.Context, visitInfo *VisitInfo) error
	FindAll(ctx context.Context, filters []query.Filter, sort query.SortOption) ([]*VisitInfo, error)
	FindByID(ctx context.Context, id string) (*VisitInfo, error)
}
