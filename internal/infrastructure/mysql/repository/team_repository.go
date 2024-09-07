package repository

import (
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
	if err := r.db.Create(team).Error; err != nil {
		return err
	}
	return nil
}

func (r *TeamRepository) FindByID(ctx context.Context, id string) (*teamDomain.Team, error) {
	var team teamDomain.Team
	if err := r.db.Where("id = ?", id).First(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *TeamRepository) FindByFacilityID(ctx context.Context, facilityID string) ([]*teamDomain.Team, error) {
	var teams []*teamDomain.Team
	if err := r.db.Where("facility_id = ?", facilityID).Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

func (r *TeamRepository) FindAll(ctx context.Context) ([]*teamDomain.Team, error) {
	var teams []*teamDomain.Team
	if err := r.db.Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}
