package visit_category

import (
	"api-buddy/domain/common"

	"github.com/Fukuemon/go-pkg/ulid"
)

type VisitCategoryType string

const (
	NightShift VisitCategoryType = "夜勤"
	Emergency  VisitCategoryType = "緊急"
	Hospital   VisitCategoryType = "入院"
)

type VisitCategory struct {
	ID   string
	Name VisitCategoryType
	common.CommonModel
}

func NewVisitCategory(Name VisitCategoryType) (*VisitCategory, error) {
	return newVisitCategory(
		ulid.NewULID(),
		Name,
	)
}

func newVisitCategory(ID string, Name VisitCategoryType) (*VisitCategory, error) {
	visitCategory := &VisitCategory{
		ID:   ID,
		Name: Name,
	}

	common.InitializeCommonModel(&visitCategory.CommonModel)

	return visitCategory, nil
}
