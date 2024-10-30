package address

import (
	"api-buddy/domain/common"

	"github.com/Fukuemon/go-pkg/ulid"
)

type Address struct {
	ID           string `gorm:"primaryKey"`
	ZipCode      string `gorm:"not null"`
	Prefecture   string `gorm:"not null"`
	City         string `gorm:"not null"`
	AddressLine1 string `gorm:"not null"`
	AddressLine2 string `gorm:"not null"`
	Latitude     float64
	Longitude    float64
	common.CommonModel
}

func NewAddress(
	ZipCode string,
	prefecture string,
	city string,
	addressLine1 string,
	addressLine2 string,
	latitude float64,
	longitude float64,
) (*Address, error) {
	return newAddress(
		ulid.NewULID(),
		ZipCode,
		prefecture,
		city,
		addressLine1,
		addressLine2,
		latitude,
		longitude,
	)
}

func newAddress(
	id string,
	ZipCode string,
	prefecture string,
	city string,
	addressLine1 string,
	addressLine2 string,
	latitude float64,
	longitude float64,
) (*Address, error) {
	address := &Address{
		ID:           id,
		ZipCode:      ZipCode,
		Prefecture:   prefecture,
		City:         city,
		AddressLine1: addressLine1,
		AddressLine2: addressLine2,
		Latitude:     latitude,
		Longitude:    longitude,
	}

	common.InitializeCommonModel(&address.CommonModel)
	return address, nil
}
