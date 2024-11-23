package visit_category

import (
	"api-buddy/domain/common"

	"github.com/Fukuemon/go-pkg/ulid"
)

type VisitCategory struct {
	ID   string
	Name string
	common.CommonModel
}

func NewVisitCategory(Name string) (*VisitCategory, error) {
	return newVisitCategory(
		ulid.NewULID(),
		Name,
	)
}

func newVisitCategory(ID string, Name string) (*VisitCategory, error) {
	visitCategory := &VisitCategory{
		ID:   ID,
		Name: Name,
	}

	common.InitializeCommonModel(&visitCategory.CommonModel)

	return visitCategory, nil
}
