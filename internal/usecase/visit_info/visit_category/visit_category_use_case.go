package visit_category

import (
	"api-buddy/domain/visit_info/visit_category"
	"context"
)

type FetchVisitCategoriesUseCase struct {
	visitCategoryRepository visit_category.VisitCategoryRepository
}

func NewFetchVisitCategoriesUseCase(visitCategoryRepository visit_category.VisitCategoryRepository) *FetchVisitCategoriesUseCase {
	return &FetchVisitCategoriesUseCase{
		visitCategoryRepository: visitCategoryRepository,
	}
}

type FetchVisitCategoriesUseCaseOutputDto struct {
	ID   string
	Name visit_category.VisitCategoryType
}

func (uc *FetchVisitCategoriesUseCase) Run(ctx context.Context) ([]FetchVisitCategoriesUseCaseOutputDto, error) {
	visitCategories, err := uc.visitCategoryRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputDto []FetchVisitCategoriesUseCaseOutputDto
	for _, visitCategory := range visitCategories {
		outputDto = append(outputDto, FetchVisitCategoriesUseCaseOutputDto{
			ID:   visitCategory.ID,
			Name: visitCategory.Name,
		})
	}

	return outputDto, nil
}
