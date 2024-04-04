// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/Styra/terraform-provider-styra/internal/sdk/internal/utils"
)

type SystemsV1AllowedMapping struct {
	Expected interface{} `json:"expected,omitempty"`
	// when set to true, decision is Allowed when the mapped property IS NOT equal to the expected value
	Negated *bool `default:"false" json:"negated"`
	// dot-separated decision property path
	Path string `json:"path"`
}

func (s SystemsV1AllowedMapping) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SystemsV1AllowedMapping) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SystemsV1AllowedMapping) GetExpected() interface{} {
	if o == nil {
		return nil
	}
	return o.Expected
}

func (o *SystemsV1AllowedMapping) GetNegated() *bool {
	if o == nil {
		return nil
	}
	return o.Negated
}

func (o *SystemsV1AllowedMapping) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}