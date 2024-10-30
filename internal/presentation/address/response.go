package address

import "time"

type CreateAddressResponse struct {
	ID           string
	ZipCode      string
	Prefecture   string
	City         string
	AddressLine1 string
	AddressLine2 string
	Latitude     float64
	Longitude    float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AddressResponse struct {
	ID           string
	ZipCode      string
	Prefecture   string
	City         string
	AddressLine1 string
	AddressLine2 string
	Latitude     float64
	Longitude    float64
}

type AddressDetailResponse struct {
	ID           string
	ZipCode      string
	Prefecture   string
	City         string
	AddressLine1 string
	AddressLine2 string
	Latitude     float64
	Longitude    float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AddressListResponse struct {
	Addresses []AddressResponse `json:"addresses"`
}
