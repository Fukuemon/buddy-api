package auth

import "api-buddy/domain/policy"

// SignInResponse is the response structure for a successful sign-in
type SignInResponse struct {
	AccessToken string `json:"access_token"`
	IdToken     string `json:"id_token"`
}

type SignUpResponse struct {
	ID             string           `json:"id"`
	UserName       string           `json:"name"`
	Email          *string          `json:"email"`
	PhoneNumber    *string          `json:"phone_number"`
	FacilityName   string           `json:"facility"`
	DepartmentName string           `json:"department"`
	PositionName   string           `json:"position"`
	TeamName       string           `json:"team"`
	Policies       []*policy.Policy `json:"policies"`
}
