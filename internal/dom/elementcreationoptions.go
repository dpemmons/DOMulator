package dom

// ElementCreationOptions represents options for element creation per WHATWG DOM specification
type ElementCreationOptions struct {
	// Is specifies the custom element name for customized built-in elements
	Is string `json:"is,omitempty"`
}

// ElementCreationOptionsInput represents the input type for element creation methods
// Can be either a string (for backward compatibility) or ElementCreationOptions
type ElementCreationOptionsInput interface{}

// parseElementCreationOptions parses element creation options from various input types
func parseElementCreationOptions(options ElementCreationOptionsInput) (ElementCreationOptions, error) {
	var result ElementCreationOptions

	if options == nil {
		return result, nil
	}

	switch opts := options.(type) {
	case string:
		// String input is treated as the "is" attribute for backward compatibility
		result.Is = opts
	case ElementCreationOptions:
		result = opts
	case map[string]interface{}:
		// Handle JavaScript object input
		if is, exists := opts["is"]; exists {
			if isStr, ok := is.(string); ok {
				result.Is = isStr
			}
		}
	default:
		// For other types, try to extract as a struct with Is field
		// This handles cases where options might come from JavaScript bindings
		// For now, we'll just ignore unknown types rather than error
	}

	return result, nil
}
