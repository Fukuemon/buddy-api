package visit_info

import (
	"api-buddy/domain/common"
	patientDomain "api-buddy/domain/patient"
	userDomain "api-buddy/domain/user"
	routeDomain "api-buddy/domain/visit_info/route"
	serviceCodeDomain "api-buddy/domain/visit_info/service_code"

	"github.com/Fukuemon/go-pkg/query"
	"github.com/Fukuemon/go-pkg/ulid"
)

type VisitInfo struct {
	ID              string
	PatientID       string
	Patient         *patientDomain.Patient `gorm:"foreignKey:PatientID"`
	AssignedStaffID string
	AssignedStaff   *userDomain.User `gorm:"foreignKey:AssignedStaffID"`
	CompanionID     string
	Companion       *userDomain.User `gorm:"foreignKey:CompanionID"`
	RouteID         string
	Route           *routeDomain.Route `gorm:"foreignKey:RouteID"`
	ServiceCodeID   string
	ServiceCode     *serviceCodeDomain.ServiceCode `gorm:"foreignKey:ServiceCodeID"`
	common.CommonModel
}

func NewVisitInfo(
	patient *patientDomain.Patient,
	assignedStaff *userDomain.User,
	companion *userDomain.User,
	route *routeDomain.Route,
	serviceCode *serviceCodeDomain.ServiceCode,
) (*VisitInfo, error) {
	return newVisitInfo(
		ulid.NewULID(),
		patient,
		assignedStaff,
		companion,
		route,
		serviceCode,
	)
}

func newVisitInfo(
	ID string,
	patient *patientDomain.Patient,
	assignedStaff *userDomain.User,
	companion *userDomain.User,
	route *routeDomain.Route,
	serviceCode *serviceCodeDomain.ServiceCode,
) (*VisitInfo, error) {
	visitInfo := &VisitInfo{
		ID:              ID,
		PatientID:       patient.ID,
		Patient:         patient,
		AssignedStaffID: assignedStaff.ID,
		AssignedStaff:   assignedStaff,
		CompanionID:     companion.ID,
		Companion:       companion,
		RouteID:         route.ID,
		Route:           route,
		ServiceCodeID:   serviceCode.ID,
		ServiceCode:     serviceCode,
	}

	common.InitializeCommonModel(&visitInfo.CommonModel)

	return visitInfo, nil
}

var VisitInfoRelationMappings = map[string]query.RelationMapping{
	"patient": {
		TableName:   "patients",
		JoinKey:     "patients.id = visit_infos.patient_id",
		FilterField: "patients.name",
	},
	"assigned_staff": {
		TableName:   "users",
		JoinKey:     "users.id = visit_infos.assigned_staff_id",
		FilterField: "users.username",
	},
	"companion": {
		TableName:   "users",
		JoinKey:     "users.id = visit_infos.companion_staff_id",
		FilterField: "users.username",
	},
	"route": {
		TableName:   "routes",
		JoinKey:     "routes.id = visit_infos.route_id",
		FilterField: "routes.name",
	},
	"service_code": {
		TableName:   "service_codes",
		JoinKey:     "service_codes.id = visit_infos.service_code_id",
		FilterField: "service_codes.code",
	},
}
