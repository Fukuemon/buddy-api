package route

import (
	addressDomain "api-buddy/domain/address"
	"context"

	"gorm.io/gorm"
)

type RouteService struct {
	routeRepository   RouteRepository
	addressRepository addressDomain.AddressRepository
}

func NewRouteService(
	routeRepository RouteRepository,
	addressRepository addressDomain.AddressRepository,
) *RouteService {
	return &RouteService{
		routeRepository:   routeRepository,
		addressRepository: addressRepository,
	}
}

type RouteModel struct {
	ID            string
	TravelTime    int
	FromAddressID string
	DestinationID string
}

func (s *RouteService) CreateRoute(
	ctx context.Context,
	tx *gorm.DB,
	routeModel *RouteModel,
) (*Route, error) {

	fromAddress, err := s.addressRepository.FindByID(ctx, routeModel.FromAddressID)
	if err != nil {
		return nil, err
	}

	toAddress, err := s.addressRepository.FindByID(ctx, routeModel.DestinationID)
	if err != nil {
		return nil, err
	}

	// 新規ルートの作成
	route, err := NewRoute(routeModel.TravelTime, fromAddress, toAddress)
	if err != nil {
		return nil, err
	}

	err = s.routeRepository.Create(ctx, tx, route)
	if err != nil {
		return nil, err
	}

	return route, nil
}
