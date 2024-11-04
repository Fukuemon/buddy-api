package route

import "context"

type RouteRepository interface {
	FindByID(ctx context.Context, id string) (*Route, error)
	Create(ctx context.Context, route *Route) error
}
