package auth

// SignInRequest is the request structure for signing in
type SignInRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignUpRequest struct {
	Username     string  `json:"username" validate:"required"`
	Password     string  `json:"password" validate:"required"`
	FacilityID   string  `json:"facility_id" validate:"required,ulid"`
	DepartmentID string  `json:"department_id" validate:"required,ulid"`
	PositionID   string  `json:"position_id" validate:"required,ulid"`
	TeamID       string  `json:"team_id" validate:"required,ulid"`
	AreaID       string  `json:"area_id" validate:"required,ulid"`
	Email        *string `json:"email" validate:"omitempty,email"`
	PhoneNumber  *string `json:"phone_number" validate:"omitempty"`
}
