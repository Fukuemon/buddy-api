package repository

import (
	"api-buddy/domain/common"
	errorDomain "api-buddy/domain/error"
	patientDomain "api-buddy/domain/patient"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"github.com/Fukuemon/go-pkg/query"
	"gorm.io/gorm"
)

type PatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository() patientDomain.PatientRepository {
	return &PatientRepository{
		db: db.GetDB(),
	}
}

func (r *PatientRepository) Create(ctx context.Context, patient *patientDomain.Patient) error {
	err := r.db.Create(patient).Error
	if err != nil {
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *PatientRepository) FindByFacilityID(ctx context.Context, facility_id string, filters []query.Filter, sort query.SortOption) ([]*patientDomain.Patient, error) {
	q := query.NewQuery()

	for _, filter := range filters {
		filter.Apply(q)
	}

	dbQuery := r.db

	for _, mapping := range patientDomain.PatientRelationMappings {
		if value, exists := q.Filters[mapping.FilterField]; exists {
			dbQuery = dbQuery.Joins("JOIN "+mapping.TableName+" ON "+mapping.JoinKey).
				Where(mapping.FilterField+" = ?", value)
		}
	}

	for key, value := range q.Filters {
		if _, isRelationField := patientDomain.PatientRelationMappings[key]; !isRelationField {
			dbQuery = dbQuery.Where(key, value)
		}
	}

	if sort.Field != "" {
		if mapping, exists := patientDomain.PatientRelationMappings[sort.Field]; exists {
			dbQuery = dbQuery.Joins("JOIN " + mapping.TableName + " ON " + mapping.JoinKey).
				Order(mapping.FilterField + " " + sort.Order)
		} else {
			dbQuery = dbQuery.Order(sort.Field + " " + sort.Order)
		}
	} else {
		dbQuery = dbQuery.Order(common.UpdatedAt + " " + query.DESC)
	}

	var patients []*patientDomain.Patient
	err := dbQuery.Preload("ServiceCode").Preload("Address").Preload("Area").Preload("Assigned_Staff").
		Where("facility_id = ?", facility_id).
		Find(&patients).Error
	if err != nil {
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return patients, nil
}

func (r *PatientRepository) FindByID(ctx context.Context, id string) (*patientDomain.Patient, error) {
	var patient patientDomain.Patient
	err := r.db.Preload("ServiceCode").Preload("Address").Preload("Area").Preload("AssignedStaff").
		Where("id = ?", id).
		First(&patient).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorDomain.WrapError(errorDomain.NotFoundErr, err)
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return &patient, nil
}
