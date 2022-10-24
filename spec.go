package connectorName

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

// Specification returns the connector's specification.
func Specification() sdk.Specification {
	return sdk.Specification{
		Name:        "connectorName",
		Summary:     "<describe your connector>",
		Description: "<describe your connector in detail>",
		Version:     "v0.1.0",
		Author:      "<your name>",
	}
}
