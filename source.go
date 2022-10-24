package connectorName

import (
	"context"

	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Source struct {
	sdk.UnimplementedSource

	config           SourceConfig
	lastPositionRead sdk.Position //nolint:unused // this is just an example
}

func NewSource() sdk.Source {
	return &Source{}
}

func (s *Source) Parameters() map[string]sdk.Parameter {
	return map[string]sdk.Parameter{
		GlobalConfigParam: {
			Default:     "localhost:10000",
			Required:    true,
			Description: "The URL of the server.",
		},
		SourceConfigParam: {
			Default:     "false",
			Required:    false,
			Description: "An optional configuration parameter.",
		},
	}
}

func (s *Source) Configure(ctx context.Context, cfg map[string]string) error {
	sdk.Logger(ctx).Info().Msg("Configuring a Source Connector...")
	config, err := ParseSourceConfig(cfg)
	if err != nil {
		return err
	}

	s.config = config
	return nil
}

func (s *Source) Open(ctx context.Context, pos sdk.Position) error {
	return nil
}

func (s *Source) Read(ctx context.Context) (sdk.Record, error) {
	return sdk.Record{}, nil
}

func (s *Source) Ack(ctx context.Context, position sdk.Position) error {
	return nil
}

func (s *Source) Teardown(ctx context.Context) error {
	return nil
}
