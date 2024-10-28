package team

import (
	"api-buddy/domain/common"

	"github.com/Fukuemon/go-pkg/ulid"
)

type Team struct {
	ID         string `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	FacilityID string `gorm:"not null"`
	common.CommonModel
}

func NewTeam(name string, facilityID string) (*Team, error) {
	return newTeam(
		ulid.NewULID(),
		name,
		facilityID,
	)
}

func newTeam(ID string, name string, facilityID string) (*Team, error) {
	team := &Team{
		ID:         ID,
		Name:       name,
		FacilityID: facilityID,
	}

	common.InitializeCommonModel(&team.CommonModel)

	return team, nil
}
