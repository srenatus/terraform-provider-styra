// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SystemsV1MigrationRecord struct {
	From           types.String `tfsdk:"from"`
	InitiatedBy    types.String `tfsdk:"initiated_by"`
	InitiatingUser types.String `tfsdk:"initiating_user"`
	MigratedAt     types.String `tfsdk:"migrated_at"`
	Recovered      types.Bool   `tfsdk:"recovered"`
	To             types.String `tfsdk:"to"`
}