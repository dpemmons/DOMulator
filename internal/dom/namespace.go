package dom

import (
	"fmt"
	"strings"
)

// Reserved namespace prefixes based on WHATWG DOM Standard
var (
	xmlNamespace    = "http://www.w3.org/XML/1998/namespace"
	xmlnsNamespace  = "http://www.w3.org/2000/xmlns/"
	htmlNamespace   = "http://www.w3.org/1999/xhtml"
	svgNamespace    = "http://www.w3.org/2000/svg"
	mathMLNamespace = "http://www.w3.org/1998/Math/MathML"
)

// NamespaceError represents a namespace-related error
type NamespaceError struct {
	Message string
}

func (e *NamespaceError) Error() string {
	return e.Message
}

// InvalidCharacterError represents an invalid character error
type InvalidCharacterError struct {
	Message string
}

func (e *InvalidCharacterError) Error() string {
	return e.Message
}

// ValidateQualifiedName validates a qualified name according to XML Name production rules
// Implements the "validate" algorithm from WHATWG DOM Standard Section 1.4
func ValidateQualifiedName(qualifiedName string) error {
	if qualifiedName == "" {
		return &InvalidCharacterError{Message: "Qualified name cannot be empty"}
	}

	// Check for invalid characters that would make this not an XML Name
	if !isValidXMLName(qualifiedName) {
		return &InvalidCharacterError{Message: fmt.Sprintf("'%s' is not a valid XML Name", qualifiedName)}
	}

	// Check for multiple colons (only one allowed for namespace prefix separation)
	colonCount := strings.Count(qualifiedName, ":")
	if colonCount > 1 {
		return &NamespaceError{Message: "Qualified name cannot contain more than one colon"}
	}

	// If there's a colon, validate prefix and local name separately
	if colonCount == 1 {
		parts := strings.SplitN(qualifiedName, ":", 2)
		prefix := parts[0]
		localName := parts[1]

		if prefix == "" {
			return &NamespaceError{Message: "Qualified name cannot have empty prefix"}
		}

		if localName == "" {
			return &NamespaceError{Message: "Qualified name cannot have empty local name"}
		}

		if !isValidXMLName(prefix) {
			return &InvalidCharacterError{Message: fmt.Sprintf("Prefix '%s' is not a valid XML Name", prefix)}
		}

		if !isValidXMLName(localName) {
			return &InvalidCharacterError{Message: fmt.Sprintf("Local name '%s' is not a valid XML Name", localName)}
		}
	}

	return nil
}

// ValidateAndExtract validates a qualified name and extracts namespace, prefix, and local name
// Implements the "validate and extract" algorithm from WHATWG DOM Standard Section 1.4
func ValidateAndExtract(namespace, qualifiedName string) (ns, prefix, localName string, err error) {
	// First validate the qualified name
	if err = ValidateQualifiedName(qualifiedName); err != nil {
		return "", "", "", err
	}

	// Extract prefix and local name
	if strings.Contains(qualifiedName, ":") {
		parts := strings.SplitN(qualifiedName, ":", 2)
		prefix = parts[0]
		localName = parts[1]
	} else {
		prefix = ""
		localName = qualifiedName
	}

	// Validate namespace and prefix combinations per WHATWG DOM Standard

	// 1. If namespace is null and prefix is not null, throw NamespaceError
	if namespace == "" && prefix != "" {
		return "", "", "", &NamespaceError{Message: "Cannot have a prefix without a namespace"}
	}

	// 2. If prefix is "xml" and namespace is not XML namespace, throw NamespaceError
	if prefix == "xml" && namespace != xmlNamespace {
		return "", "", "", &NamespaceError{Message: "Prefix 'xml' can only be used with XML namespace"}
	}

	// 3. If qualified name or prefix is "xmlns" and namespace is not XMLNS namespace, throw NamespaceError
	if (qualifiedName == "xmlns" || prefix == "xmlns") && namespace != xmlnsNamespace {
		return "", "", "", &NamespaceError{Message: "Prefix 'xmlns' can only be used with XMLNS namespace"}
	}

	// 4. If namespace is XMLNS namespace and neither qualified name nor prefix is "xmlns", throw NamespaceError
	if namespace == xmlnsNamespace && qualifiedName != "xmlns" && prefix != "xmlns" {
		return "", "", "", &NamespaceError{Message: "XMLNS namespace can only be used with 'xmlns' prefix"}
	}

	return namespace, prefix, localName, nil
}

// isValidXMLName checks if a string is a valid XML Name according to XML 1.0 production rules
func isValidXMLName(name string) bool {
	if name == "" {
		return false
	}

	runes := []rune(name)

	// First character must be NameStartChar
	if !isNameStartChar(runes[0]) {
		return false
	}

	// Remaining characters must be NameChar
	for i := 1; i < len(runes); i++ {
		if !isNameChar(runes[i]) {
			return false
		}
	}

	return true
}

// isNameStartChar checks if a rune is a valid XML NameStartChar
func isNameStartChar(r rune) bool {
	return (r >= 'A' && r <= 'Z') ||
		(r >= 'a' && r <= 'z') ||
		r == '_' || r == ':' ||
		(r >= 0x00C0 && r <= 0x00D6) ||
		(r >= 0x00D8 && r <= 0x00F6) ||
		(r >= 0x00F8 && r <= 0x02FF) ||
		(r >= 0x0370 && r <= 0x037D) ||
		(r >= 0x037F && r <= 0x1FFF) ||
		(r >= 0x200C && r <= 0x200D) ||
		(r >= 0x2070 && r <= 0x218F) ||
		(r >= 0x2C00 && r <= 0x2FEF) ||
		(r >= 0x3001 && r <= 0xD7FF) ||
		(r >= 0xF900 && r <= 0xFDCF) ||
		(r >= 0xFDF0 && r <= 0xFFFD)
}

// isNameChar checks if a rune is a valid XML NameChar
func isNameChar(r rune) bool {
	return isNameStartChar(r) ||
		(r >= '0' && r <= '9') ||
		r == '-' || r == '.' || r == 0x00B7 ||
		(r >= 0x0300 && r <= 0x036F) ||
		(r >= 0x203F && r <= 0x2040)
}

// GetWellKnownNamespace returns the URI for well-known namespace prefixes
func GetWellKnownNamespace(prefix string) string {
	switch prefix {
	case "xml":
		return xmlNamespace
	case "xmlns":
		return xmlnsNamespace
	case "html":
		return htmlNamespace
	case "svg":
		return svgNamespace
	case "mathml", "math":
		return mathMLNamespace
	default:
		return ""
	}
}

// NamespaceInfo holds extracted namespace information
type NamespaceInfo struct {
	NamespaceURI string
	Prefix       string
	LocalName    string
}

// ParseQualifiedName parses a qualified name and returns namespace information
func ParseQualifiedName(qualifiedName string) *NamespaceInfo {
	if strings.Contains(qualifiedName, ":") {
		parts := strings.SplitN(qualifiedName, ":", 2)
		return &NamespaceInfo{
			NamespaceURI: GetWellKnownNamespace(parts[0]),
			Prefix:       parts[0],
			LocalName:    parts[1],
		}
	}

	return &NamespaceInfo{
		NamespaceURI: "",
		Prefix:       "",
		LocalName:    qualifiedName,
	}
}
