package repository

import (
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

func (r *RecurringScheduleRepository) Create(ctx context.Context, tx *gorm.DB, recurringSchedule *recurringScheduleDomain.RecurringSchedule) error {
	db := tx
	if db == nil {
		db = r.db
	}
	err := db.Create(recurringSchedule).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *RecurringScheduleRepository) FindByFacilityID(ctx context.Context, facility_id string, filters []query.Filter, sort query.SortOption) ([]*recurringScheduleDomain.RecurringSchedule, error) {
	dbQuery := r.db.Table("recurring_schedules") // 明示的にテーブルを指定

	// フィルタとソートを適用
	dbQuery = db.ApplyFiltersAndSort(dbQuery, filters, sort, recurringScheduleDomain.RecurringScheduleRelationMappings)

	var recurringSchedules []*recurringScheduleDomain.RecurringSchedule
	err := dbQuery.
		Preload("Facility").
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
		Preload("RecurringRule").
		Where("recurring_schedules.facility_id = ?", facility_id).Find(&recurringSchedules).Error
	if err != nil {
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return recurringSchedules, nil
}

func (r *RecurringScheduleRepository) FindByID(ctx context.Context, id string) (*recurringScheduleDomain.RecurringSchedule, error) {
	var recurringSchedule *recurringScheduleDomain.RecurringSchedule
	err := r.db.
		Preload("Facility").
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
		Preload("RecurringRule").
		Where("id = ?", id).First(&recurringSchedule).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return recurringSchedule, nil
}
