// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/Styra/terraform-provider-styra/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *PolicyDataSourceModel) RefreshFromSharedPoliciesV1PolicyGetResponse(resp *shared.PoliciesV1PolicyGetResponse) {
	if resp != nil {
		r.RequestID = types.StringPointerValue(resp.RequestID)
		resultResult, _ := json.Marshal(resp.Result)
		r.Result = types.StringValue(string(resultResult))
	}
}