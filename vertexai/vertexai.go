// Package vertexai provides validators for Google Cloud Vertex AI identifiers.
package vertexai

import (
	"regexp"

	"github.com/UnitVectorY-Labs/gcpvalidate/internal"
)

// Compiled regex for Vertex AI name validation
var (
	vertexNameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)
)

// IsValidVertexModelName validates a Vertex AI model name.
//
// A valid model name must:
//   - Contain only letters, numbers, dashes, and underscores
//   - Be case-sensitive
//   - Start with a letter
//   - Be no more than 128 characters long
//
// Note: This validates the display name, not server-assigned numeric IDs
// or publisher model IDs.
func IsValidVertexModelName(name string) bool {
	// Empty string is invalid
	if name == "" {
		return false
	}

	// Check for leading/trailing whitespace
	if !internal.HasTrimmedWhitespace(name) {
		return false
	}

	// Check length first
	if len(name) > 128 {
		return false
	}

	// Must start with a letter
	if !((name[0] >= 'a' && name[0] <= 'z') || (name[0] >= 'A' && name[0] <= 'Z')) {
		return false
	}

	// Validate format with regex
	return vertexNameRegex.MatchString(name)
}

// IsValidVertexEndpointName validates a Vertex AI endpoint name.
//
// A valid endpoint name must:
//   - Contain only letters, numbers, dashes, and underscores
//   - Be case-sensitive
//   - Start with a letter
//   - Be no more than 128 characters long
//
// Note: The naming rules are identical to IsValidVertexModelName.
func IsValidVertexEndpointName(name string) bool {
	// Uses the same validation rules as model names
	return IsValidVertexModelName(name)
}
