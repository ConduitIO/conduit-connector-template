package connectorname

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

// version is set during the build process with ldflags (see Makefile).
// Default version matches default from runtime/debug.
var version = "(devel)"

// Specification returns the connector's specification.
func Specification() sdk.Specification {
	return sdk.Specification{
		Name:        "connectorname",
		Summary:     "<describe your connector>",
		Description: "<describe your connector in detail>",
		Version:     version,
		Author:      "<your name>",
	}
}
