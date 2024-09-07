package position

type CreatePositionRequest struct {
	Name      string   `json:"name" binding:"required"`
	PolicyIDs []string `json:"policy_ids" binding:"required"`
}
