package repository

import (
	errorDomain "api-buddy/domain/error"
	teamDomain "api-buddy/domain/facility/team"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository() teamDomain.TeamRepository {
	return &TeamRepository{
		db: db.GetDB(),
	}
}

func (r *TeamRepository) Create(ctx context.Context, team *teamDomain.Team) error {
	err := r.db.Create(team).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *TeamRepository) FindByID(ctx context.Context, id string) (*teamDomain.Team, error) {
	var team teamDomain.Team
	err := r.db.Where("id = ?", id).First(&team).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorDomain.WrapError(errorDomain.NotFoundErr, err)
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return &team, nil
}

func (r *TeamRepository) FindByFacilityID(ctx context.Context, facilityID string) ([]*teamDomain.Team, error) {
	var teams []*teamDomain.Team
	err := r.db.Where("facility_id = ?", facilityID).Find(&teams).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return teams, nil
}

func (r *TeamRepository) FindAll(ctx context.Context) ([]*teamDomain.Team, error) {
	var teams []*teamDomain.Team
	err := r.db.Find(&teams).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return teams, nil
}
