package patient

import (
	addressDomain "api-buddy/domain/address"
	"api-buddy/domain/common"
	facilityDomain "api-buddy/domain/facility"
	areaDomain "api-buddy/domain/facility/area"
	userDomain "api-buddy/domain/user"
	serviceCodeDomain "api-buddy/domain/visit_info/service_code"

	"github.com/Fukuemon/go-pkg/query"
	"github.com/Fukuemon/go-pkg/ulid"
)

type PreferredGender string
type PreferredTime string

func (p *PreferredGender) String() string {
	return string(*p)
}

func (p *PreferredTime) IsValid() bool {
	if p == nil {
		return false
	}
	switch *p {
	case AM, PM:
		return true
	}
	return false
}

func (p *PreferredTime) String() string {
	return string(*p)
}

func (p *PreferredGender) IsValid() bool {
	if p == nil {
		return false
	}
	switch *p {
	case Man, Woman:
		return true
	}
	return false
}

const (
	Man   PreferredGender = "男性"
	Woman PreferredGender = "女性"
	PM    PreferredTime   = "午後"
	AM    PreferredTime   = "午前"
)

type Patient struct {
	ID              string
	Name            string
	PreferredTime   *PreferredTime
	PreferredGender *PreferredGender
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

func NewPatient(
	name string, preferredTime *PreferredTime, preferredGender *PreferredGender, serviceCode *serviceCodeDomain.ServiceCode, address *addressDomain.Address, area *areaDomain.Area, assignedStaff *userDomain.User, facility *facilityDomain.Facility,
) (*Patient, error) {
	return newPatient(
		ulid.NewULID(),
		name,
		preferredTime,
		preferredGender,
		serviceCode,
		address,
		area,
		assignedStaff,
		facility,
	), nil
}

func newPatient(
	id string, name string, preferredTime *PreferredTime, preferredGender *PreferredGender, serviceCode *serviceCodeDomain.ServiceCode, address *addressDomain.Address, area *areaDomain.Area, assignedStaff *userDomain.User, facility *facilityDomain.Facility,
) *Patient {
	return &Patient{
		ID:              id,
		Name:            name,
		PreferredTime:   preferredTime,
		PreferredGender: preferredGender,
		ServiceCodeID:   serviceCode.ID,
		ServiceCode:     serviceCode,
		AddressID:       address.ID,
		Address:         address,
		AreaID:          area.ID,
		Area:            area,
		AssignedStaffID: assignedStaff.ID,
		AssignedStaff:   assignedStaff,
		FacilityID:      facility.ID,
		Facility:        facility,
	}
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
		FilterField: "addresses.zip_code",
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
