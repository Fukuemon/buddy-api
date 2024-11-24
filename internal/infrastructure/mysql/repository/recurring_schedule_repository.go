package repository

import (
	"api-buddy/domain/common"
	errorDomain "api-buddy/domain/error"
	recurringScheduleDomain "api-buddy/domain/schedule/recurring_schedule"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"github.com/Fukuemon/go-pkg/query"
	"gorm.io/gorm"
)

type RecurringScheduleRepository struct {
	db *gorm.DB
}

func NewRecurringScheduleRepository() recurringScheduleDomain.RecurringScheduleRepository {
	return &RecurringScheduleRepository{
		db: db.GetDB(),
	}
}

func (r *RecurringScheduleRepository) Create(ctx context.Context, recurringSchedule *recurringScheduleDomain.RecurringSchedule) error {
	err := r.db.Create(recurringSchedule).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}

	return nil
}

func (r *RecurringScheduleRepository) FindByFacilityID(ctx context.Context, facility_id string, filters []query.Filter, sort query.SortOption) ([]*recurringScheduleDomain.RecurringSchedule, error) {
	q := query.NewQuery()

	// Queryオブジェクトにフィルターを適用
	for _, filter := range filters {
		filter.Apply(q)
	}

	dbQuery := r.db

	for _, mapping := range recurringScheduleDomain.RecurringScheduleRelationMappings {
		if value, exists := q.Filters[mapping.FilterField]; exists {
			dbQuery = dbQuery.Joins("JOIN "+mapping.TableName+" ON "+mapping.JoinKey).
				Where(mapping.FilterField+" = ?", value)
		}
	}

	for key, value := range q.Filters {
		if _, isRelationField := recurringScheduleDomain.RecurringScheduleRelationMappings[key]; !isRelationField {
			dbQuery = dbQuery.Where(key, value)
		}
	}

	if sort.Field != "" {
		if mapping, exists := recurringScheduleDomain.RecurringScheduleRelationMappings[sort.Field]; exists {
			dbQuery = dbQuery.Joins("JOIN " + mapping.TableName + " ON " + mapping.JoinKey).
				Order(mapping.FilterField + " " + sort.Order)
		} else {
			dbQuery = dbQuery.Order(sort.Field + " " + sort.Order)
		}
	} else {
		dbQuery = dbQuery.Order(common.UpdatedAt + " " + query.DESC)
	}

	var recurringSchedules []*recurringScheduleDomain.RecurringSchedule
	err := dbQuery.Preload("Facility").Preload("RecurringRule").Preload("VisitInfo").Preload("ScheduleType").Preload("Staff").
		Where("facility_id = ?", facility_id).Find(&recurringSchedules).Error
	if err != nil {
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return recurringSchedules, nil
}

func (r *RecurringScheduleRepository) FindByID(ctx context.Context, id string) (*recurringScheduleDomain.RecurringSchedule, error) {
	var recurringSchedule *recurringScheduleDomain.RecurringSchedule
	err := r.db.Preload("Facility").Preload("RecurringRule").Preload("VisitInfo").Preload("ScheduleType").Preload("Staff").
		Where("id = ?", id).First(&recurringSchedule).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return recurringSchedule, nil
}
