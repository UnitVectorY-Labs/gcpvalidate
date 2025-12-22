---
layout: default
title: gcpvalidate
nav_order: 1
permalink: /
---

# gcpvalidate Validation Rules

Client-side syntactic validation for Google Cloud resource identifiers. Validates format and conventions only—no existence checks, API calls, or IAM verification.

## Validators

| Package | Validator | Description |
|---------|-----------|-------------|
| [project](project.md) | `IsValidProjectID` | GCP project identifier (6-30 chars, lowercase, starts with letter) |
| [project](project.md) | `IsValidProjectName` | GCP project display name (4-30 chars, letters/numbers/spaces/punctuation) |
| [location](location.md) | `IsValidRegion` | Region identifier (e.g., `us-central1`, `europe-west4`) |
| [location](location.md) | `IsValidZone` | Zone identifier (e.g., `us-central1-a`) |
| [location](location.md) | `IsValidLocation` | Location identifier (region, zone, or `global`) |
| [storage](storage.md) | `IsValidBucketName` | Cloud Storage bucket name (3-63 chars, lowercase, alphanumeric with dots/dashes/underscores) |
| [vertexai](vertexai.md) | `IsValidVertexModelName` | Vertex AI model display name (max 128 chars, starts with letter) |
| [vertexai](vertexai.md) | `IsValidVertexEndpointName` | Vertex AI endpoint display name (max 128 chars, starts with letter) |
| [resource](resource.md) | `IsValidVertexModelResourceName` | Full Vertex AI model resource path |
| [resource](resource.md) | `IsValidProjectLocationParent` | Project/location parent path (`projects/{project}/locations/{location}`) |

## Important Disclaimers

- **No existence checks**: Validators confirm syntax only, not whether a resource exists
- **No API calls**: All validation is local
- **No IAM or permission validation**: Access control is not checked
- **No future guarantees**: Google may change naming rules; validators reflect documented conventions at time of release
- **Not affiliated with Google**: This is an independent library

## Usage Example

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/project"

if !project.IsValidProjectID("my-project-123") {
    return fmt.Errorf("invalid project ID")
}
```

All validators return `bool`—no error messages for low friction and fast validation.
