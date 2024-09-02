package user

import (
	"api-buddy/domain/common"
	policyDomain "api-buddy/domain/policy"

	"github.com/Fukuemon/go-pkg/ulid"
)

type Option struct {
	Email       *string
	PhoneNumber *string
}

type User struct {
	ID          string `gorm:"primaryKey"`
	Username    string `gorm:"unique"`
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"unique"`
	PositionID  string
	TeamID      string
	FacilityID  string
	AreaID      string
	Policies    []*policyDomain.Policy `gorm:"many2many:user_policies;"`
	common.CommonModel
}

func Reconstruct(
	ID string,
	username string,
	positionID string,
	teamID string,
	facilityID string,
	areaID string,
	policies []*policyDomain.Policy,
	options *Option,
) (*User, error) {
	return newUser(
		ID,
		username,
		positionID,
		teamID,
		facilityID,
		areaID,
		policies,
		options,
	)
}

func NewUser(
	username string,
	positionID string,
	teamID string,
	facilityID string,
	areaID string,
	policies []*policyDomain.Policy,
	options *Option,
) (*User, error) {
	return newUser(
		ulid.NewULID(),
		username,
		positionID,
		teamID,
		facilityID,
		areaID,
		policies,
		options,
	)
}

func newUser(
	ID string,
	username string,
	positionID string,
	teamID string,
	facilityID string,
	areaID string,
	policies []*policyDomain.Policy,
	options *Option,
) (*User, error) {
	user := &User{
		ID:         ID,
		Username:   username,
		PositionID: positionID,
		TeamID:     teamID,
		FacilityID: facilityID,
		AreaID:     areaID,
		Policies:   policies,
	}

	if options != nil {
		options = &Option{}
	}

	if options.Email == nil {
		user.Email = *options.Email
	}

	if options.PhoneNumber == nil {
		user.PhoneNumber = *options.PhoneNumber
	}

	common.InitializeCommonModel(&user.CommonModel)

	return user, nil
}
