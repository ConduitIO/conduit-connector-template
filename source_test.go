package connectorname_test

import (
	"context"
	"testing"

	connectorname "github.com/conduitio/conduit-connector-connectorname"
)

func TestTeardownSource_NoOpen(t *testing.T) {
	con := connectorname.NewSource()
	err := con.Teardown(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
