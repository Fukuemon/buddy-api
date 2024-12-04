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
	ServiceCode     *serviceCodeDomain.ServiceCode       `gorm:"foreignKey:ServiceCodeID"`
	VisitCategories []*visitCategoryDomain.VisitCategory `gorm:"many2many:visit_info_visit_categories"`
	common.CommonModel
}

// Functional option type
type VisitInfoOption func(*VisitInfo)

// WithCompanion sets the companion for the VisitInfo
func WithCompanion(companion *userDomain.User) VisitInfoOption {
	return func(vi *VisitInfo) {
		if companion != nil {
			vi.CompanionID = companion.ID
			vi.Companion = companion
		}
	}
}

// WithRoute sets the route for the VisitInfo
func WithRoute(route *routeDomain.Route) VisitInfoOption {
	return func(vi *VisitInfo) {
		if route != nil {
			vi.RouteID = route.ID
			vi.Route = route
		}
	}
}

// WithVisitCategories sets the visit categories for the VisitInfo
func WithVisitCategories(categories []*visitCategoryDomain.VisitCategory) VisitInfoOption {
	return func(vi *VisitInfo) {
		if categories != nil {
			vi.VisitCategories = categories
		}
	}
}

// NewVisitInfo initializes a new VisitInfo with required fields and options
func NewVisitInfo(
	patient *patientDomain.Patient,
	assignedStaff *userDomain.User,
	serviceCode *serviceCodeDomain.ServiceCode,
	options ...VisitInfoOption,
) (*VisitInfo, error) {
	visitInfo := &VisitInfo{
		ID:              ulid.NewULID(),
		PatientID:       patient.ID,
		Patient:         patient,
		AssignedStaffID: assignedStaff.ID,
		AssignedStaff:   assignedStaff,
		ServiceCodeID:   serviceCode.ID,
		ServiceCode:     serviceCode,
	}

	// Apply options
	for _, option := range options {
		option(visitInfo)
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
