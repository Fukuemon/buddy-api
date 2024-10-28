package facility

import "api-buddy/domain/common"

type Facility struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	common.CommonModel
}

func NewFacility(name string) (*Facility, error) {
	return newFacility(name)
}

func newFacility(name string) (*Facility, error) {
	facility := &Facility{
		Name: name,
	}

	common.InitializeCommonModel(&facility.CommonModel)

	return facility, nil
}
