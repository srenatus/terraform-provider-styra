// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"github.com/Styra/terraform-provider-styra/internal/sdk/internal/hooks"
	"github.com/Styra/terraform-provider-styra/internal/sdk/internal/utils"
	"github.com/Styra/terraform-provider-styra/internal/sdk/models/shared"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://TENANT.styra.com/",
	// The Styra DAS API server for the current tenant.
	"https://{dasId}.styra.com",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	ServerDefaults    []map[string]string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
	Hooks             *hooks.Hooks
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], c.ServerDefaults[c.ServerIndex]
}

// StyraDas - Styra API: Styra DAS is entirely API-driven.
//
// Access to the APIs requires authentication that should be provided as an Authorization HTTP header including a Styra DAS-issued token:
//
// `Authorization: Bearer <YOURTOKENHERE>`
//
// To request a token you need to have an Styra account, and create a token via the API Tokens menu.
//
// https://docs.styra.com - Styra DAS Documentation
type StyraDas struct {
	// API to create and manage libraries
	Libraries *Libraries
	// Policy management
	Policies *Policies
	// Secrets Management
	Secrets *Secrets
	// Stacks management
	Stacks *Stacks
	// Systems management
	Systems *Systems

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*StyraDas)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *StyraDas) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *StyraDas) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *StyraDas) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithDasID allows setting the dasId variable for url substitution
func WithDasID(dasID string) SDKOption {
	return func(sdk *StyraDas) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["dasId"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["dasId"] = fmt.Sprintf("%v", dasID)
		}
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *StyraDas) {
		sdk.sdkConfiguration.Client = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *StyraDas) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *StyraDas) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *StyraDas) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *StyraDas {
	sdk := &StyraDas{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "2.0.0",
			SDKVersion:        "0.0.1",
			GenVersion:        "2.298.2",
			UserAgent:         "speakeasy-sdk/go 0.0.1 2.298.2 2.0.0 github.com/Styra/terraform-provider-styra/internal/sdk",
			ServerDefaults: []map[string]string{
				{},
				{
					"dasId": "example",
				},
			},
			Hooks: hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Libraries = newLibraries(sdk.sdkConfiguration)

	sdk.Policies = newPolicies(sdk.sdkConfiguration)

	sdk.Secrets = newSecrets(sdk.sdkConfiguration)

	sdk.Stacks = newStacks(sdk.sdkConfiguration)

	sdk.Systems = newSystems(sdk.sdkConfiguration)

	return sdk
}