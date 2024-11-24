package schedule_type

type ScheduleTypeRepository interface {
	FindAll(facility_id string) ([]*ScheduleType, error)
	FindByID(id string) (*ScheduleType, error)
}
