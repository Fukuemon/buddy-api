package repository

import (
	policyDomain "api-buddy/domain/policy"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type policyRepository struct {
	db *gorm.DB
}

func NewPolicyRepository() policyDomain.PolicyRepository {
	return &policyRepository{
		db: db.GetDB(),
	}
}

func (r *policyRepository) Create(ctx context.Context, policy *policyDomain.Policy) error {
	err := r.db.Create(&policy).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *policyRepository) FindByID(ctx context.Context, id string) (*policyDomain.Policy, error) {
	var policy policyDomain.Policy
	err := r.db.Where("id = ?", id).First(&policy).Error
	if err != nil {
		return nil, err
	}
	return &policy, nil
}

func (r *policyRepository) FindByIDs(ctx context.Context, ids []string) ([]*policyDomain.Policy, error) {
	var policies []*policyDomain.Policy
	err := r.db.Where("id IN ?", ids).Find(&policies).Error
	if err != nil {
		return nil, err
	}
	return policies, nil
}

func (r *policyRepository) FindByPositionID(ctx context.Context, positionID string) ([]*policyDomain.Policy, error) {
	var policies []*policyDomain.Policy
	err := r.db.Preload("Positions").Begin().Where("position_id = ?", positionID).Find(&policies).Error
	if err != nil {
		return nil, err
	}
	return policies, nil
}

func (r *policyRepository) FindAll(ctx context.Context) ([]*policyDomain.Policy, error) {
	var policies []*policyDomain.Policy
	err := r.db.Find(&policies).Error
	if err != nil {
		return nil, err
	}
	return policies, nil
}
