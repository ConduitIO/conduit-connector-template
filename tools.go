//go:build tools

package connectorname

// Keeps the dependencies of the connector spec generator in go.mod and go.sum.
// "make generate" runs "conn-sdk-cli specgen", which builds a small program
// inside this module that imports the specgen package; without this import
// "go mod tidy" drops the modules that program needs (for example
// golang.org/x/exp) and "make generate" fails with a missing go.sum entry.
import _ "github.com/conduitio/conduit-connector-sdk/conn-sdk-cli/specgen"
