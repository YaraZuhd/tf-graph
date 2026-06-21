package parser

// TerraformState mirrors the top-level shape of `terraform show -json` output.
type TerraformState struct {
	FormatVersion    string `json:"format_version"`
	TerraformVersion string `json:"terraform_version"`
	Values           Values `json:"values"`
}

type Values struct {
	RootModule RootModule `json:"root_module"`
}

type RootModule struct {
	Resources []Resource `json:"resources"`
}

// Resource represents a single managed resource in the state.
// We only need a subset of fields for graph rendering: identity + dependencies.
type Resource struct {
	Address      string                 `json:"address"`
	Mode         string                 `json:"mode"`
	Type         string                 `json:"type"`
	Name         string                 `json:"name"`
	ProviderName string                 `json:"provider_name"`
	Values       map[string]interface{} `json:"values"`
	DependsOn    []string               `json:"depends_on"`
}
