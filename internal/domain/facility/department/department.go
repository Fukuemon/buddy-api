package department

import (
	"api-buddy/domain/common"

	"github.com/Fukuemon/go-pkg/ulid"
)

type Department struct {
	ID         string `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	FacilityID string `gorm:"not null"`
	common.CommonModel
}

func NewDepartment(name string, facilityID string) (*Department, error) {
	return newDepartment(
		ulid.NewULID(),
		name,
		facilityID,
	)
}

func newDepartment(ID string, name string, facilityID string) (*Department, error) {
	department := &Department{
		ID:         ID,
		Name:       name,
		FacilityID: facilityID,
	}

	common.InitializeCommonModel(&department.CommonModel)

	return department, nil
}
