package recurring_schedule

import (
	"api-buddy/domain/common"
	facilityDomain "api-buddy/domain/facility"
	recurringRuleDomain "api-buddy/domain/schedule/recurring_rule"
	recurringScheduleDomain "api-buddy/domain/schedule/recurring_schedule"
	scheduleTypeDomain "api-buddy/domain/schedule/schedule_type"
	userDomain "api-buddy/domain/user"
	visitInfoDomain "api-buddy/domain/visit_info"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type CreateRecurringScheduleUseCase struct {
	recurringRuleRepository     recurringRuleDomain.RecurringRuleRepository
	facilityRepository          facilityDomain.FacilityRepository
	scheduleTypeRepository      scheduleTypeDomain.ScheduleTypeRepository
	userRepository              userDomain.UserRepository
	recurringScheduleRepository recurringScheduleDomain.RecurringScheduleRepository
	visitInfoService            *visitInfoDomain.VisitInfoService
}

func NewCreateRecurringScheduleUseCase(
	recurringRuleRepository recurringRuleDomain.RecurringRuleRepository,
	facilityRepository facilityDomain.FacilityRepository,
	scheduleTypeRepository scheduleTypeDomain.ScheduleTypeRepository,
	userRepository userDomain.UserRepository,
	recurringScheduleRepository recurringScheduleDomain.RecurringScheduleRepository,
	visitInfoService *visitInfoDomain.VisitInfoService,
) *CreateRecurringScheduleUseCase {
	return &CreateRecurringScheduleUseCase{
		recurringRuleRepository:     recurringRuleRepository,
		facilityRepository:          facilityRepository,
		scheduleTypeRepository:      scheduleTypeRepository,
		userRepository:              userRepository,
		recurringScheduleRepository: recurringScheduleRepository,
		visitInfoService:            visitInfoService,
	}
}

type CreateRecurringScheduleUseCaseInputDto struct {
	ScheduleTypeID string
	RecurringRule  RecurringRuleModel
	Date           common.Date
	StartTime      common.Time
	EndTime        common.Time
	StaffID        string
	FacilityID     string
	Title          *string
	Description    *string
	VisitInfo      *visitInfoDomain.VisitInfoModel
}

type RecurringRuleModel struct {
	Frequency   recurringRuleDomain.FrequencyEnum
	DayOfWeek   *int
	DayOfMonth  *int
	WeekOfMonth *int
	StartDate   common.Date
	EndDate     *common.Date
}

func (uc *CreateRecurringScheduleUseCase) Run(ctx context.Context, input CreateRecurringScheduleUseCaseInputDto) (*recurringScheduleDomain.RecurringSchedule, error) {
	var recurringSchedule *recurringScheduleDomain.RecurringSchedule

	err := db.WithTransaction(ctx, func(tx *gorm.DB) error {
		recurringScheduleOptions := []recurringScheduleDomain.RecurringScheduleOption{}

		// 予定の種類を取得
		scheduleType, err := uc.scheduleTypeRepository.FindByID(ctx, input.ScheduleTypeID)
		if err != nil {
			return err
		}

		// スタッフを取得
		staff, err := uc.userRepository.FindByID(ctx, input.StaffID)
		if err != nil {
			return err
		}

		// 施設を取得
		facility, err := uc.facilityRepository.FindByID(ctx, input.FacilityID)
		if err != nil {
			return err
		}

		recurringScheduleRuleOptions := []recurringRuleDomain.RecurringRuleOption{}

		// 曜日が指定されている場合
		if input.RecurringRule.DayOfWeek != nil {
			recurringScheduleRuleOptions = append(recurringScheduleRuleOptions, recurringRuleDomain.WithDayOfWeek(*input.RecurringRule.DayOfWeek))
		}

		// 月の日が指定されている場合
		if input.RecurringRule.DayOfMonth != nil {
			recurringScheduleRuleOptions = append(recurringScheduleRuleOptions, recurringRuleDomain.WithDayOfMonth(*input.RecurringRule.DayOfMonth))
		}

		// 週の日が指定されている場合
		if input.RecurringRule.WeekOfMonth != nil {
			recurringScheduleRuleOptions = append(recurringScheduleRuleOptions, recurringRuleDomain.WithWeekOfMonth(*input.RecurringRule.WeekOfMonth))
		}

		// 終了日が指定されている場合
		if input.RecurringRule.EndDate != nil {
			recurringScheduleRuleOptions = append(recurringScheduleRuleOptions, recurringRuleDomain.WithEndDate(*input.RecurringRule.EndDate))
		}

		// リクエストされた繰り返しルールを作成
		recurringRule, err := recurringRuleDomain.NewRecurringRule(
			input.RecurringRule.Frequency,
			input.RecurringRule.StartDate,
			recurringScheduleRuleOptions...,
		)
		if err != nil {
			return err
		}

		// 繰り返しルールの登録
		if err := uc.recurringRuleRepository.Create(ctx, tx, recurringRule); err != nil {
			return err
		}

		// 通常予定の場合
		if scheduleType.Name == scheduleTypeDomain.Normal && input.Title != nil {
			recurringScheduleOptions = append(recurringScheduleOptions, recurringScheduleDomain.WithTitle(*input.Title))
		}

		// 訪問情報がある場合
		if scheduleType.Name == scheduleTypeDomain.Visit && input.VisitInfo != nil {
			visitInfo, err := uc.visitInfoService.CreateVisitInfo(ctx, tx, *input.VisitInfo)
			if err != nil {
				return err
			}

			recurringScheduleOptions = append(recurringScheduleOptions, recurringScheduleDomain.WithVisitInfo(visitInfo))
		}

		// 予定の説明がある場合
		if input.Description != nil {
			recurringScheduleOptions = append(recurringScheduleOptions, recurringScheduleDomain.WithDescription(*input.Description))
		}

		// 繰り返し予定の作成
		recurringSchedule, err = recurringScheduleDomain.NewRecurringSchedule(
			recurringRule,
			scheduleType,
			input.Date,
			input.StartTime,
			input.EndTime,
			staff,
			facility,
			recurringScheduleOptions...,
		)
		if err != nil {
			return err
		}

		// 繰り返し予定を登録
		if err := uc.recurringScheduleRepository.Create(ctx, tx, recurringSchedule); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return recurringSchedule, nil
}
