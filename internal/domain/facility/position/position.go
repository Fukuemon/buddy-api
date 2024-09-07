package position

import (
	"api-buddy/domain/common"
	policyDomain "api-buddy/domain/policy"

	"github.com/Fukuemon/go-pkg/ulid"
)

type Position struct {
	ID         string                 `gorm:"primaryKey"`
	Name       string                 `gorm:"not null"`
	FacilityID string                 `gorm:"not null"`
	Policies   []*policyDomain.Policy `gorm:"many2many:position_policies;"`
	common.CommonModel
}

func NewPosition(name string, facilityID string, policies []*policyDomain.Policy) (*Position, error) {
	return newPosition(ulid.NewULID(), name, facilityID, policies)
}

func newPosition(ID string, name string, facilityID string, policies []*policyDomain.Policy) (*Position, error) {
	position := &Position{
		ID:         ID,
		Name:       name,
		FacilityID: facilityID,
		Policies:   policies,
	}

	common.InitializeCommonModel(&position.CommonModel)

	return position, nil
}
