package department

import "time"

type DepartmentResponse struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	FacilityID string     `json:"facility_id"`
	CreateAt   time.Time  `json:"created_at"`
	UpdateAt   time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type FetchDepartmentsResponse struct {
	Departments []DepartmentResponse `json:"departments"`
}

type ErrorResponse struct {
	Error string `json:"message"`
}
