package policy

import "time"

type CreatePolicyResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PolicyResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PolicyListResponse []PolicyResponse
