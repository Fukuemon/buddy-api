package user_test

import (
	"api-buddy/domain/common"
	errorDomain "api-buddy/domain/error"
	facilityDomain "api-buddy/domain/facility"
	areaDomain "api-buddy/domain/facility/area"
	departmentDomain "api-buddy/domain/facility/department"
	positionDomain "api-buddy/domain/facility/position"
	teamDomain "api-buddy/domain/facility/team"
	"api-buddy/domain/policy"
	userDomain "api-buddy/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	// 正常値の定義
	trueName := "test_user"
	facility := &facilityDomain.Facility{ID: "facility_1"}
	department := &departmentDomain.Department{ID: "department_1"}
	position := &positionDomain.Position{ID: "position_1"}
	team := &teamDomain.Team{ID: "team_1"}
	area := &areaDomain.Area{ID: "area_1"}
	policies := []*policy.Policy{{ID: "policy_1"}, {ID: "policy_2"}}
	trueEmail := "test@example.com"
	truePhoneNumber := "09012345678"

	// テストケースの定義
	cases := map[string]struct {
		username     string
		options      *userDomain.Option
		wantErrorMsg string
	}{
		"正常系": {
			username: trueName,
			options: &userDomain.Option{
				Email:       &trueEmail,
				PhoneNumber: &truePhoneNumber,
			},
			wantErrorMsg: "",
		},
		"異常系: ユーザー名が128文字以上": {
			username: "Il2wTMIUYSgYc3oRBwKjx1O2STuuur32vXR06k3Pqu3OahfWE1xR7yWXeBft8utbkfQh1v66TowFixygHKuUbSIr5NUrs00NMCvaLlHSAsAL0QvLpCKn8Hj2jMxJhUOSp",
			options: &userDomain.Option{
				Email:       &trueEmail,
				PhoneNumber: &truePhoneNumber,
			},
			wantErrorMsg: errorDomain.NewError("ユーザー名は128文字以内です").Error(),
		},
		"異常系：メールアドレスが正しい形式でない": {
			username: trueName,
			options: &userDomain.Option{
				Email:       common.StringPointer("testexample.com"),
				PhoneNumber: &truePhoneNumber,
			},
			wantErrorMsg: errorDomain.NewError("Emailの形式が正しくありません").Error(),
		},
		"異常系: メールアドレスが320文字以上": {
			username: trueName,
			options: &userDomain.Option{
				Email:       common.StringPointer("MQ8GaEaSPw8p1ngbSsNXiFgRPK21fwPTIHl2N7aLUW8hOo5xi8UnAHHKqh62FBMTrmheUUbD6FCA0MP2yI2TbVCyxrUpokE0PcEcMQ5xdKn3YL0U7MdgYsiGDAvGJHSEVO4i5NcWs4X8NPMVRg4IW0McN8ga07pttL6TNytGGofk@EQhGG6OpTPP4rOkWAkWlrn3bdObfN658LHvheDxseoHacEnWVlJibMYatQVuFsuBS8tu1nQJTclx5AdWl6XvrDK2qkVQcQBN3DAlgmyYmYeJTaYg7m7aoTfkn1QI.MaJFoUxKxg1U03WgqYfeac8g"),
				PhoneNumber: &truePhoneNumber,
			},
			wantErrorMsg: errorDomain.NewError("Emailは320文字以内です").Error(),
		},
		"異常系: 電話番号が数字以外を含む": {
			username: trueName,
			options: &userDomain.Option{
				Email:       &trueEmail,
				PhoneNumber: common.StringPointer("0901234567a"),
			},
			wantErrorMsg: errorDomain.NewError("電話番号は数字のみで構成されています").Error(),
		},
		"異常系: 電話番号が11文字以上": {
			username: trueName,
			options: &userDomain.Option{
				Email:       &trueEmail,
				PhoneNumber: common.StringPointer("090123456789"),
			},
			wantErrorMsg: errorDomain.NewError("電話番号は11文字以内です").Error(),
		},
	}

	// テストケースの実行
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			user, err := userDomain.NewUser(tt.username, position, team, facility, department, area, policies, tt.options)

			if tt.wantErrorMsg == "" {
				assert.NotNil(t, user)
				assert.Equal(t, tt.username, user.Username)
				if tt.options.Email != nil {
					assert.Equal(t, *tt.options.Email, user.Email)
				}
				if tt.options.PhoneNumber != nil {
					assert.Equal(t, *tt.options.PhoneNumber, user.PhoneNumber)
				}
			} else {
				assert.Error(t, err)
				if err != nil {
					assert.Equal(t, tt.wantErrorMsg, err.Error())
				}
			}
		})
	}
}
