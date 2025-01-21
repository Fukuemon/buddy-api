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
	DayOfWeek   *int
	DayOfMonth  *int
	WeekOfMonth *int
	EndDate     *common.Date
	common.CommonModel
}

type RecurringRuleOption func(*RecurringRule) error

func WithDayOfWeek(daysOfWeek *int) RecurringRuleOption {
	return func(r *RecurringRule) error {
		r.DayOfWeek = daysOfWeek
		return nil
	}
}

func WithDayOfMonth(dayOfMonth *int) RecurringRuleOption {
	return func(r *RecurringRule) error {
		r.DayOfMonth = dayOfMonth
		return nil
	}
}

func WithWeekOfMonth(weekOfMonth *int) RecurringRuleOption {
	return func(r *RecurringRule) error {
		r.WeekOfMonth = weekOfMonth
		return nil
	}
}

func WithEndDate(endDate *common.Date) RecurringRuleOption {
	return func(r *RecurringRule) error {
		r.EndDate = endDate
		return nil
	}
}

func NewRecurringRule(
	frequency FrequencyEnum,
	options ...RecurringRuleOption,
) (*RecurringRule, error) {
	return newRecurringRule(
		ulid.NewULID(),
		frequency,
		options...,
	)
}

func newRecurringRule(
	id string,
	frequency FrequencyEnum,
	options ...RecurringRuleOption,
) (*RecurringRule, error) {
	recurring_rule := &RecurringRule{
		ID:          id,
		Frequency:   frequency,
		DayOfWeek:   nil,
		DayOfMonth:  nil,
		WeekOfMonth: nil,
		EndDate:     nil,
	}

	for _, option := range options {
		if err := option(recurring_rule); err != nil {
			return nil, err
		}
	}
	// Todo: 繰り返しルールの制約
	//DayOfWeek, DayOfMonth, WeekOfMonthがいずれも設定されていない場合は許容する
	if recurring_rule.DayOfWeek == nil && recurring_rule.DayOfMonth == nil && recurring_rule.WeekOfMonth == nil {
		common.InitializeCommonModel(&recurring_rule.CommonModel)
		return recurring_rule, nil
	}

	if recurring_rule.Frequency == FrequencyMonthly && recurring_rule.DayOfMonth == nil {
		err := errorDomain.NewError("月の日を指定してください")
		return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
	}

	if recurring_rule.Frequency == FrequencyWeekly && recurring_rule.DayOfWeek == nil {
		err := errorDomain.NewError("曜日を指定してください")
		return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
	}

	//DayOfWeek, DayOfMonth, WeekOfMonthのいずれかが重複していないか確認
	if recurring_rule.DayOfWeek != nil && recurring_rule.DayOfMonth != nil {
		err := errorDomain.NewError("曜日と月の日が重複しています")
		return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
	}

	if recurring_rule.WeekOfMonth != nil && recurring_rule.DayOfMonth != nil {
		err := errorDomain.NewError("週の日と月の日が重複しています")
		return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
	}

	if recurring_rule.DayOfWeek != nil && (*recurring_rule.DayOfWeek < 1 || *recurring_rule.DayOfWeek > 7) {
		err := errorDomain.NewError("曜日が不正です")
		return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
	}

	if recurring_rule.DayOfMonth != nil && (*recurring_rule.DayOfMonth < 1 || *recurring_rule.DayOfMonth > 31) {
		err := errorDomain.NewError("月の日が不正です")
		return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
	}

	if recurring_rule.WeekOfMonth != nil && (*recurring_rule.WeekOfMonth < 1 || *recurring_rule.WeekOfMonth > 5) {
		err := errorDomain.NewError("週の日が不正です")
		return nil, errorDomain.WrapError(errorDomain.InvalidInputErr, err)
	}

	common.InitializeCommonModel(&recurring_rule.CommonModel)
	return recurring_rule, nil
}
