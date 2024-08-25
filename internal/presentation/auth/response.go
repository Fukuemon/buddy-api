package auth

// SignInResponse is the response structure for a successful sign-in
type SignInResponse struct {
	AccessToken string `json:"access_token"`
	IdToken     string `json:"id_token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
