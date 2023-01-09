package main

import (
	sdk "github.com/conduitio/conduit-connector-sdk"

	connectorname "github.com/conduitio/conduit-connector-connectorname"
)

func main() {
	sdk.Serve(connectorname.Connector)
}
