package repository

import (
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
	if err := r.db.Create(position).Error; err != nil {
		return err
	}
	return nil
}

func (r *PositionRepository) FindByID(ctx context.Context, id string) (*positionDomain.Position, error) {
	var position positionDomain.Position
	if err := r.db.Preload("Policies").Begin().Where("id = ?", id).First(&position).Error; err != nil {
		return nil, err
	}
	return &position, nil
}

func (r *PositionRepository) FindByFacilityID(ctx context.Context, facilityID string) ([]*positionDomain.Position, error) {
	var positions []*positionDomain.Position
	if err := r.db.Preload("Policies").Begin().Where("facility_id = ?", facilityID).Find(&positions).Error; err != nil {
		return nil, err
	}
	return positions, nil
}
