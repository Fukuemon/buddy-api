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

func (r *ScheduleRepository) Create(ctx context.Context, tx *gorm.DB, schedule *scheduleDomain.Schedule) error {
	// トランザクションが渡されていない場合、通常のDB接続を使用
	db := tx
	if db == nil {
		db = r.db
	}

	err := db.Create(schedule).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *ScheduleRepository) FindByFacilityID(ctx context.Context, facility_id string, filters []query.Filter, sort query.SortOption) ([]*scheduleDomain.Schedule, error) {
	dbQuery := r.db.Table("schedules") // 明示的にテーブルを指定

	// フィルタとソートを適用
	dbQuery = db.ApplyFiltersAndSort(dbQuery, filters, sort, scheduleDomain.ScheduleRelationMappings)

	var schedules []*scheduleDomain.Schedule
	err := dbQuery.
		Preload("Facility").
		Preload("RecurringSchedule").
		Preload("VisitInfo").
		Preload("VisitInfo.Patient").
		Preload("VisitInfo.AssignedStaff").
		Preload("VisitInfo.Companion").
		Preload("VisitInfo.Route").
		Preload("VisitInfo.Route.Address").
		Preload("VisitInfo.Route.Destination").
		Preload("VisitInfo.ServiceCode").
		Preload("VisitInfo.VisitCategories").
		Preload("ScheduleType").
		Preload("Staff").
		Preload("ScheduleCancel").
		Where("schedules.facility_id = ?", facility_id).Find(&schedules).Error
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
	err := r.db.
		Preload("Facility").
		Preload("RecurringSchedule").
		Preload("VisitInfo").
		Preload("VisitInfo.Patient").
		Preload("VisitInfo.AssignedStaff").
		Preload("VisitInfo.Companion").
		Preload("VisitInfo.Route").
		Preload("VisitInfo.Route.Address").
		Preload("VisitInfo.Route.Destination").
		Preload("VisitInfo.ServiceCode").
		Preload("VisitInfo.VisitCategories").
		Preload("ScheduleType").
		Preload("Staff").
		Preload("ScheduleCancel").
		Where("id = ?", id).First(&schedule).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return schedule, nil
}
