package schedule_cancel

import (
	"api-buddy/domain/common"

	"github.com/Fukuemon/go-pkg/ulid"
)

type ScheduleCancel struct {
	ID         string
	ScheduleID string
	Reason     string
	common.CommonModel
}

func NewScheduleCancel(schedule_id string, reason string) (*ScheduleCancel, error) {
	return newScheduleCancel(ulid.NewULID(), schedule_id, reason)
}

func newScheduleCancel(id string, schedule_id string, reason string) (*ScheduleCancel, error) {
	scheduleCancel := &ScheduleCancel{
		ID:         id,
		ScheduleID: schedule_id,
		Reason:     reason,
	}

	common.InitializeCommonModel(&scheduleCancel.CommonModel)
	return scheduleCancel, nil
}
