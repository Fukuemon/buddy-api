package route

import (
	addressDomain "api-buddy/domain/address"

	"github.com/Fukuemon/go-pkg/query"
	"github.com/Fukuemon/go-pkg/ulid"
)

type Route struct {
	ID            string
	TravelTime    int
	AddressID     string
	Address       *addressDomain.Address `gorm:"foreignKey:AddressID"`
	DestinationID string
	Destination   *addressDomain.Address `gorm:"foreignKey:DestinationID"`
}

func NewRoute(
	TravelTime int,
	Address *addressDomain.Address,
	Destination *addressDomain.Address,
) (*Route, error) {
	return newRoute(
		ulid.NewULID(),
		TravelTime,
		Address,
		Destination,
	)
}

func newRoute(
	ID string,
	TravelTime int,
	Address *addressDomain.Address,
	Destination *addressDomain.Address,
) (*Route, error) {
	route := &Route{
		ID:            ID,
		TravelTime:    TravelTime,
		AddressID:     Address.ID,
		Address:       Address,
		DestinationID: Destination.ID,
		Destination:   Destination,
	}

	return route, nil
}

var RouteRelationMappings = map[string]query.RelationMapping{
	"address": {
		TableName:   "addresses",
		JoinKey:     "addresses.id = routes.address_id",
		FilterField: "addresses.name",
	},
	"destination": {
		TableName:   "addresses",
		JoinKey:     "addresses.id = routes.destination_id",
		FilterField: "addresses.name",
	},
}
