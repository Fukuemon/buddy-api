package user

import (
	"time"
)

type UserDetailResponse struct {
	ID          string        `json:"id"`
	Username    string        `json:"username"`
	Position    string        `json:"position"`
	Team        string        `json:"team"`
	Facility    string        `json:"facility"`
	Department  string        `json:"department"`
	Policies    []PolicyModel `json:"policies"`
	Email       *string       `json:"email"`
	PhoneNumber *string       `json:"phone"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type UserResponse struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Position   string `json:"position"`
	Team       string `json:"team"`
	Department string `json:"department"`
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
}

type PolicyModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
