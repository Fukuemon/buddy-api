package visit_info

import (
	"api-buddy/domain/common"
	patientDomain "api-buddy/domain/patient"
	userDomain "api-buddy/domain/user"
	routeDomain "api-buddy/domain/visit_info/route"
	serviceCodeDomain "api-buddy/domain/visit_info/service_code"
	visitCategoryDomain "api-buddy/domain/visit_info/visit_category"

	"github.com/Fukuemon/go-pkg/query"
	"github.com/Fukuemon/go-pkg/ulid"
)

type Option struct {
	Companion       *userDomain.User
	CompanionID     *string
	VisitCategory   *visitCategoryDomain.VisitCategory
	VisitCategoryID *string
}

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
	VisitCategoryID string
	VisitCategory   *visitCategoryDomain.VisitCategory `gorm:"foreignKey:VisitCategoryID"`
	common.CommonModel
}

func NewVisitInfo(
	patient *patientDomain.Patient,
	assignedStaff *userDomain.User,
	route *routeDomain.Route,
	serviceCode *serviceCodeDomain.ServiceCode,
	options *Option,
) (*VisitInfo, error) {
	return newVisitInfo(
		ulid.NewULID(),
		patient,
		assignedStaff,
		route,
		serviceCode,
		options,
	)
}

func newVisitInfo(
	ID string,
	patient *patientDomain.Patient,
	assignedStaff *userDomain.User,
	route *routeDomain.Route,
	serviceCode *serviceCodeDomain.ServiceCode,
	options *Option,
) (*VisitInfo, error) {
	visitInfo := &VisitInfo{
		ID:              ID,
		PatientID:       patient.ID,
		Patient:         patient,
		AssignedStaffID: assignedStaff.ID,
		AssignedStaff:   assignedStaff,
		RouteID:         route.ID,
		Route:           route,
		ServiceCodeID:   serviceCode.ID,
		ServiceCode:     serviceCode,
	}

	if options != nil {
		if options.Companion != nil {
			visitInfo.CompanionID = options.Companion.ID
			visitInfo.Companion = options.Companion
		}
		if options.VisitCategory != nil {
			visitInfo.VisitCategoryID = options.VisitCategory.ID
			visitInfo.VisitCategory = options.VisitCategory
		}
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
	"visit_category": {
		TableName:   "visit_categories",
		JoinKey:     "visit_categories.id = visit_infos.visit_category_id",
		FilterField: "visit_categories.name",
	},
}
