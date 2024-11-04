package repository

import (
	errorDomain "api-buddy/domain/error"
	routeDomain "api-buddy/domain/visit_info/route"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type RouteRepository struct {
	db *gorm.DB
}

func NewRouteRepository() routeDomain.RouteRepository {
	return &RouteRepository{
		db: db.GetDB(),
	}
}

func (r *RouteRepository) Create(ctx context.Context, route *routeDomain.Route) error {
	err := r.db.Create(route).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *RouteRepository) FindAll(ctx context.Context) ([]*routeDomain.Route, error) {
	var routes []*routeDomain.Route
	err := r.db.Find(&routes).Error
	if err != nil {
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return routes, nil
}

func (r *RouteRepository) FindByID(ctx context.Context, id string) (*routeDomain.Route, error) {
	var route routeDomain.Route
	err := r.db.Where("id = ?", id).First(&route).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorDomain.WrapError(errorDomain.NotFoundErr, err)
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return &route, nil
}
