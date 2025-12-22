// Package project provides validators for Google Cloud Project identifiers.
package project

import "regexp"

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
	// Check length first (security requirement)
	if len(id) < 6 || len(id) > 30 {
		return false
	}

	// Check for leading/trailing whitespace
	if len(id) > 0 && (id[0] == ' ' || id[len(id)-1] == ' ') {
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
	// Check length first (security requirement)
	if len(name) < 4 || len(name) > 30 {
		return false
	}

	// Check for leading/trailing whitespace
	if len(name) > 0 && (name[0] == ' ' || name[len(name)-1] == ' ') {
		return false
	}

	// Validate format with regex
	return projectNameRegex.MatchString(name)
}
