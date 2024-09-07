package position

import (
	"api-buddy/usecase/facility/position"
	"time"
)

type CreatePositionResponse struct {
	ID       string               `json:"id"`
	Name     string               `json:"name" binding:"required"`
	Policies []position.PolicyDto `json:"policies"`
}

type PositionResponse struct {
	ID         string               `json:"id"`
	Name       string               `json:"name"`
	FacilityID string               `json:"facility_id"`
	Policies   []position.PolicyDto `json:"policies"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
	DeletedAt  *time.Time           `json:"deleted_at"`
}

type PositionListResponse []PositionResponse

type ErrorResponse struct {
	Error string `json:"error"`
}
