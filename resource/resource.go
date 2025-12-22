// Package resource provides validators for Google Cloud resource path structures.
package resource

import (
	"strings"

	"github.com/UnitVectorY-Labs/gcpvalidate/internal"
	"github.com/UnitVectorY-Labs/gcpvalidate/location"
	"github.com/UnitVectorY-Labs/gcpvalidate/project"
)

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
		// Model ID must be non-empty, contain no slashes, and have no whitespace
		if modelID == "" || !internal.HasTrimmedWhitespace(modelID) || strings.ContainsAny(modelID, "/\t\n\r") {
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
		// Publisher must be non-empty, contain no slashes, and have no whitespace
		if publisher == "" || !internal.HasTrimmedWhitespace(publisher) || strings.ContainsAny(publisher, "/\t\n\r") {
			return false
		}
		// Model ID must be non-empty, contain no slashes, and have no whitespace
		if modelID == "" || !internal.HasTrimmedWhitespace(modelID) || strings.ContainsAny(modelID, "/\t\n\r") {
			return false
		}

		return true
	}

	return false
}

// IsValidProjectLocationParent validates a project/location parent path.
//
// Accepted structure:
//   - projects/{project}/locations/{location}
//
// Segment rules:
//   - {project} must satisfy project.IsValidProjectID
//   - {location} must satisfy location.IsValidLocation
//
// Note: This validates structure only, not whether the resource exists.
func IsValidProjectLocationParent(parent string) bool {
	// Empty string is invalid
	if parent == "" {
		return false
	}

	// Check for leading/trailing whitespace
	if !internal.HasTrimmedWhitespace(parent) {
		return false
	}

	// Split the path into segments
	segments := strings.Split(parent, "/")

	// Must have exactly 4 segments
	if len(segments) != 4 {
		return false
	}

	// Check the structure
	if segments[0] != "projects" || segments[2] != "locations" {
		return false
	}

	projectID := segments[1]
	loc := segments[3]

	// Validate each segment
	if !project.IsValidProjectID(projectID) {
		return false
	}
	if !location.IsValidLocation(loc) {
		return false
	}

	return true
}
