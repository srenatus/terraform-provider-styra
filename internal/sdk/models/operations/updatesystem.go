// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Styra/terraform-provider-styra/internal/sdk/models/shared"
	"net/http"
)

type UpdateSystemRequest struct {
	// system ID
	System string `pathParam:"style=simple,explode=false,name=system"`
	// if set to '*' then creates a new system with type-specific related objects
	IfNoneMatch                *string                           `header:"style=simple,explode=false,name=If-None-Match"`
	SystemsV1SystemsPutRequest shared.SystemsV1SystemsPutRequest `request:"mediaType=application/json"`
}

func (o *UpdateSystemRequest) GetSystem() string {
	if o == nil {
		return ""
	}
	return o.System
}

func (o *UpdateSystemRequest) GetIfNoneMatch() *string {
	if o == nil {
		return nil
	}
	return o.IfNoneMatch
}

func (o *UpdateSystemRequest) GetSystemsV1SystemsPutRequest() shared.SystemsV1SystemsPutRequest {
	if o == nil {
		return shared.SystemsV1SystemsPutRequest{}
	}
	return o.SystemsV1SystemsPutRequest
}

type UpdateSystemResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	SystemsV1SystemsPutResponse *shared.SystemsV1SystemsPutResponse
	// Bad Request
	MetaV1ErrorResponse *shared.MetaV1ErrorResponse
}

func (o *UpdateSystemResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSystemResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSystemResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSystemResponse) GetSystemsV1SystemsPutResponse() *shared.SystemsV1SystemsPutResponse {
	if o == nil {
		return nil
	}
	return o.SystemsV1SystemsPutResponse
}

func (o *UpdateSystemResponse) GetMetaV1ErrorResponse() *shared.MetaV1ErrorResponse {
	if o == nil {
		return nil
	}
	return o.MetaV1ErrorResponse
}