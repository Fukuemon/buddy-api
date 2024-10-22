package position

type CreatePositionRequest struct {
	Name      string   `json:"name" validate:"required"`
	PolicyIDs []string `json:"policy_ids" validate:"required,dive,ulid"`
}
