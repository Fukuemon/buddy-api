package patient

import patientDomain "api-buddy/domain/patient"

type CreatePatientResponse struct {
	ID              string                         `json:"id"`
	Name            string                         `json:"name"`
	PreferredTime   *patientDomain.PreferredTime   `json:"preferred_time"`
	PreferredGender *patientDomain.PreferredGender `json:"preferred_gender"`
	ServiceCode     string                         `json:"service_code"`
	Address         string                         `json:"address"`
	Area            string                         `json:"area"`
	AssignedStaff   string                         `json:"assigned_staff"`
	Facility        string                         `json:"facility"`
}

type PatientResponse struct {
	ID              string                         `json:"id"`
	Name            string                         `json:"name"`
	PreferredTime   *patientDomain.PreferredTime   `json:"preferred_time"`
	PreferredGender *patientDomain.PreferredGender `json:"preferred_gender"`
	AssignedStaff   string                         `json:"assigned_staff"`
	Address         string                         `json:"address"`
	Area            string                         `json:"area"`
}

type PatientListResponse []PatientResponse
