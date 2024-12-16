package schedule_type

import (
	"github.com/Fukuemon/go-pkg/ulid"
)

type ScheduleTypeEnum string

const (
	Normal ScheduleTypeEnum = "通常"
	Visit  ScheduleTypeEnum = "訪問"
)

type ScheduleType struct {
	ID   string `grom:"primaryKey"`
	Name ScheduleTypeEnum
}

func NewScheduleType(name ScheduleTypeEnum) *ScheduleType {
	return newScheduleType(ulid.NewULID(), name)
}

func newScheduleType(id string, name ScheduleTypeEnum) *ScheduleType {
	return &ScheduleType{
		ID:   id,
		Name: name,
	}
}
