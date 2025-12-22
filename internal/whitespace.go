// Package internal provides shared utilities for gcpvalidate.
package internal

import "strings"

// HasTrimmedWhitespace checks if a string has no leading or trailing whitespace.
// Returns true if the string is trimmed (valid), false if it has whitespace (invalid).
func HasTrimmedWhitespace(s string) bool {
	return strings.TrimSpace(s) == s
}

// IsValidPathSegment checks if a string is valid for use in a resource path segment.
// It must be non-empty, have no whitespace, and contain no slashes.
func IsValidPathSegment(s string) bool {
	return s != "" && HasTrimmedWhitespace(s) && !strings.Contains(s, "/")
}
