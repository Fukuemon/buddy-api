package user_test

// import (
// 	"api-buddy/domain/facility"
// 	"api-buddy/domain/facility/department"
// 	"api-buddy/domain/facility/position"
// 	"api-buddy/domain/facility/team"
// 	"api-buddy/domain/policy"
// 	userDomain "api-buddy/domain/user"
// 	"testing"

// 	"github.com/Fukuemon/go-pkg/ulid"
// )

// func TestNewUser(t *testing.T) {
// 	facilityID := ulid.NewULID()
// 	departmentID := ulid.NewULID()
// 	positionID := ulid.NewULID()
// 	teamID := ulid.NewULID()
// 	policyID1 := ulid.NewULID()
// 	policyID2 := ulid.NewULID()
// 	facility := &facility.Facility{
// 		ID:   facilityID,
// 		Name: "testFacility",
// 	}
// 	department := &department.Department{
// 		ID:         departmentID,
// 		FacilityID: facilityID,
// 		Name:       "testDepartment",
// 	}
// 	position := &position.Position{
// 		ID:         positionID,
// 		FacilityID: facilityID,
// 		Name:       "testPosition",
// 	}
// 	team := &team.Team{
// 		ID:         teamID,
// 		FacilityID: facilityID,
// 		Name:       "testTeam",
// 	}
// 	policies := []*policy.Policy{
// 		{
// 			ID:   policyID1,
// 			Name: "testPolicy1",
// 		},
// 		{
// 			ID:   policyID2,
// 			Name: "testPolicy2",
// 		},
// 	}

// 	type args struct {
// 		username    string
// 		email       string
// 		phoneNumber string
// 		facility    *facility.Facility
// 		department  *department.Department
// 		position    *position.Position
// 		team        *team.Team
// 		policies    []*policy.Policy
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *userDomain.User
// 		wantErr bool
// 	}{
// 		{
// 			name: "正常系",
// 			args: args{
// 				username:    "testUser",
// 				email:       "test@example.com",
// 				phoneNumber: "09012345678",
// 				facility:    facility,
// 				department:  department,
// 				position:    position,
// 				team:        team,
// 				policies:    policies,
// 			},
// 			want: &userDomain.User{
// 				ID:          ulid.NewULID(),
// 				Username:    "testUser",
// 				Email:       "test@example.com",
// 				PhoneNumber: "09012345678",
// 				FacilityID:  facilityID,
// 				Facility:    facility,
// 				DepartmentID: departmentID,
// 				Department:   department,
// 				PositionID:   positionID,
// 				Position:     position,
// 				TeamID:       teamID,
// 				Team:         team,
// 				Policies:     policies,
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "異常系：ユーザー名が128文字を超える場合",
// 			args: args{
// 				username:    "testUser123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
// 				email:       "test@example.com",
// 				phoneNumber: "09012345678",
// 				facility:    facility,
// 				department:  department,
// 				position:    position,
// 				team:        team,
// 				policies:    policies,
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 		{
// 			name: "異常系：Emailが320文字を超える場合",
// 			args: args{
// 				username:    "testUser",
// 				email:			 "test@examplep2P5WP5ZT1O4NCxo1TKqCu6NjMPNlf5gazKSJLe2cjb5bPWMgQ20ImL1UYP4Brv5Nkh5uHrqCuQbWkpKY9QqSzQNvk58dYdi0hO7zzeahEoV7VueO0XBixvbmBUdNnKDR5KW2NjtbwwRC2vWA3gVpFDdLU4JkdCMLwihuoneh0bgMt21rQuRJRGCCPkzcevWSaVAIP2Jk0119qzWSBPtXvvKvw5P5gs1qpiLGH5raKOgbZQ3V5VflKiDq3gN3RjT9wZrkSBUGA0g8useWbJ3p2w8F16rU55Vp4rpmKL11IMBffQD.com",
// 				phoneNumber: "09012345678",
// 				facility:    facility,
// 				department:  department,
// 				position: 	position,
// 				team:        team,
// 				policies:    policies,
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 		{
// 			name: "異常系：PhoneNumberが11文字を超える場合",
// 			args: args{
// 				username:    "testUser",
// 				phoneNumber: "090123456789",
// 				facility:    facility,
// 				department:  department,
// 				position:    position,
// 				team:        team,
// 				policies:    policies,
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := userDomain.NewUser(tt.args.username, tt.args.email, tt.args.phoneNumber, tt.args.facility, tt.args.department, tt.args.position, tt.args.team, tt.args.policies)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if tt.wantErr {
// 				return
// 			}
// 			if got.ID == "" {
// 				t.Errorf("NewUser() got.ID = %v, want not empty", got.ID)
// 			}
// 			if got.Username != tt.want.Username {
// 				t.Errorf("NewUser() got.Username = %v, want %v", got.Username, tt.want.Username)
// 			}
// 			if got.Email != tt.want.Email {
// 				t.Errorf("NewUser() got.Email = %v, want %v", got.Email, tt.want.Email)
// 			}
// 			if got.PhoneNumber != tt.want.PhoneNumber {
// 				t.Errorf("NewUser() got.PhoneNumber = %v, want %v", got.PhoneNumber, tt.want.PhoneNumber)
// 			}
// 			if got.FacilityID != tt.want.FacilityID {
// 				t.Errorf("NewUser() got.FacilityID = %v, want %v", got.FacilityID, tt.want.FacilityID)
// 			}
// 			if got.Facility != tt.want.Facility {
// 				t.Errorf("NewUser() got.Facility = %v, want %v", got.Facility, tt.want.Facility)
// 			}
// 			if got.DepartmentID != tt.want.DepartmentID {
// 				t.Errorf("NewUser() got.DepartmentID = %v, want %v", got.DepartmentID, tt.want.DepartmentID)
// 			}
// 			if got.Department != tt.want.Department {
// 				t.Errorf("NewUser() got.Department = %v, want %v", got.Department, tt.want.Department)
// 			}
// 			if got.PositionID != tt.want.PositionID {
// 				t.Errorf("NewUser() got.PositionID = %v, want %v", got.PositionID, tt.want.PositionID)
// 			}
// 			if got.Position != tt.want.Position {
// 				t.Errorf("NewUser() got.Position = %v, want %v", got.Position, tt.want
// .Position)
// 			}
// 			if got.TeamID != tt.want.TeamID {
// 				t.Errorf("NewUser() got.TeamID = %v, want %v", got.TeamID, tt.want.TeamID)
// 			}
