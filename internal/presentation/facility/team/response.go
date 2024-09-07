package team

import "time"

type CreateTeamResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FacilityID string `json:"facility_id"`
}

type TeamResponse struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	FacilityID string     `json:"facility_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type TeamListResponse []TeamResponse

type ErrorResponse struct {
	Error string `json:"error"`
}
