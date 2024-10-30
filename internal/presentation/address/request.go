package address

type CreateAddressRequest struct {
	ZipCode      string  `json:"zip_code" validate:"required"`
	Prefecture   string  `json:"prefecture" validate:"required"`
	City         string  `json:"city" validate:"required"`
	AddressLine1 string  `json:"address_line1" validate:"required"`
	AddressLine2 string  `json:"address_line2" validate:"required"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}
