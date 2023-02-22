package connectorname_test

import (
	"context"
	"testing"

	connectorname "github.com/conduitio/conduit-connector-connectorname"
)

func TestTeardown_NoOpen(t *testing.T) {
	con := connectorname.NewDestination()
	err := con.Teardown(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
