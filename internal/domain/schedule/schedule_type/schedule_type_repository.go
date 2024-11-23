package schedule_type

type ScheduleTypeRepository interface {
	FindByFacilityID(facility_id string) ([]*ScheduleType, error)
	FindByID(id string) (*ScheduleType, error)
}
