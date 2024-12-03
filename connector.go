package connectorname

import (
	_ "embed"

	sdk "github.com/conduitio/conduit-connector-sdk"
)

//go:embed connector.yaml
var specs string

var Connector = sdk.Connector{
	NewSpecification: sdk.YAMLSpecification(specs),
	NewSource:        NewSource,
	NewDestination:   NewDestination,
}
