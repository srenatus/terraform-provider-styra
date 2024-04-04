// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SystemsV1VerificationConfig struct {
	// files in the bundle to exclude during verification
	ExcludeFiles []string `json:"exclude_files,omitempty"`
	// name of the key to use for bundle signature verification
	Keyid *string `json:"keyid,omitempty"`
	// information about necessary public signing keys
	PublicKeys map[string]SystemsV1KeyConfig `json:"public_keys,omitempty"`
	// scope to use for bundle signature verification
	Scope *string `json:"scope,omitempty"`
}

func (o *SystemsV1VerificationConfig) GetExcludeFiles() []string {
	if o == nil {
		return nil
	}
	return o.ExcludeFiles
}

func (o *SystemsV1VerificationConfig) GetKeyid() *string {
	if o == nil {
		return nil
	}
	return o.Keyid
}

func (o *SystemsV1VerificationConfig) GetPublicKeys() map[string]SystemsV1KeyConfig {
	if o == nil {
		return nil
	}
	return o.PublicKeys
}

func (o *SystemsV1VerificationConfig) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}