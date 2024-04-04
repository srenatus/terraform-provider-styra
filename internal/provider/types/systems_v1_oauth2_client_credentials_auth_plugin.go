// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SystemsV1Oauth2ClientCredentialsAuthPlugin struct {
	AdditionalClaims     TypeParameters          `tfsdk:"additional_claims"`
	AdditionalHeaders    map[string]types.String `tfsdk:"additional_headers"`
	AdditionalParameters map[string]types.String `tfsdk:"additional_parameters"`
	ClientID             types.String            `tfsdk:"client_id"`
	ClientSecret         types.String            `tfsdk:"client_secret"`
	GrantType            types.String            `tfsdk:"grant_type"`
	IncludeJtiClaim      types.Bool              `tfsdk:"include_jti_claim"`
	Scopes               []types.String          `tfsdk:"scopes"`
	SigningKey           types.String            `tfsdk:"signing_key"`
	Thumbprint           types.String            `tfsdk:"thumbprint"`
	TokenURL             types.String            `tfsdk:"token_url"`
}