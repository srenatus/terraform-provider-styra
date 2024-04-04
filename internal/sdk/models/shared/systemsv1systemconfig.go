// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/Styra/terraform-provider-styra/internal/sdk/internal/utils"
)

// SystemsV1SystemConfigTypeParameters - system type parameter values (for template.* types)
type SystemsV1SystemConfigTypeParameters struct {
}

type SystemsV1SystemConfig struct {
	Authz          SystemsV1AuthzConfig           `json:"authz"`
	BundleDownload *SystemsV1BundleDownloadConfig `json:"bundle_download,omitempty"`
	BundleRegistry *SystemsV1BundleRegistryConfig `json:"bundle_registry,omitempty"`
	// only put data in the context bundle
	ContextBundleDataOnly *bool `json:"context_bundle_data_only,omitempty"`
	// list of path prefixes for policies/datasources that go into the second (context) bundle
	ContextBundleRoots []string `json:"context_bundle_roots,omitempty"`
	// datasources created for the system
	Datasources []SystemsV1DatasourceConfig `json:"datasources,omitempty"`
	// location of key attributes and additional columns in the decisions grouped by policy entry point path
	DecisionMappings     map[string]SystemsV1RuleDecisionMappings `json:"decision_mappings,omitempty"`
	DeploymentParameters *SystemsV1SystemDeploymentParameters     `json:"deployment_parameters,omitempty"`
	// description for the system
	Description *string `json:"description,omitempty"`
	// error/warning configuration: one of "all", "errors", "none"
	ErrorSetting *string `json:"error_setting,omitempty"`
	// current deployment errors
	Errors          map[string]SystemsV1AgentErrors `json:"errors,omitempty"`
	ExternalBundles *SystemsV1ExternalBundleConfig  `json:"external_bundles,omitempty"`
	// optional parameter to map Styra DAS system ID to external IDs used by a customer. (mapping can be retrieved with TranslateExternalIds operation)
	ExternalID *string `json:"external_id,omitempty"`
	// when set, stacks that are not linked to this system will be filtered out of its bundles
	FilterStacks *bool `json:"filter_stacks,omitempty"`
	// system ID
	ID string `json:"id"`
	// optional parameter to specify the Kafka topic where the decision logs for this system should be published (ignored if Kafka is not configured for the workspace for decision export)
	KafkaTopic *string `json:"kafka_topic,omitempty"`
	// IDs of stacks matching the system
	MatchingStacks []string         `json:"matching_stacks,omitempty"`
	Metadata       MetaV1ObjectMeta `json:"metadata"`
	// A history of any migrations performed on this system
	MigrationHistory []SystemsV1MigrationRecord `json:"migration_history,omitempty"`
	// minimum running OPA version for the systems
	MinimumOpaVersion *string `json:"minimum_opa_version,omitempty"`
	// enable mock OPAs for this system
	MockOpaEnabled *bool `json:"mock_opa_enabled,omitempty"`
	// system name
	Name string `json:"name"`
	// policies created for the system
	Policies []SystemsV1PolicyConfig `json:"policies"`
	// prevents users from modifying policies using Styra UIs
	ReadOnly      *bool                     `default:"false" json:"read_only"`
	SourceControl *GitV1SourceControlConfig `json:"source_control,omitempty"`
	// system status
	Status string `json:"status"`
	// tokens created for the system
	Tokens []TokensV1Token `json:"tokens,omitempty"`
	// system type e.g. kubernetes
	Type string `json:"type"`
	// system type parameter values (for template.* types)
	TypeParameters *SystemsV1SystemConfigTypeParameters `json:"type_parameters,omitempty"`
	// uninstallation instructions by installation method (deprecated)
	Uninstall map[string]string `json:"uninstall,omitempty"`
}

func (s SystemsV1SystemConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SystemsV1SystemConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SystemsV1SystemConfig) GetAuthz() SystemsV1AuthzConfig {
	if o == nil {
		return SystemsV1AuthzConfig{}
	}
	return o.Authz
}

func (o *SystemsV1SystemConfig) GetBundleDownload() *SystemsV1BundleDownloadConfig {
	if o == nil {
		return nil
	}
	return o.BundleDownload
}

func (o *SystemsV1SystemConfig) GetBundleRegistry() *SystemsV1BundleRegistryConfig {
	if o == nil {
		return nil
	}
	return o.BundleRegistry
}

func (o *SystemsV1SystemConfig) GetContextBundleDataOnly() *bool {
	if o == nil {
		return nil
	}
	return o.ContextBundleDataOnly
}

func (o *SystemsV1SystemConfig) GetContextBundleRoots() []string {
	if o == nil {
		return nil
	}
	return o.ContextBundleRoots
}

func (o *SystemsV1SystemConfig) GetDatasources() []SystemsV1DatasourceConfig {
	if o == nil {
		return nil
	}
	return o.Datasources
}

func (o *SystemsV1SystemConfig) GetDecisionMappings() map[string]SystemsV1RuleDecisionMappings {
	if o == nil {
		return nil
	}
	return o.DecisionMappings
}

func (o *SystemsV1SystemConfig) GetDeploymentParameters() *SystemsV1SystemDeploymentParameters {
	if o == nil {
		return nil
	}
	return o.DeploymentParameters
}

func (o *SystemsV1SystemConfig) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SystemsV1SystemConfig) GetErrorSetting() *string {
	if o == nil {
		return nil
	}
	return o.ErrorSetting
}

func (o *SystemsV1SystemConfig) GetErrors() map[string]SystemsV1AgentErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *SystemsV1SystemConfig) GetExternalBundles() *SystemsV1ExternalBundleConfig {
	if o == nil {
		return nil
	}
	return o.ExternalBundles
}

func (o *SystemsV1SystemConfig) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *SystemsV1SystemConfig) GetFilterStacks() *bool {
	if o == nil {
		return nil
	}
	return o.FilterStacks
}

func (o *SystemsV1SystemConfig) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SystemsV1SystemConfig) GetKafkaTopic() *string {
	if o == nil {
		return nil
	}
	return o.KafkaTopic
}

func (o *SystemsV1SystemConfig) GetMatchingStacks() []string {
	if o == nil {
		return nil
	}
	return o.MatchingStacks
}

func (o *SystemsV1SystemConfig) GetMetadata() MetaV1ObjectMeta {
	if o == nil {
		return MetaV1ObjectMeta{}
	}
	return o.Metadata
}

func (o *SystemsV1SystemConfig) GetMigrationHistory() []SystemsV1MigrationRecord {
	if o == nil {
		return nil
	}
	return o.MigrationHistory
}

func (o *SystemsV1SystemConfig) GetMinimumOpaVersion() *string {
	if o == nil {
		return nil
	}
	return o.MinimumOpaVersion
}

func (o *SystemsV1SystemConfig) GetMockOpaEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.MockOpaEnabled
}

func (o *SystemsV1SystemConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SystemsV1SystemConfig) GetPolicies() []SystemsV1PolicyConfig {
	if o == nil {
		return []SystemsV1PolicyConfig{}
	}
	return o.Policies
}

func (o *SystemsV1SystemConfig) GetReadOnly() *bool {
	if o == nil {
		return nil
	}
	return o.ReadOnly
}

func (o *SystemsV1SystemConfig) GetSourceControl() *GitV1SourceControlConfig {
	if o == nil {
		return nil
	}
	return o.SourceControl
}

func (o *SystemsV1SystemConfig) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *SystemsV1SystemConfig) GetTokens() []TokensV1Token {
	if o == nil {
		return nil
	}
	return o.Tokens
}

func (o *SystemsV1SystemConfig) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *SystemsV1SystemConfig) GetTypeParameters() *SystemsV1SystemConfigTypeParameters {
	if o == nil {
		return nil
	}
	return o.TypeParameters
}

func (o *SystemsV1SystemConfig) GetUninstall() map[string]string {
	if o == nil {
		return nil
	}
	return o.Uninstall
}