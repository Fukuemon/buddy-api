package user

import (
	"api-buddy/domain/common"
	facilityDomain "api-buddy/domain/facility"
	departmentDomain "api-buddy/domain/facility/department"
	positionDomain "api-buddy/domain/facility/position"
	teamDomain "api-buddy/domain/facility/team"
	policyDomain "api-buddy/domain/policy"

	"github.com/Fukuemon/go-pkg/ulid"
)

type Option struct {
	Email       *string
	PhoneNumber *string
}

// Todo: 担当エリアの追加
type User struct {
	ID           string `gorm:"primaryKey"`
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	PhoneNumber  string `gorm:"unique"`
	FacilityID   string
	Facility     *facilityDomain.Facility `gorm:"foreignKey:FacilityID"`
	DepartmentID string
	Department   *departmentDomain.Department `gorm:"foreignKey:DepartmentID"`
	PositionID   string
	Position     *positionDomain.Position `gorm:"foreignKey:PositionID"`
	TeamID       string
	Team         *teamDomain.Team       `gorm:"foreignKey:TeamID"`
	Policies     []*policyDomain.Policy `gorm:"many2many:user_policies;"`
	common.CommonModel
}

func Reconstruct(
	ID string,
	username string,
	position *positionDomain.Position,
	team *teamDomain.Team,
	facility *facilityDomain.Facility,
	department *departmentDomain.Department,
	policies []*policyDomain.Policy,
	options *Option,
) (*User, error) {
	return newUser(
		ID,
		username,
		position,
		team,
		facility,
		department,
		policies,
		options,
	)
}

func NewUser(
	username string,
	position *positionDomain.Position,
	team *teamDomain.Team,
	facility *facilityDomain.Facility,
	department *departmentDomain.Department,
	policies []*policyDomain.Policy,
	options *Option,
) (*User, error) {
	return newUser(
		ulid.NewULID(),
		username,
		position,
		team,
		facility,
		department,
		policies,
		options,
	)
}

func newUser(
	ID string,
	username string,
	position *positionDomain.Position,
	team *teamDomain.Team,
	facility *facilityDomain.Facility,
	department *departmentDomain.Department,
	policies []*policyDomain.Policy,
	options *Option,
) (*User, error) {
	user := &User{
		ID:           ID,
		Username:     username,
		PositionID:   position.ID,
		TeamID:       team.ID,
		FacilityID:   facility.ID,
		DepartmentID: department.ID,
		Position:     position,
		Team:         team,
		Facility:     facility,
		Department:   department,
		Policies:     policies,
	}

	// Optionがnilでない場合のみ、EmailとPhoneNumberを設定
	if options != nil {
		if options.Email != nil {
			user.Email = *options.Email
		}

		if options.PhoneNumber != nil {
			user.PhoneNumber = *options.PhoneNumber
		}
	}

	common.InitializeCommonModel(&user.CommonModel)

	return user, nil
}
