# location package validation rules

[← Back to index](index.md)

The `location` package validates Google Cloud location identifiers.

This package supports **syntax validation**: the string looks like a region or zone token. It does not check whether the location currently exists or is available for a specific product.

## IsValidRegion

**Signature**: `location.IsValidRegion(region string) bool`

### Rules

A region token must:
- Be non-empty.
- Contain only lowercase letters, digits, and hyphens.
- Match the general "region code" style used by Google Cloud products, such as `us-central1`, `us-east5`, and `northamerica-northeast1`.

### Usage example

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/location"

if !location.IsValidRegion(cfg.Region) {
    return fmt.Errorf("invalid region")
}
```

### Notes

- This does not guarantee the region exists or is available for a given service.
- Many Google APIs accept `projects/{project}/locations/{location}` and state that `{location}` must be a valid region.
- Leading and trailing whitespace is rejected.

## IsValidZone

**Signature**: `location.IsValidZone(zone string) bool`

### Rules

A zone token must:
- Be non-empty.
- Follow the zone naming pattern described in Compute Engine docs: a zone name is formed from `<region>-<zone>`, for example `us-central1-a`.

### Usage example

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/location"

if !location.IsValidZone(cfg.Zone) {
    return fmt.Errorf("invalid zone")
}
```

### Notes

- Zones are a Compute concept; many products are regional-only. Still, zone syntax is useful when validating inputs for zonal services.
- Leading and trailing whitespace is rejected.

## IsValidLocation

**Signature**: `location.IsValidLocation(loc string) bool`

### Rules

A location token must:
- Pass either `IsValidRegion` or `IsValidZone`, OR
- Be the literal string `"global"`.

### Usage example

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/location"

if !location.IsValidLocation(cfg.Location) {
    return fmt.Errorf("invalid location")
}
```

### Notes

- This does not guarantee the location currently exists or is available for a given service.
- The `"global"` literal is supported as some services use it.
- Leading and trailing whitespace is rejected.

## References

- [Compute Engine: Regions and zones](https://cloud.google.com/compute/docs/regions-zones) - Describes zone name formation and example `us-central1-a`
- [Vertex AI: Locations](https://cloud.google.com/vertex-ai/docs/general/locations) - Shows region code examples and resource paths like `projects/PROJECT/locations/us-central1/...`
- [Cloud Run API](https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services/list) - States location must be a valid region and uses `projects/{project}/locations/{location}`
