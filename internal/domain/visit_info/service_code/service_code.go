package service_code

import (
	"api-buddy/domain/common"

	"github.com/Fukuemon/go-pkg/ulid"
)

type ServiceCode struct {
	ID                    string
	Code                  string
	ServiceTimeRangeStart int
	ServiceTimeRangeEnd   int
	common.CommonModel
}

func NewServiceCode(
	Code string,
	ServiceTimeRangeStart int,
	ServiceTimeRangeEnd int,
) (*ServiceCode, error) {
	return newServiceCode(
		ulid.NewULID(),
		Code,
		ServiceTimeRangeStart,
		ServiceTimeRangeEnd,
	)
}

func newServiceCode(
	ID string,
	Code string,
	ServiceTimeRangeStart int,
	ServiceTimeRangeEnd int,
) (*ServiceCode, error) {
	serviceCode := &ServiceCode{
		ID:                    ID,
		Code:                  Code,
		ServiceTimeRangeStart: ServiceTimeRangeStart,
		ServiceTimeRangeEnd:   ServiceTimeRangeEnd,
	}

	common.InitializeCommonModel(&serviceCode.CommonModel)

	return serviceCode, nil
}
