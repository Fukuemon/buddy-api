package area

import (
	addressDomain "api-buddy/domain/address"
	"api-buddy/domain/common"

	"github.com/Fukuemon/go-pkg/ulid"
)

type Area struct {
	ID         string                   `gorm:"primaryKey"`
	Name       string                   `gorm:"not null"`
	FacilityID string                   `gorm:"not null"`
	Addresses  []*addressDomain.Address `gorm:"many2many:area_addresses;"`
	common.CommonModel
}

func NewArea(name string, facilityID string, addresses []*addressDomain.Address) (*Area, error) {
	return newArea(ulid.NewULID(), name, facilityID, addresses)
}

func newArea(id string, name string, facilityID string, addresses []*addressDomain.Address) (*Area, error) {
	area := &Area{
		ID:         id,
		Name:       name,
		FacilityID: facilityID,
		Addresses:  addresses,
	}
	common.InitializeCommonModel(&area.CommonModel)
	return area, nil
}
