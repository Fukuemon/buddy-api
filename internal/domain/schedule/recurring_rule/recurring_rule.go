package recurring_rule

import (
	"api-buddy/domain/common"
	"time"

	"github.com/Fukuemon/go-pkg/ulid"
)

type RecurringRule struct {
	ID          string `grom:"primaryKey"`
	Frequency   string
	DaysOfWeek  int
	DayOfMonth  int
	WeekOfMonth int
	StartDate   time.Time
	EndDate     time.Time
	common.CommonModel
}

func NewRecurringRule(
	frequency string,
	daysOfWeek int,
	dayOfMonth int,
	weekOfMonth int,
	startDate time.Time,
	endDate time.Time,
) (*RecurringRule, error) {
	return newRecurringRule(
		ulid.NewULID(),
		frequency, daysOfWeek,
		dayOfMonth, weekOfMonth,
		startDate,
		endDate,
	)
}

func newRecurringRule(
	id string,
	frequency string,
	daysOfWeek int,
	dayOfMonth int,
	weekOfMonth int,
	startDate time.Time, endDate time.Time) (*RecurringRule, error) {
	recurring_rule := &RecurringRule{
		ID:          id,
		Frequency:   frequency,
		DaysOfWeek:  daysOfWeek,
		DayOfMonth:  dayOfMonth,
		WeekOfMonth: weekOfMonth,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	common.InitializeCommonModel(&recurring_rule.CommonModel)
	return recurring_rule, nil
}
