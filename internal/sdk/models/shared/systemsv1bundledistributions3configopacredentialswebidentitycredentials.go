// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials struct {
	AwsRegion   string `json:"aws_region"`
	SessionName string `json:"session_name"`
}

func (o *SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials) GetAwsRegion() string {
	if o == nil {
		return ""
	}
	return o.AwsRegion
}

func (o *SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials) GetSessionName() string {
	if o == nil {
		return ""
	}
	return o.SessionName
}