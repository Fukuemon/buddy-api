package repository

import (
	errorDomain "api-buddy/domain/error"
	scheduleTypeDomain "api-buddy/domain/schedule/schedule_type"
	"api-buddy/infrastructure/mysql/db"

	"gorm.io/gorm"
)

type ScheduleTypeRepository struct {
	db *gorm.DB
}

func NewScheduleTypeRepository() scheduleTypeDomain.ScheduleTypeRepository {
	return &ScheduleTypeRepository{
		db: db.GetDB(),
	}
}

func (r *ScheduleTypeRepository) FindAll(facility_id string) ([]*scheduleTypeDomain.ScheduleType, error) {
	var scheduleTypes []*scheduleTypeDomain.ScheduleType
	err := r.db.Find(&scheduleTypes).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return scheduleTypes, nil
}

func (r *ScheduleTypeRepository) FindByID(id string) (*scheduleTypeDomain.ScheduleType, error) {
	var scheduleType scheduleTypeDomain.ScheduleType
	err := r.db.Where("id = ?", id).First(&scheduleType).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return &scheduleType, nil
}
