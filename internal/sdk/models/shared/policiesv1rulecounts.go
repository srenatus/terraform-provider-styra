// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PoliciesV1RuleCounts struct {
	// number of allow rules
	Allow int `json:"allow"`
	// number of deny rules
	Deny int `json:"deny"`
	// number of enforce rules
	Enforce int `json:"enforce"`
	// number of ignore rules
	Ignore int `json:"ignore"`
	// number of monitor rules
	Monitor int `json:"monitor"`
	// number of notify rules
	Notify int `json:"notify"`
	// number of unclassified rules
	Other int `json:"other"`
	// number of test rules
	Test int `json:"test"`
	// total number of rules
	Total int `json:"total"`
}

func (o *PoliciesV1RuleCounts) GetAllow() int {
	if o == nil {
		return 0
	}
	return o.Allow
}

func (o *PoliciesV1RuleCounts) GetDeny() int {
	if o == nil {
		return 0
	}
	return o.Deny
}

func (o *PoliciesV1RuleCounts) GetEnforce() int {
	if o == nil {
		return 0
	}
	return o.Enforce
}

func (o *PoliciesV1RuleCounts) GetIgnore() int {
	if o == nil {
		return 0
	}
	return o.Ignore
}

func (o *PoliciesV1RuleCounts) GetMonitor() int {
	if o == nil {
		return 0
	}
	return o.Monitor
}

func (o *PoliciesV1RuleCounts) GetNotify() int {
	if o == nil {
		return 0
	}
	return o.Notify
}

func (o *PoliciesV1RuleCounts) GetOther() int {
	if o == nil {
		return 0
	}
	return o.Other
}

func (o *PoliciesV1RuleCounts) GetTest() int {
	if o == nil {
		return 0
	}
	return o.Test
}

func (o *PoliciesV1RuleCounts) GetTotal() int {
	if o == nil {
		return 0
	}
	return o.Total
}