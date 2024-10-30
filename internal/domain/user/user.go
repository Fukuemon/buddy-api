package user

import (
	"api-buddy/domain/common"
	facilityDomain "api-buddy/domain/facility"
	areaDomain "api-buddy/domain/facility/area"
	departmentDomain "api-buddy/domain/facility/department"
	positionDomain "api-buddy/domain/facility/position"
	teamDomain "api-buddy/domain/facility/team"
	policyDomain "api-buddy/domain/policy"

	errorDomain "api-buddy/domain/error"

	"github.com/Fukuemon/go-pkg/query"
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
	Email        string
	PhoneNumber  string
	FacilityID   string
	Facility     *facilityDomain.Facility `gorm:"foreignKey:FacilityID"`
	DepartmentID string
	Department   *departmentDomain.Department `gorm:"foreignKey:DepartmentID"`
	PositionID   string
	Position     *positionDomain.Position `gorm:"foreignKey:PositionID"`
	TeamID       string
	Team         *teamDomain.Team `gorm:"foreignKey:TeamID"`
	AreaID       string
	Area         *areaDomain.Area       `gorm:"foreignKey:AreaID"`
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
	area *areaDomain.Area,
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
		area,
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
	area *areaDomain.Area,
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
		area,
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
	area *areaDomain.Area,
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
		AreaID:       area.ID,
		Area:         area,
		Policies:     policies,
	}

	if len(user.Username) > 128 {
		err := errorDomain.NewError("ユーザー名は128文字以内です")
		return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
	}

	// Optionがnilでない場合のみ、EmailとPhoneNumberを設定
	if options != nil {
		if options.Email != nil {
			if len(*options.Email) > 320 {
				err := errorDomain.NewError("Emailは320文字以内です")
				return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
			}
			user.Email = *options.Email
		}

		if options.PhoneNumber != nil {
			if len(*options.PhoneNumber) > 11 {
				err := errorDomain.NewError("電話番号は11文字以内です")
				return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
			}
			// 先頭に"+"を加える処理を追加
			*options.PhoneNumber = common.AddPlusToPhoneNumber(*options.PhoneNumber)
			user.PhoneNumber = *options.PhoneNumber
		}
	}

	common.InitializeCommonModel(&user.CommonModel)

	return user, nil
}

var UserRelationMappings = map[string]query.RelationMapping{
	"position": {
		TableName:   "positions",
		JoinKey:     "positions.id = users.position_id",
		FilterField: "positions.name",
	},
	"department": {
		TableName:   "departments",
		JoinKey:     "departments.id = users.department_id",
		FilterField: "departments.name",
	},
	"team": {
		TableName:   "teams",
		JoinKey:     "teams.id = users.team_id",
		FilterField: "teams.name",
	},
	"area": {
		TableName:   "areas",
		JoinKey:     "areas.id = users.area_id",
		FilterField: "areas.name",
	},
}
