package repository

import (
	errorDomain "api-buddy/domain/error"
	facilityDomain "api-buddy/domain/facility"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type FacilityRepository struct {
	db *gorm.DB
}

func NewFacilityRepository() facilityDomain.FacilityRepository {
	return &FacilityRepository{
		db: db.GetDB(),
	}
}

func (r *FacilityRepository) Create(ctx context.Context, facility *facilityDomain.Facility) error {
	err := r.db.Create(facility).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *FacilityRepository) FindByID(ctx context.Context, id string) (*facilityDomain.Facility, error) {
	var facility facilityDomain.Facility
	err := r.db.Where("id = ?", id).First(&facility).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorDomain.WrapError(errorDomain.NotFoundErr, err)
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return &facility, nil
}
