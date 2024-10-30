package repository

import (
	errorDomain "api-buddy/domain/error"
	areaDomain "api-buddy/domain/facility/area"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type AreaRepository struct {
	db *gorm.DB
}

func NewAreaRepository() areaDomain.AreaRepository {
	return &AreaRepository{
		db: db.GetDB(),
	}
}

func (r *AreaRepository) Create(ctx context.Context, area *areaDomain.Area) error {
	err := r.db.Create(area).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *AreaRepository) FindByID(ctx context.Context, id string) (*areaDomain.Area, error) {
	var area areaDomain.Area
	err := r.db.Preload("Addresses").Begin().Where("id = ?", id).First(&area).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorDomain.WrapError(errorDomain.NotFoundErr, err)
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return &area, nil
}

func (r *AreaRepository) FindByFacilityID(ctx context.Context, facilityID string) ([]*areaDomain.Area, error) {
	var areas []*areaDomain.Area
	err := r.db.Preload("Addresses").Begin().Where("facility_id = ?", facilityID).Find(&areas).Error
	if err != nil {
		return nil, err
	}
	return areas, nil
}
