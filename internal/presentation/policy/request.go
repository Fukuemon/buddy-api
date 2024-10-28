package policy

type CreatePolicyRequest struct {
	Name string `json:"name" validate:"required"`
}
