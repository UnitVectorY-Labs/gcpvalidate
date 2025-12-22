// Package storage provides validators for Google Cloud Storage identifiers.
package storage

import (
	"regexp"
	"strings"
)

// Compiled regex for bucket name validation
var (
	bucketNameRegex = regexp.MustCompile(`^[a-z0-9][a-z0-9._-]*[a-z0-9]$`)
	// IP address pattern to reject (e.g., 192.168.1.1)
	ipAddressRegex = regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$`)
)

// IsValidBucketName validates a Google Cloud Storage bucket name.
//
// A valid bucket name must:
//   - Contain only lowercase letters, digits, dashes (-), underscores (_), and dots (.)
//   - Start and end with a number or letter
//   - Be 3 to 63 characters long (or up to 222 for dotted names)
//   - Not contain consecutive dots
//   - Not look like an IP address
//
// Note: This validates format only, not whether the bucket name is available.
func IsValidBucketName(name string) bool {
	// Empty string is invalid
	if name == "" {
		return false
	}

	// Check for leading/trailing whitespace
	if name[0] == ' ' || name[len(name)-1] == ' ' {
		return false
	}

	// Check length constraints
	if len(name) < 3 {
		return false
	}

	// If name contains dots, max length is 222, otherwise 63
	if strings.Contains(name, ".") {
		if len(name) > 222 {
			return false
		}
		// Check each dot-separated component is <= 63 characters
		parts := strings.Split(name, ".")
		for _, part := range parts {
			if len(part) > 63 {
				return false
			}
		}
		// Check for consecutive dots
		if strings.Contains(name, "..") {
			return false
		}
	} else {
		if len(name) > 63 {
			return false
		}
	}

	// Must start and end with alphanumeric
	if !isAlphanumeric(name[0]) || !isAlphanumeric(name[len(name)-1]) {
		return false
	}

	// Must not look like an IP address
	if ipAddressRegex.MatchString(name) {
		return false
	}

	// Validate format with regex
	return bucketNameRegex.MatchString(name)
}

// isAlphanumeric checks if a byte is a lowercase letter or digit
func isAlphanumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}
