package team

type CreateTeamRequest struct {
	Name string `json:"name" validate:"required"`
}
