package repository

import (
	departmentDomain "api-buddy/domain/facility/department"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository() departmentDomain.DepartmentRepository {
	return &DepartmentRepository{
		db: db.GetDB(),
	}
}

func (r *DepartmentRepository) Create(ctx context.Context, department *departmentDomain.Department) error {
	err := r.db.Create(&department).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *DepartmentRepository) FindByID(ctx context.Context, id string) (*departmentDomain.Department, error) {
	var department *departmentDomain.Department
	err := r.db.Where("id = ?", id).First(&department).Error
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (r *DepartmentRepository) FindByFacilityID(ctx context.Context, facilityID string) ([]*departmentDomain.Department, error) {
	var departments []*departmentDomain.Department
	err := r.db.Where("facility_id = ?", facilityID).Find(&departments).Error
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (r *DepartmentRepository) FindAll(ctx context.Context) ([]*departmentDomain.Department, error) {
	var departments []*departmentDomain.Department
	err := r.db.Find(&departments).Error
	if err != nil {
		return nil, err
	}
	return departments, nil
}
