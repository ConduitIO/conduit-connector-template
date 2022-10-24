package connectorName

import "errors"

const (
	GlobalConfigParam      = "global_config"
	SourceConfigParam      = "source_config"
	DestinationConfigParam = "destination_config"
)

var Required = []string{GlobalConfigParam}

var (
	ErrEmptyConfig = errors.New("missing or empty config")
)

type Config struct {
	globalConfigParam string
}

type SourceConfig struct {
	Config
	sourceConfigParam string
}

type DestinationConfig struct {
	Config
	destinationConfigParam string
}

func ParseSourceConfig(cfg map[string]string) (SourceConfig, error) {
	err := checkEmpty(cfg)
	if err != nil {
		return SourceConfig{}, err
	}
	return SourceConfig{
		Config:            Config{globalConfigParam: cfg[GlobalConfigParam]},
		sourceConfigParam: cfg[SourceConfigParam],
	}, nil
}

func ParseDestinationConfig(cfg map[string]string) (DestinationConfig, error) {
	err := checkEmpty(cfg)
	if err != nil {
		return DestinationConfig{}, err
	}
	return DestinationConfig{
		Config:                 Config{globalConfigParam: cfg[GlobalConfigParam]},
		destinationConfigParam: cfg[DestinationConfigParam],
	}, nil
}

func checkEmpty(cfg map[string]string) error {
	if len(cfg) == 0 {
		return ErrEmptyConfig
	}
	return nil
}
