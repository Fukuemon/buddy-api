package repository

import (
	errorDomain "api-buddy/domain/error"
	scheduleDomain "api-buddy/domain/schedule"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"github.com/Fukuemon/go-pkg/query"
	"gorm.io/gorm"
)

type ScheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository() scheduleDomain.ScheduleRepository {
	return &ScheduleRepository{
		db: db.GetDB(),
	}
}

func (r *ScheduleRepository) Create(ctx context.Context, schedule *scheduleDomain.Schedule) error {
	err := r.db.Create(schedule).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *ScheduleRepository) FindByFacilityID(ctx context.Context, facility_id string, filters []query.Filter, sort query.SortOption) ([]*scheduleDomain.Schedule, error) {
	// Queryオブジェクトを初期化
	q := query.NewQuery()

	// Queryオブジェクトにフィルターを適用
	for _, filter := range filters {
		filter.Apply(q)
	}

	dbQuery := r.db

	// リレーションテーブルに基づいたフィルタリングのQueryを適用
	for _, mapping := range scheduleDomain.ScheduleRelationMappings {
		if value, exists := q.Filters[mapping.FilterField]; exists {
			dbQuery = dbQuery.Joins("JOIN "+mapping.TableName+" ON "+mapping.JoinKey).
				Where(mapping.FilterField+" = ?", value)
		}
	}

	// 残りのフィルタリング（`schedules` テーブルに対するフィルター）を適用
	for key, value := range q.Filters {
		if _, isRelationField := scheduleDomain.ScheduleRelationMappings[key]; !isRelationField {
			dbQuery = dbQuery.Where(key, value)
		}
	}

	// ソートオプションの適用
	if sort.Field != "" {
		if mapping, exists := scheduleDomain.ScheduleRelationMappings[sort.Field]; exists {
			dbQuery = dbQuery.Joins("JOIN " + mapping.TableName + " ON " + mapping.JoinKey).
				Order(mapping.TableName + "." + sort.Field + " " + sort.Order)
		} else {
			dbQuery = dbQuery.Order(sort.Field + " " + sort.Order)
		}
	} else {
		dbQuery = dbQuery.Order(sort.Field + " " + sort.Order)
	}

	var schedules []*scheduleDomain.Schedule
	err := dbQuery.Preload("Facility").Preload("RecurringSchedule").Preload("VisitInfo").Preload("ScheduleType").Preload("Staff").Preload("ScheduleCancel").
		Find(&schedules).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return schedules, nil
}

func (r *ScheduleRepository) FindByID(ctx context.Context, id string) (*scheduleDomain.Schedule, error) {
	var schedule *scheduleDomain.Schedule
	err := r.db.Preload("Facility").Preload("RecurringSchedule").Preload("VisitInfo").Preload("ScheduleType").Preload("Staff").Preload("ScheduleCancel").
		Where("id = ?", id).First(&schedule).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return schedule, nil
}
