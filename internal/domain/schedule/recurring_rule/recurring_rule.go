package recurring_rule

import (
	"api-buddy/domain/common"

	errorDomain "api-buddy/domain/error"

	"github.com/Fukuemon/go-pkg/ulid"
)

type FrequencyEnum string

const (
	FrequencyMonthly FrequencyEnum = "monthly"
	FrequencyWeekly  FrequencyEnum = "weekly"
)

type RecurringRule struct {
	ID          string `grom:"primaryKey"`
	Frequency   FrequencyEnum
	DaysOfWeek  int
	DayOfMonth  int
	WeekOfMonth int
	StartDate   common.Date
	EndDate     common.Date
	common.CommonModel
}

type RecurringRuleOption func(*RecurringRule) error

func WithDaysOfWeek(daysOfWeek int) RecurringRuleOption {
	return func(r *RecurringRule) error {
		r.DaysOfWeek = daysOfWeek
		return nil
	}
}

func WithDayOfMonth(dayOfMonth int) RecurringRuleOption {
	return func(r *RecurringRule) error {
		r.DayOfMonth = dayOfMonth
		return nil
	}
}

func WithWeekOfMonth(weekOfMonth int) RecurringRuleOption {
	return func(r *RecurringRule) error {
		r.WeekOfMonth = weekOfMonth
		return nil
	}
}

func WithEndDate(endDate common.Date) RecurringRuleOption {
	return func(r *RecurringRule) error {
		r.EndDate = endDate
		return nil
	}
}

func NewRecurringRule(
	frequency FrequencyEnum,
	startDate common.Date,
	options ...RecurringRuleOption,
) (*RecurringRule, error) {
	return newRecurringRule(
		ulid.NewULID(),
		frequency,
		startDate,
		options...,
	)
}

func newRecurringRule(
	id string,
	frequency FrequencyEnum,
	startDate common.Date,
	options ...RecurringRuleOption,
) (*RecurringRule, error) {
	recurring_rule := &RecurringRule{
		ID:        id,
		Frequency: frequency,
		StartDate: startDate,
	}

	for _, option := range options {
		if err := option(recurring_rule); err != nil {
			return nil, err
		}
	}
	// Todo: 繰り返しルールの制約
	//DaysOfWeek, DayOfMonth, WeekOfMonthのいずれかが設定されているか確認
	if recurring_rule.DaysOfWeek == 0 && recurring_rule.DayOfMonth == 0 && recurring_rule.WeekOfMonth == 0 {
		return nil, errorDomain.NewError("繰り返しルールが不正です")
	}

	common.InitializeCommonModel(&recurring_rule.CommonModel)
	return recurring_rule, nil
}
