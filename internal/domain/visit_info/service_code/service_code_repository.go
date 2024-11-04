package service_code

import "context"

type ServiceCodeRepository interface {
	Create(ctx context.Context, serviceCode *ServiceCode) error
	FindAll(ctx context.Context) ([]*ServiceCode, error)
	FindByID(ctx context.Context, id string) (*ServiceCode, error)
}
