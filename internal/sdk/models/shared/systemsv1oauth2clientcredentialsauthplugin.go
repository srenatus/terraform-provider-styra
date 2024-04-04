// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AdditionalClaims - map of claims to include in the JWT
type AdditionalClaims struct {
}

type SystemsV1Oauth2ClientCredentialsAuthPlugin struct {
	// map of claims to include in the JWT
	AdditionalClaims AdditionalClaims `json:"additional_claims"`
	// map of additional headers to send to token endpoint at the OAuth2 authorization server
	AdditionalHeaders map[string]string `json:"additional_headers,omitempty"`
	// map of additional body parameters to send token endpoint at the OAuth2 authorization server
	AdditionalParameters map[string]string `json:"additional_parameters,omitempty"`
	// the client ID to use for authentication
	ClientID string `json:"client_id"`
	// the client secret to use for authentication
	ClientSecret string `json:"client_secret"`
	// defaults to client_credentials
	GrantType string `json:"grant_type"`
	// include a uniquely generated jti claim in any issued JWT
	IncludeJtiClaim bool `json:"include_jti_claim"`
	// optional list of scopes to request for the token
	Scopes []string `json:"scopes,omitempty"`
	// reference to private key used for signing the JWT
	SigningKey string `json:"signing_key"`
	// certificate thumbprint to use for x5t header generation
	Thumbprint string `json:"thumbprint"`
	// URL pointing to the token endpoint at the OAuth2 authorization server
	TokenURL string `json:"token_url"`
}

func (o *SystemsV1Oauth2ClientCredentialsAuthPlugin) GetAdditionalClaims() AdditionalClaims {
	if o == nil {
		return AdditionalClaims{}
	}
	return o.AdditionalClaims
}

func (o *SystemsV1Oauth2ClientCredentialsAuthPlugin) GetAdditionalHeaders() map[string]string {
	if o == nil {
		return nil
	}
	return o.AdditionalHeaders
}

func (o *SystemsV1Oauth2ClientCredentialsAuthPlugin) GetAdditionalParameters() map[string]string {
	if o == nil {
		return nil
	}
	return o.AdditionalParameters
}

func (o *SystemsV1Oauth2ClientCredentialsAuthPlugin) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SystemsV1Oauth2ClientCredentialsAuthPlugin) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SystemsV1Oauth2ClientCredentialsAuthPlugin) GetGrantType() string {
	if o == nil {
		return ""
	}
	return o.GrantType
}

func (o *SystemsV1Oauth2ClientCredentialsAuthPlugin) GetIncludeJtiClaim() bool {
	if o == nil {
		return false
	}
	return o.IncludeJtiClaim
}

func (o *SystemsV1Oauth2ClientCredentialsAuthPlugin) GetScopes() []string {
	if o == nil {
		return nil
	}
	return o.Scopes
}

func (o *SystemsV1Oauth2ClientCredentialsAuthPlugin) GetSigningKey() string {
	if o == nil {
		return ""
	}
	return o.SigningKey
}

func (o *SystemsV1Oauth2ClientCredentialsAuthPlugin) GetThumbprint() string {
	if o == nil {
		return ""
	}
	return o.Thumbprint
}

func (o *SystemsV1Oauth2ClientCredentialsAuthPlugin) GetTokenURL() string {
	if o == nil {
		return ""
	}
	return o.TokenURL
}