package policy

import (
	"api-buddy/domain/common"

	"github.com/Fukuemon/go-pkg/ulid"
)

type Policy struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	common.CommonModel
}

func NewPolicy(name string) (*Policy, error) {
	return newPolicy(
		ulid.NewULID(),
		name,
	)
}

func newPolicy(ID string, name string) (*Policy, error) {
	policy := &Policy{
		ID:   ID,
		Name: name,
	}
	common.InitializeCommonModel(&policy.CommonModel)

	return policy, nil
}
