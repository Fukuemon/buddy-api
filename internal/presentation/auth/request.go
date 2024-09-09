package auth

// SignInRequest is the request structure for signing in
type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpRequest struct {
	Username     string  `json:"username" binding:"required"`
	Password     string  `json:"password" binding:"required"`
	FacilityID   string  `json:"facility_id" binding:"required"`
	DepartmentID string  `json:"department_id" binding:"required"`
	PositionID   string  `json:"position_id" binding:"required"`
	TeamID       string  `json:"team_id" binding:"required"`
	Email        *string `json:"email"`
	PhoneNumber  *string `json:"phone_number"`
}
