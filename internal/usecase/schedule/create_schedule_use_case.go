package schedule

import (
	"api-buddy/domain/common"
	facilityDomain "api-buddy/domain/facility"
	scheduleDomain "api-buddy/domain/schedule"
	recurringScheduleDomain "api-buddy/domain/schedule/recurring_schedule"
	scheduleTypeDomain "api-buddy/domain/schedule/schedule_type"
	userDomain "api-buddy/domain/user"
	visitInfoDomain "api-buddy/domain/visit_info"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type CreateScheduleUseCase struct {
	scheduleRepository          scheduleDomain.ScheduleRepository
	facilityRepository          facilityDomain.FacilityRepository
	scheduleTypeRepository      scheduleTypeDomain.ScheduleTypeRepository
	userRepository              userDomain.UserRepository
	recurringScheduleRepository recurringScheduleDomain.RecurringScheduleRepository
	visitInfoService            *visitInfoDomain.VisitInfoService
}

func NewCreateScheduleUseCase(
	scheduleRepository scheduleDomain.ScheduleRepository,
	facilityRepository facilityDomain.FacilityRepository,
	scheduleTypeRepository scheduleTypeDomain.ScheduleTypeRepository,
	userRepository userDomain.UserRepository,
	recurringScheduleRepository recurringScheduleDomain.RecurringScheduleRepository,
	visitInfoService *visitInfoDomain.VisitInfoService,
) *CreateScheduleUseCase {
	return &CreateScheduleUseCase{
		scheduleRepository:          scheduleRepository,
		facilityRepository:          facilityRepository,
		scheduleTypeRepository:      scheduleTypeRepository,
		userRepository:              userRepository,
		recurringScheduleRepository: recurringScheduleRepository,
		visitInfoService:            visitInfoService,
	}
}

type CreateUseCaseInputDto struct {
	ScheduleTypeID      string
	Date                common.Date
	StartTime           common.Time
	EndTime             common.Time
	StaffID             string
	FacilityID          string
	Title               *string
	Description         *string
	VisitInfo           *visitInfoDomain.VisitInfoModel
	RecurringScheduleID *string
}

func (uc *CreateScheduleUseCase) Run(ctx context.Context, input CreateUseCaseInputDto) (*scheduleDomain.Schedule, error) {
	var schedule *scheduleDomain.Schedule

	err := db.WithTransaction(ctx, func(tx *gorm.DB) error {
		// スケジュールオプションの初期化
		scheduleOptions := []scheduleDomain.ScheduleOption{}

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

		// 通常予定の場合
		if scheduleType.Name == scheduleTypeDomain.Normal && input.Title != nil {
			scheduleOptions = append(scheduleOptions, scheduleDomain.WithTitle(input.Title))
		}

		// 訪問情報がある場合
		if scheduleType.Name == scheduleTypeDomain.Visit && input.VisitInfo != nil {
			visitInfo, err := uc.visitInfoService.CreateVisitInfo(ctx, tx, *input.VisitInfo)
			if err != nil {
				return err
			}
			scheduleOptions = append(scheduleOptions, scheduleDomain.WithVisitInfo(visitInfo))
		}

		// 予定の説明がある場合
		if input.Description != nil {
			scheduleOptions = append(scheduleOptions, scheduleDomain.WithDescription(input.Description))
		}

		// 繰り返し予定の変更予定の場合
		if input.RecurringScheduleID != nil {
			recurringSchedule, err := uc.recurringScheduleRepository.FindByID(ctx, *input.RecurringScheduleID)
			if err != nil {
				return err
			}
			scheduleOptions = append(scheduleOptions, scheduleDomain.WithRecurringSchedule(recurringSchedule))
		}

		// 予定のインスタンスを生成
		schedule, err = scheduleDomain.NewSchedule(scheduleType, input.Date, input.StartTime, input.EndTime, staff, facility, scheduleOptions...)
		if err != nil {
			return err
		}

		// 予定を登録
		if err := uc.scheduleRepository.Create(ctx, tx, schedule); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return schedule, nil
}
