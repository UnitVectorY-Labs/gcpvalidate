// Package internal provides shared utilities for gcpvalidate.
package internal

import "strings"

// HasTrimmedWhitespace checks if a string has no leading or trailing whitespace.
// Returns true if the string is trimmed (valid), false if it has whitespace (invalid).
func HasTrimmedWhitespace(s string) bool {
	return strings.TrimSpace(s) == s
}

// IsValidPathSegment checks if a string is valid for use in a resource path segment.
// It must be non-empty, contain no whitespace anywhere, and contain no slashes.
func IsValidPathSegment(s string) bool {
	if s == "" {
		return false
	}
	// Check for any whitespace character anywhere in the string
	if strings.ContainsAny(s, " \t\n\r\v\f") {
		return false
	}
	// Check for slashes
	if strings.Contains(s, "/") {
		return false
	}
	return true
}
