// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GitV1SourceControlConfig struct {
	Origin GitV1GitRepoConfig `json:"origin"`
}

func (o *GitV1SourceControlConfig) GetOrigin() GitV1GitRepoConfig {
	if o == nil {
		return GitV1GitRepoConfig{}
	}
	return o.Origin
}