package repository

import (
	"api-buddy/domain/common"
	errorDomain "api-buddy/domain/error"
	visitInfoDomain "api-buddy/domain/visit_info"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"github.com/Fukuemon/go-pkg/query"
	"gorm.io/gorm"
)

type VisitInfoRepository struct {
	db *gorm.DB
}

func NewVisitInfoRepository() visitInfoDomain.VisitInfoRepository {
	return &VisitInfoRepository{
		db: db.GetDB(),
	}
}

func (r *VisitInfoRepository) Create(ctx context.Context, visitInfo *visitInfoDomain.VisitInfo) error {
	err := r.db.Create(visitInfo).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *VisitInfoRepository) FindAll(ctx context.Context, filters []query.Filter, sort query.SortOption) ([]*visitInfoDomain.VisitInfo, error) {
	q := query.NewQuery()

	for _, filter := range filters {
		filter.Apply(q)
	}

	dbQuery := r.db

	for _, mapping := range visitInfoDomain.VisitInfoRelationMappings {
		if value, exists := q.Filters[mapping.FilterField]; exists {
			dbQuery = dbQuery.Joins("JOIN "+mapping.TableName+" ON "+mapping.JoinKey).
				Where(mapping.FilterField+" = ?", value)
		}
	}

	for key, value := range q.Filters {
		if _, isRelationField := visitInfoDomain.VisitInfoRelationMappings[key]; !isRelationField {
			dbQuery = dbQuery.Where(key, value)
		}
	}

	if sort.Field != "" {
		if mapping, exists := visitInfoDomain.VisitInfoRelationMappings[sort.Field]; exists {
			dbQuery = dbQuery.Joins("JOIN " + mapping.TableName + " ON " + mapping.JoinKey).
				Order(mapping.FilterField + " " + sort.Order)
		} else {
			dbQuery = dbQuery.Order(sort.Field + " " + sort.Order)
		}
	} else {
		dbQuery = dbQuery.Order(common.UpdatedAt + " " + query.DESC)
	}

	var visitInfos []*visitInfoDomain.VisitInfo
	err := dbQuery.Preload("Patient").Preload("User").Preload("Route").Find(&visitInfos).Error
	if err != nil {
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return visitInfos, nil
}

func (r *VisitInfoRepository) FindByID(ctx context.Context, id string) (*visitInfoDomain.VisitInfo, error) {
	var visitInfo visitInfoDomain.VisitInfo
	err := r.db.Preload("Patient").Preload("User").Preload("Route").Where("id = ?", id).First(&visitInfo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorDomain.WrapError(errorDomain.NotFoundErr, err)
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return &visitInfo, nil
}
