package patient

import (
	addressDomain "api-buddy/domain/address"
	"api-buddy/domain/common"
	facilityDomain "api-buddy/domain/facility"
	areaDomain "api-buddy/domain/facility/area"
	userDomain "api-buddy/domain/user"
	serviceCodeDomain "api-buddy/domain/visit_info/service_code"

	"github.com/Fukuemon/go-pkg/query"
)

type Patient struct {
	ID              string
	Name            string
	PreferredTime   string
	PreferredGender string
	ServiceCodeID   string
	ServiceCode     *serviceCodeDomain.ServiceCode `gorm:"foreignKey:ServiceCodeID"`
	AddressID       string
	Address         *addressDomain.Address `gorm:"foreignKey:AddressID"`
	AreaID          string
	Area            *areaDomain.Area `gorm:"foreignKey:AreaID"`
	AssignedStaffID string
	AssignedStaff   *userDomain.User `gorm:"foreignKey:AssignedStaffID"`
	FacilityID      string
	Facility        *facilityDomain.Facility `gorm:"foreignKey:FacilityID"`
	common.CommonModel
}

var PatientRelationMappings = map[string]query.RelationMapping{
	"service_code": {
		TableName:   "service_codes",
		JoinKey:     "service_codes.id = patients.service_code_id",
		FilterField: "service_codes.name",
	},
	"address": {
		TableName:   "addresses",
		JoinKey:     "addresses.id = patients.address_id",
		FilterField: "addresses.name",
	},
	"area": {
		TableName:   "areas",
		JoinKey:     "areas.id = patients.area_id",
		FilterField: "areas.name",
	},
	"assigned_staff": {
		TableName:   "users",
		JoinKey:     "users.id = patients.assigned_staff_id",
		FilterField: "users.username",
	},
}
