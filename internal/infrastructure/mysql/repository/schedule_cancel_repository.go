package repository

import (
	errorDomain "api-buddy/domain/error"
	scheduleCancelDomain "api-buddy/domain/schedule/schedule_cancel"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type ScheduleCancelRepository struct {
	db *gorm.DB
}

func NewScheduleCancelRepository() scheduleCancelDomain.ScheduleCancelRepository {
	return &ScheduleCancelRepository{
		db: db.GetDB(),
	}
}

func (r *ScheduleCancelRepository) Create(ctx context.Context, scheduleCancel *scheduleCancelDomain.ScheduleCancel) error {
	err := r.db.Create(scheduleCancel).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *ScheduleCancelRepository) FindByID(ctx context.Context, id string) (*scheduleCancelDomain.ScheduleCancel, error) {
	var scheduleCancel *scheduleCancelDomain.ScheduleCancel
	err := r.db.Where("id = ?", id).First(&scheduleCancel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return scheduleCancel, nil
}
