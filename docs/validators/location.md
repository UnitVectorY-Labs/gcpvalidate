---
layout: default
title: Location
parent: Validators
nav_order: 2
permalink: /validators/location
---

# Location

Validators for Google Cloud location identifiers. These validate **syntax only**—not whether the location currently exists or is available for a specific service.

## IsValidRegion

**Signature**: `location.IsValidRegion(region string) bool`

**Example**:

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/location"

// Valid regions
location.IsValidRegion("us-central1")             // true
location.IsValidRegion("europe-west4")            // true
location.IsValidRegion("northamerica-northeast1") // true

// Invalid regions
location.IsValidRegion("us-central1-a")           // false - zone, not region
location.IsValidRegion("US-CENTRAL1")             // false - uppercase not allowed
location.IsValidRegion("global")                  // false - use IsValidLocation for "global"
```

**Rules**:

- Contain only lowercase letters, digits, and hyphens
- Match region code style (e.g., `us-central1`, `europe-west4`)

**Note**: Accepts region-like patterns. Does not verify the region exists.

## IsValidZone

**Signature**: `location.IsValidZone(zone string) bool`

**Example**:

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/location"

// Valid zones
location.IsValidZone("us-central1-a")             // true
location.IsValidZone("europe-west4-b")            // true

// Invalid zones
location.IsValidZone("us-central1")               // false - region, not zone
location.IsValidZone("US-CENTRAL1-A")             // false - uppercase not allowed
```

**Rules**:

- Follow zone naming pattern: `<region>-<zone-letter>` (e.g., `us-central1-a`)

**Note**: Zones are primarily a Compute Engine concept. Many products are regional-only.

## IsValidLocation

**Signature**: `location.IsValidLocation(loc string) bool`

**Example**:

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/location"

// Valid locations
location.IsValidLocation("global")                // true
location.IsValidLocation("us-central1")           // true (region)
location.IsValidLocation("us-central1-a")         // true (zone)

// Invalid locations
location.IsValidLocation("GLOBAL")                // false - case sensitive
location.IsValidLocation("invalid location")      // false - spaces not allowed
```

**Rules**:

- Pass `IsValidRegion` or `IsValidZone`, OR
- Be the literal string `"global"`

**References**:
- [Compute Engine: Regions and zones](https://cloud.google.com/compute/docs/regions-zones)
- [Vertex AI: Locations](https://cloud.google.com/vertex-ai/docs/general/locations)

