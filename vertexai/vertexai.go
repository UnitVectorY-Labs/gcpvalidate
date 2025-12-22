// Package vertexai provides validators for Google Cloud Vertex AI identifiers.
package vertexai

import (
	"regexp"
	"strings"

	"github.com/UnitVectorY-Labs/gcpvalidate/internal"
	"github.com/UnitVectorY-Labs/gcpvalidate/location"
	"github.com/UnitVectorY-Labs/gcpvalidate/project"
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

// IsValidVertexModelResourceName validates a Vertex AI model resource path.
//
// Accepted path structures:
//   - projects/{project}/locations/{location}/models/{modelId}
//   - projects/{project}/locations/{location}/publishers/{publisher}/models/{modelId}
//
// Segment rules:
//   - {project} must satisfy project.IsValidProjectID
//   - {location} must satisfy location.IsValidLocation
//   - {publisher} must be non-empty and contain no slashes
//   - {modelId} must be non-empty and contain no slashes
//
// Note: This validates structure only, not whether the resource exists.
func IsValidVertexModelResourceName(path string) bool {
	// Empty string is invalid
	if path == "" {
		return false
	}

	// Check for leading/trailing whitespace
	if !internal.HasTrimmedWhitespace(path) {
		return false
	}

	// Split the path into segments
	segments := strings.Split(path, "/")

	// Check for the two valid formats
	if len(segments) == 6 {
		// Format: projects/{project}/locations/{location}/models/{modelId}
		if segments[0] != "projects" || segments[2] != "locations" || segments[4] != "models" {
			return false
		}

		projectID := segments[1]
		loc := segments[3]
		modelID := segments[5]

		// Validate each segment
		if !project.IsValidProjectID(projectID) {
			return false
		}
		if !location.IsValidLocation(loc) {
			return false
		}
		// Model ID must be valid path segment
		if !internal.IsValidPathSegment(modelID) {
			return false
		}

		return true
	} else if len(segments) == 8 {
		// Format: projects/{project}/locations/{location}/publishers/{publisher}/models/{modelId}
		if segments[0] != "projects" || segments[2] != "locations" || 
		   segments[4] != "publishers" || segments[6] != "models" {
			return false
		}

		projectID := segments[1]
		loc := segments[3]
		publisher := segments[5]
		modelID := segments[7]

		// Validate each segment
		if !project.IsValidProjectID(projectID) {
			return false
		}
		if !location.IsValidLocation(loc) {
			return false
		}
		// Publisher must be valid path segment
		if !internal.IsValidPathSegment(publisher) {
			return false
		}
		// Model ID must be valid path segment
		if !internal.IsValidPathSegment(modelID) {
			return false
		}

		return true
	}

	return false
}
