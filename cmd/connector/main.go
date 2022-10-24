package main

import (
	sdk "github.com/conduitio/conduit-connector-sdk"

	connectorName "github.com/conduitio/conduit-connector-connectorName"
)

func main() {
	sdk.Serve(connectorName.Connector)
}
