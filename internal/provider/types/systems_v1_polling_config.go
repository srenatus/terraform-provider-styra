// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SystemsV1PollingConfig struct {
	LongPollingTimeoutSeconds types.Int64 `tfsdk:"long_polling_timeout_seconds"`
	MaxDelaySeconds           types.Int64 `tfsdk:"max_delay_seconds"`
	MinDelaySeconds           types.Int64 `tfsdk:"min_delay_seconds"`
}