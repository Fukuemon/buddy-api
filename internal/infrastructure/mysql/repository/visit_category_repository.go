package repository

import (
	errorDomain "api-buddy/domain/error"
	visitCategoryDomain "api-buddy/domain/visit_info/visit_category"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type VisitCategoryRepository struct {
	db *gorm.DB
}

func NewVisitCategoryRepository() visitCategoryDomain.VisitCategoryRepository {
	return &VisitCategoryRepository{
		db: db.GetDB(),
	}
}

func (r *VisitCategoryRepository) FindAll(ctx context.Context) ([]*visitCategoryDomain.VisitCategory, error) {
	var visitCategories []*visitCategoryDomain.VisitCategory
	err := r.db.Find(&visitCategories).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return visitCategories, nil
}

func (r *VisitCategoryRepository) FindByID(ctx context.Context, id string) (*visitCategoryDomain.VisitCategory, error) {
	var visitCategory visitCategoryDomain.VisitCategory
	err := r.db.Where("id = ?", id).First(&visitCategory).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return &visitCategory, nil
}
