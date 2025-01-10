package patient

import patientDomain "api-buddy/domain/patient"

type CreatePatientRequest struct {
	Name            string                         `json:"name" validate:"required"`
	PreferredTime   *patientDomain.PreferredTime   `json:"preferred_time" validate:"required"`
	PreferredGender *patientDomain.PreferredGender `json:"preferred_gender" validate:"required"`
	ServiceCodeID   string                         `json:"service_code_id" validate:"required,ulid"`
	AddressID       string                         `json:"address_id" validate:"required,ulid"`
	AreaID          string                         `json:"area_id" validate:"required,ulid"`
	AssignedStaffID string                         `json:"assigned_staff_id" validate:"required,ulid"`
}
