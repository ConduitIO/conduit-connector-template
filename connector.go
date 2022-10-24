package connectorName

import sdk "github.com/conduitio/conduit-connector-sdk"

var Connector = sdk.Connector{
	NewSpecification: Specification,
	NewSource:        NewSource,
	NewDestination:   NewDestination,
}
