package schedule_type

import (
	"api-buddy/domain/schedule/schedule_type"
	"context"
)

type FetchScheduleTypesUseCase struct {
	scheduleTypeRepository schedule_type.ScheduleTypeRepository
}

func NewFetchScheduleTypesUseCase(scheduleTypeRepository schedule_type.ScheduleTypeRepository) *FetchScheduleTypesUseCase {
	return &FetchScheduleTypesUseCase{
		scheduleTypeRepository: scheduleTypeRepository,
	}
}

type FetchScheduleTypesUseCaseOutputDto struct {
	ID   string
	Name schedule_type.ScheduleTypeEnum
}

func (uc *FetchScheduleTypesUseCase) Run(ctx context.Context, facilityId string) ([]FetchScheduleTypesUseCaseOutputDto, error) {
	scheduleTypes, err := uc.scheduleTypeRepository.FindAll(ctx, facilityId)
	if err != nil {
		return nil, err
	}

	var outputDto []FetchScheduleTypesUseCaseOutputDto
	for _, scheduleType := range scheduleTypes {
		outputDto = append(outputDto, FetchScheduleTypesUseCaseOutputDto{
			ID:   scheduleType.ID,
			Name: scheduleType.Name,
		})
	}

	return outputDto, nil
}
