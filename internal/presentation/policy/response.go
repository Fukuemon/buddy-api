package policy

import "time"

type CreatePolicyResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PolicyResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type PolicyListResponse []PolicyResponse

type ErrorResponse struct {
	Error string `json:"error"`
}