package area

type CreateAreaResponse struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	FacilityID string   `json:"facility_id"`
	AddressIDs []string `json:"address_ids"`
}

type AreaResponse struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	FacilityID string         `json:"facility_id"`
	Addresses  []AddressModel `json:"addresses"`
}

type AddressModel struct {
	ID           string  `json:"id"`
	Prefecture   string  `json:"prefecture"`
	City         string  `json:"city"`
	AddressLine1 string  `json:"address_line1"`
	AddressLine2 string  `json:"address_line2"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

type AreaListResponse []AreaResponse
