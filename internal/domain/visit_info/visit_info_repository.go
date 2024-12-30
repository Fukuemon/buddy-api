package visit_info

import (
	"context"

	"github.com/Fukuemon/go-pkg/query"
	"gorm.io/gorm"
)

type VisitInfoRepository interface {
	Create(ctx context.Context, tx *gorm.DB, visitInfo *VisitInfo) error
	FindAll(ctx context.Context, filters []query.Filter, sort query.SortOption) ([]*VisitInfo, error)
	FindByID(ctx context.Context, id string) (*VisitInfo, error)
}
