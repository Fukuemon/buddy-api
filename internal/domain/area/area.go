package area

import "api-buddy/domain/common"

type Area struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	common.CommonModel
}

func NewArea(name string) (*Area, error) {
	return newArea(name)
}

func newArea(name string) (*Area, error) {
	area := &Area{
		Name: name,
	}

	common.InitializeCommonModel(&area.CommonModel)

	return area, nil
}
