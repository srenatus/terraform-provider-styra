// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SystemsV1OpaConfigServiceDeclarationCredentials struct {
	AzureManagedIdentity *SystemsV1AzureManagedIdentitiesAuthPlugin  `json:"azure_managed_identity,omitempty"`
	Bearer               *SystemsV1BearerAuthPlugin                  `json:"bearer,omitempty"`
	ClientTLS            *SystemsV1ClientTLSAuthPlugin               `json:"client_tls,omitempty"`
	GcpMetadata          *SystemsV1GcpMetadataAuthPlugin             `json:"gcp_metadata,omitempty"`
	Oauth2               *SystemsV1Oauth2ClientCredentialsAuthPlugin `json:"oauth2,omitempty"`
	// authenticate using a custom plugin
	Plugin    *string                        `json:"plugin,omitempty"`
	S3Signing *SystemsV1AwsSigningAuthPlugin `json:"s3_signing,omitempty"`
}

func (o *SystemsV1OpaConfigServiceDeclarationCredentials) GetAzureManagedIdentity() *SystemsV1AzureManagedIdentitiesAuthPlugin {
	if o == nil {
		return nil
	}
	return o.AzureManagedIdentity
}

func (o *SystemsV1OpaConfigServiceDeclarationCredentials) GetBearer() *SystemsV1BearerAuthPlugin {
	if o == nil {
		return nil
	}
	return o.Bearer
}

func (o *SystemsV1OpaConfigServiceDeclarationCredentials) GetClientTLS() *SystemsV1ClientTLSAuthPlugin {
	if o == nil {
		return nil
	}
	return o.ClientTLS
}

func (o *SystemsV1OpaConfigServiceDeclarationCredentials) GetGcpMetadata() *SystemsV1GcpMetadataAuthPlugin {
	if o == nil {
		return nil
	}
	return o.GcpMetadata
}

func (o *SystemsV1OpaConfigServiceDeclarationCredentials) GetOauth2() *SystemsV1Oauth2ClientCredentialsAuthPlugin {
	if o == nil {
		return nil
	}
	return o.Oauth2
}

func (o *SystemsV1OpaConfigServiceDeclarationCredentials) GetPlugin() *string {
	if o == nil {
		return nil
	}
	return o.Plugin
}

func (o *SystemsV1OpaConfigServiceDeclarationCredentials) GetS3Signing() *SystemsV1AwsSigningAuthPlugin {
	if o == nil {
		return nil
	}
	return o.S3Signing
}