package repository

import (
	errorDomain "api-buddy/domain/error"
	recurringRuleDomain "api-buddy/domain/schedule/recurring_rule"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type RecurringRuleRepository struct {
	db *gorm.DB
}

func NewRecurringRuleRepository() recurringRuleDomain.RecurringRuleRepository {
	return &RecurringRuleRepository{
		db: db.GetDB(),
	}
}

func (r *RecurringRuleRepository) Create(ctx context.Context, recurringRule *recurringRuleDomain.RecurringRule) error {
	err := r.db.Create(recurringRule).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}

	return nil
}

func (r *RecurringRuleRepository) FindByFacilityID(ctx context.Context, facilityID string) ([]*recurringRuleDomain.RecurringRule, error) {
	var recurringRules []*recurringRuleDomain.RecurringRule
	err := r.db.Where("facility_id = ?", facilityID).Find(&recurringRules).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return recurringRules, nil
}

func (r *RecurringRuleRepository) FindByID(ctx context.Context, id string) (*recurringRuleDomain.RecurringRule, error) {
	var recurringRule *recurringRuleDomain.RecurringRule
	err := r.db.Where("id = ?", id).First(&recurringRule).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return recurringRule, nil
}
