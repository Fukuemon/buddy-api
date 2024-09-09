package repository

import (
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
	if err := r.db.Create(facility).Error; err != nil {
		return err
	}
	return nil
}

func (r *FacilityRepository) FindByID(ctx context.Context, id string) (*facilityDomain.Facility, error) {
	var facility facilityDomain.Facility
	err := r.db.Where("id = ?", id).First(&facility).Error
	if err != nil {
		return nil, err
	}
	return &facility, nil
}
