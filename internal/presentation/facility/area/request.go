package area

type CreateAreaRequest struct {
	Name       string   `json:"name" validate:"required"`
	FacilityID string   `json:"facility_id" validate:"required,ulid"`
	AddressIDs []string `json:"address_ids" validate:"required,dive,ulid"`
}
