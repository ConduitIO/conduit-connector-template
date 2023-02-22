package connectorname

import sdk "github.com/conduitio/conduit-connector-sdk"

// Connector combines all constructors for each plugin in one struct.
var Connector = sdk.Connector{
	NewSpecification: Specification,
	NewSource:        NewSource,
	NewDestination:   NewDestination,
}
