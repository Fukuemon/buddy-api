package repository

import (
	errorDomain "api-buddy/domain/error"
	serviceCodeDomain "api-buddy/domain/visit_info/service_code"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type ServiceCodeRepository struct {
	db *gorm.DB
}

func NewServiceCodeRepository() serviceCodeDomain.ServiceCodeRepository {
	return &ServiceCodeRepository{
		db: db.GetDB(),
	}
}

func (r *ServiceCodeRepository) Create(ctx context.Context, serviceCode *serviceCodeDomain.ServiceCode) error {
	err := r.db.Create(serviceCode).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *ServiceCodeRepository) FindAll(ctx context.Context) ([]*serviceCodeDomain.ServiceCode, error) {
	var serviceCodes []*serviceCodeDomain.ServiceCode
	err := r.db.Find(&serviceCodes).Error
	if err != nil {
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return serviceCodes, nil
}

func (r *ServiceCodeRepository) FindByID(ctx context.Context, id string) (*serviceCodeDomain.ServiceCode, error) {
	var serviceCode serviceCodeDomain.ServiceCode
	err := r.db.Where("id = ?", id).First(&serviceCode).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorDomain.WrapError(errorDomain.NotFoundErr, err)
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return &serviceCode, nil
}
