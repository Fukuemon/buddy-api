package route

import (
	"context"

	"gorm.io/gorm"
)

type RouteRepository interface {
	FindByID(ctx context.Context, id string) (*Route, error)
	Create(ctx context.Context, tx *gorm.DB, route *Route) error
}
