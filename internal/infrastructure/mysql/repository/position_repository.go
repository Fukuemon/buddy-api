package repository

import (
	errorDomain "api-buddy/domain/error"
	positionDomain "api-buddy/domain/facility/position"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type PositionRepository struct {
	db *gorm.DB
}

func NewPositionRepository() positionDomain.PositionRepository {
	return &PositionRepository{
		db: db.GetDB(),
	}
}

func (r *PositionRepository) Create(ctx context.Context, position *positionDomain.Position) error {
	err := r.db.Create(position).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *PositionRepository) FindByID(ctx context.Context, id string) (*positionDomain.Position, error) {
	var position positionDomain.Position
	err := r.db.Preload("Policies").Begin().Where("id = ?", id).First(&position).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorDomain.WrapError(errorDomain.NotFoundErr, err)
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return &position, nil
}

func (r *PositionRepository) FindByFacilityID(ctx context.Context, facilityID string) ([]*positionDomain.Position, error) {
	var positions []*positionDomain.Position
	err := r.db.Preload("Policies").Begin().Where("facility_id = ?", facilityID).Find(&positions).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return positions, nil
}
