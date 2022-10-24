package connectorName

import (
	"context"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Destination struct {
	sdk.UnimplementedDestination

	config DestinationConfig
}

func NewDestination() sdk.Destination {
	return &Destination{}
}

func (d *Destination) Parameters() map[string]sdk.Parameter {
	return map[string]sdk.Parameter{
		GlobalConfigParam: {
			Default:     "localhost:10000",
			Required:    true,
			Description: "The URL of the server.",
		},
		DestinationConfigParam: {
			Default:     "false",
			Required:    false,
			Description: "An optional configuration parameter.",
		},
	}
}

func (d *Destination) Configure(ctx context.Context, cfg map[string]string) error {
	sdk.Logger(ctx).Info().Msg("Configuring a Destination connector...")
	config, err := ParseDestinationConfig(cfg)
	if err != nil {
		return err
	}
	d.config = config
	return nil
}

func (d *Destination) Open(ctx context.Context) error {
	return nil
}

func (d *Destination) Write(ctx context.Context, records []sdk.Record) (int, error) {
	return 0, nil
}

func (d *Destination) Teardown(ctx context.Context) error {
	return nil
}
