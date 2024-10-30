package repository

import (
	addressDomain "api-buddy/domain/address"
	"api-buddy/infrastructure/mysql/db"
	"context"

	errorDomain "api-buddy/domain/error"

	"github.com/Fukuemon/go-pkg/query"
	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository() addressDomain.AddressRepository {
	return &AddressRepository{
		db: db.GetDB(),
	}
}

func (r *AddressRepository) Create(ctx context.Context, address *addressDomain.Address) error {
	err := r.db.Create(address).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *AddressRepository) FindByID(ctx context.Context, id string) (*addressDomain.Address, error) {
	var address addressDomain.Address
	err := r.db.Where("id = ?", id).First(&address).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorDomain.WrapError(errorDomain.NotFoundErr, err)
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return &address, nil
}

func (r *AddressRepository) FindByIDs(ctx context.Context, ids []string) ([]*addressDomain.Address, error) {
	var addresses []*addressDomain.Address
	err := r.db.Where("id IN ?", ids).Find(&addresses).Error
	if err != nil {
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return addresses, nil
}

func (r *AddressRepository) FindByAreaID(ctx context.Context, areaID string) ([]*addressDomain.Address, error) {
	var addresses []*addressDomain.Address
	err := r.db.Preload("Areas").Begin().Where("area_id = ?", areaID).Find(&addresses).Error
	if err != nil {
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return addresses, nil
}

func (r *AddressRepository) Fetch(ctx context.Context, filters []query.Filter) ([]*addressDomain.Address, error) {
	q := query.NewQuery()

	for _, filter := range filters {
		filter.Apply(q)
	}

	dbQuery := r.db

	for key, value := range q.Filters {
		dbQuery = dbQuery.Where(key, value)
	}

	var addresses []*addressDomain.Address
	err := dbQuery.Find(&addresses).Error
	if err != nil {
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return addresses, nil
}
