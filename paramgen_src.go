// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-commons/tree/main/paramgen

package connectorname

import (
	"github.com/conduitio/conduit-commons/config"
)

const (
	SourceConfigFoo                   = "foo"
	SourceConfigGlobalConfigParamName = "global_config_param_name"
)

func (SourceConfig) Parameters() map[string]config.Parameter {
	return map[string]config.Parameter{
		SourceConfigFoo: {
			Default:     "",
			Description: "SourceConfigParam is named foo and must be provided by the user.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		SourceConfigGlobalConfigParamName: {
			Default:     "",
			Description: "GlobalConfigParam is named global_config_param_name and needs to be\nprovided by the user.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
	}
}
