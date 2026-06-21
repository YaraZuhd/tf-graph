package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadState reads a terraform show -json file from disk and parses it
// into a TerraformState struct.
func LoadState(path string) (*TerraformState, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading state file %q: %w", path, err)
	}

	var state TerraformState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("parsing terraform JSON: %w", err)
	}

	return &state, nil
}
