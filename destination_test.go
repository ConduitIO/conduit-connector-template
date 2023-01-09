package connectorname_test

import (
	"context"
	"strings"
	"testing"

	connectorname "github.com/conduitio/conduit-connector-connectorname"
)

func TestConfigureDestination_FailsWhenConfigEmpty(t *testing.T) {
	con := connectorname.Destination{}
	err := con.Configure(context.Background(), make(map[string]string))
	if err == nil {
		t.Error("expected error for missing config params")
	}

	if strings.HasPrefix(err.Error(), "config is invalid:") {
		t.Errorf("expected error to be about missing config, got %v", err)
	}
}

func TestTeardown_NoOpen(t *testing.T) {
	con := connectorname.NewDestination()
	err := con.Teardown(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
