// Package project provides validators for Google Cloud Project identifiers.
package project

import (
	"regexp"
	"strings"

	"github.com/UnitVectorY-Labs/gcpvalidate/internal"
	"github.com/UnitVectorY-Labs/gcpvalidate/location"
)

// Compiled regexes for project validation
var (
	projectIDRegex   = regexp.MustCompile(`^[a-z][a-z0-9-]{4,28}[a-z0-9]$`)
	projectNameRegex = regexp.MustCompile(`^[a-zA-Z0-9 '!-]{4,30}$`)
)

// IsValidProjectID validates a Google Cloud Project ID.
//
// A valid Project ID must:
//   - Be 6 to 30 characters long
//   - Start with a lowercase letter
//   - Contain only ASCII lowercase letters, digits, and hyphens
//   - Not end with a hyphen
func IsValidProjectID(id string) bool {
	// Check length first
	if len(id) < 6 || len(id) > 30 {
		return false
	}

	// Check for leading/trailing whitespace
	if !internal.HasTrimmedWhitespace(id) {
		return false
	}

	// Must start with lowercase letter
	if id[0] < 'a' || id[0] > 'z' {
		return false
	}

	// Must not end with hyphen
	if id[len(id)-1] == '-' {
		return false
	}

	// Validate format with regex
	return projectIDRegex.MatchString(id)
}

// IsValidProjectName validates a Google Cloud Project display name.
//
// A valid Project Name must:
//   - Be 4 to 30 characters long
//   - Contain only letters, numbers, single quotes, hyphens, spaces, or exclamation points
func IsValidProjectName(name string) bool {
	// Check length first
	if len(name) < 4 || len(name) > 30 {
		return false
	}

	// Check for leading/trailing whitespace
	if !internal.HasTrimmedWhitespace(name) {
		return false
	}

	// Validate format with regex
	return projectNameRegex.MatchString(name)
}

// IsValidProjectLocationParent validates a project/location parent path.
//
// Accepted structure:
//   - projects/{project}/locations/{location}
//
// Segment rules:
//   - {project} must satisfy IsValidProjectID
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
	if !IsValidProjectID(projectID) {
		return false
	}
	if !location.IsValidLocation(loc) {
		return false
	}

	return true
}
