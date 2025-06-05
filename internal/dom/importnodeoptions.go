package dom

import (
	"encoding/json"
	"fmt"
)

// ImportNodeOptions represents the options for importNode operations per WHATWG DOM specification
type ImportNodeOptions struct {
	CustomElementRegistry interface{} `json:"customElementRegistry,omitempty"`
	SelfOnly              bool        `json:"selfOnly,omitempty"`
}

// ImportNodeOptionsInput represents flexible input for importNode options
// Can be bool (legacy), ImportNodeOptions struct, or map (JavaScript compatibility)
type ImportNodeOptionsInput interface{}

// parseImportNodeOptions parses various input types into ImportNodeOptions
func parseImportNodeOptions(input ImportNodeOptionsInput) (ImportNodeOptions, error) {
	if input == nil {
		return ImportNodeOptions{SelfOnly: false}, nil
	}

	switch v := input.(type) {
	case bool:
		// Legacy boolean parameter: true means deep (subtree), false means shallow (selfOnly)
		return ImportNodeOptions{
			SelfOnly: !v, // Invert: deep=true becomes selfOnly=false
		}, nil

	case ImportNodeOptions:
		// Direct struct
		return v, nil

	case map[string]interface{}:
		// JavaScript object/map input
		opts := ImportNodeOptions{}

		if selfOnly, exists := v["selfOnly"]; exists {
			if selfOnlyBool, ok := selfOnly.(bool); ok {
				opts.SelfOnly = selfOnlyBool
			}
		}

		if registry, exists := v["customElementRegistry"]; exists {
			opts.CustomElementRegistry = registry
		}

		return opts, nil

	case string:
		// JSON string input
		var opts ImportNodeOptions
		if err := json.Unmarshal([]byte(v), &opts); err != nil {
			return ImportNodeOptions{}, fmt.Errorf("invalid JSON for ImportNodeOptions: %v", err)
		}
		return opts, nil

	default:
		return ImportNodeOptions{}, fmt.Errorf("unsupported ImportNodeOptionsInput type: %T", input)
	}
}
