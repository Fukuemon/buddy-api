package auth

// SignInRequest is the request structure for signing in
type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
