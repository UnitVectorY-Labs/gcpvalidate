// Package location provides validators for Google Cloud location identifiers.
package location

import "regexp"

// Compiled regexes for location validation
var (
	// Matches region patterns like us-central1, europe-west4, northamerica-northeast1
	regionRegex = regexp.MustCompile(`^[a-z]+(-[a-z]+)*[0-9]+$`)
	// Matches zone patterns like us-central1-a
	zoneRegex = regexp.MustCompile(`^[a-z]+(-[a-z]+)*[0-9]+-[a-z]$`)
)

// IsValidRegion validates a Google Cloud region identifier.
//
// A valid region must:
//   - Be non-empty
//   - Contain only lowercase letters, digits, and hyphens
//   - Match the general region code style (e.g., us-central1, europe-west4)
//
// Note: This validates syntax only, not whether the region currently exists
// or is available for a given service.
func IsValidRegion(region string) bool {
	// Empty string is invalid
	if region == "" {
		return false
	}

	// Check length limit (reasonable maximum)
	if len(region) > 100 {
		return false
	}

	// Check for leading/trailing whitespace
	if region[0] == ' ' || region[len(region)-1] == ' ' {
		return false
	}

	// Validate format with regex
	return regionRegex.MatchString(region)
}

// IsValidZone validates a Google Cloud zone identifier.
//
// A valid zone must:
//   - Be non-empty
//   - Follow the zone naming pattern: <region>-<zone-letter> (e.g., us-central1-a)
//
// Note: This validates syntax only, not whether the zone currently exists.
func IsValidZone(zone string) bool {
	// Empty string is invalid
	if zone == "" {
		return false
	}

	// Check length limit (reasonable maximum)
	if len(zone) > 100 {
		return false
	}

	// Check for leading/trailing whitespace
	if zone[0] == ' ' || zone[len(zone)-1] == ' ' {
		return false
	}

	// Validate format with regex
	return zoneRegex.MatchString(zone)
}

// IsValidLocation validates a Google Cloud location identifier.
//
// A valid location must be either:
//   - A valid region (validated by IsValidRegion)
//   - A valid zone (validated by IsValidZone)
//   - The literal string "global"
//
// Note: This validates syntax only, not whether the location currently exists
// or is available for a given service.
func IsValidLocation(loc string) bool {
	// Check for "global" literal
	if loc == "global" {
		return true
	}

	// Check if it's a valid region or zone
	return IsValidRegion(loc) || IsValidZone(loc)
}
