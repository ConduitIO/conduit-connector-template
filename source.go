package connectorname

//go:generate paramgen -output=paramgen_src.go SourceConfig

import (
	"context"
	"fmt"

	"github.com/conduitio/conduit-commons/config"
	"github.com/conduitio/conduit-commons/opencdc"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Source struct {
	sdk.UnimplementedSource

	config           SourceConfig
	lastPositionRead opencdc.Position //nolint:unused // this is just an example
}

type SourceConfig struct {
	// Config includes parameters that are the same in the source and destination.
	Config
	// SourceConfigParam is named foo and must be provided by the user.
	SourceConfigParam string `json:"foo" validate:"required"`
}

func NewSource() sdk.Source {
	// Create Source and wrap it in the default middleware.
	return sdk.SourceWithMiddleware(&Source{}, sdk.DefaultSourceMiddleware()...)
}

func (s *Source) Parameters() config.Parameters {
	// Parameters is a map of named Parameters that describe how to configure
	// the Source. Parameters can be generated from SourceConfig with paramgen.
	return s.config.Parameters()
}

func (s *Source) Configure(ctx context.Context, cfg config.Config) error {
	// Configure is the first function to be called in a connector. It provides
	// the connector with the configuration that can be validated and stored.
	// In case the configuration is not valid it should return an error.
	// Testing if your connector can reach the configured data source should be
	// done in Open, not in Configure.
	// The SDK will validate the configuration and populate default values
	// before calling Configure. If you need to do more complex validations you
	// can do them manually here.

	sdk.Logger(ctx).Info().Msg("Configuring Source...")
	err := sdk.Util.ParseConfig(ctx, cfg, &s.config, NewSource().Parameters())
	if err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}
	return nil
}

func (s *Source) Open(_ context.Context, _ opencdc.Position) error {
	// Open is called after Configure to signal the plugin it can prepare to
	// start producing records. If needed, the plugin should open connections in
	// this function. The position parameter will contain the position of the
	// last record that was successfully processed, Source should therefore
	// start producing records after this position. The context passed to Open
	// will be cancelled once the plugin receives a stop signal from Conduit.
	return nil
}

func (s *Source) ReadN(context.Context, int) ([]opencdc.Record, error) {
	// ReadN is the same as Read, but returns a batch of records. The connector
	// is expected to return at most n records. If there are fewer records
	// available, it should return all of them. If there are no records available
	// it should block until there are records available or the context is
	// cancelled. If the context is cancelled while ReadN is running, it should
	// return the context error.
	return []opencdc.Record{}, nil
}

func (s *Source) Ack(_ context.Context, _ opencdc.Position) error {
	// Ack signals to the implementation that the record with the supplied
	// position was successfully processed. This method might be called after
	// the context of Read is already cancelled, since there might be
	// outstanding acks that need to be delivered. When Teardown is called it is
	// guaranteed there won't be any more calls to Ack.
	// Ack can be called concurrently with Read.
	return nil
}

func (s *Source) Teardown(_ context.Context) error {
	// Teardown signals to the plugin that there will be no more calls to any
	// other function. After Teardown returns, the plugin should be ready for a
	// graceful shutdown.
	return nil
}
