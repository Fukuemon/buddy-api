package policy

type CreatePolicyRequest struct {
	Name string `json:"name" binding:"required"`
}
