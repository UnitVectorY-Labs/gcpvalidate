---
layout: default
title: gcpvalidate
nav_order: 1
permalink: /
---

# gcpvalidate Validation Rules

Client-side syntactic validation for Google Cloud resource identifiers. Validates format and conventions only—no existence checks, API calls, or IAM verification.

## Usage Example

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/project"

if !project.IsValidProjectID("my-project-123") {
    return fmt.Errorf("invalid project ID")
}
```

All validators return `bool`—no error messages for low friction and fast validation.

## Validators

| Package | Validator | Description |
|---------|-----------|-------------|
| [project](./validators/project.md) | `IsValidProjectID` | Validates GCP project identifier |
| [project](./validators/project.md) | `IsValidProjectName` | Validates GCP project display name |
| [project](./validators/project.md) | `IsValidProjectLocationParent` | Validates project/location parent path |
| [location](./validators/location.md) | `IsValidRegion` | Validates region identifier |
| [location](./validators/location.md) | `IsValidZone` | Validates zone identifier |
| [location](./validators/location.md) | `IsValidLocation` | Validates location identifier (region, zone, or global) |
| [storage](./validators/storage.md) | `IsValidBucketName` | Validates Cloud Storage bucket name |
| [vertexai](./validators/vertexai.md) | `IsValidVertexModelName` | Validates Vertex AI model display name |
| [vertexai](./validators/vertexai.md) | `IsValidVertexEndpointName` | Validates Vertex AI endpoint display name |
| [vertexai](./validators/vertexai.md) | `IsValidVertexModelResourceName` | Validates full Vertex AI model resource path |

## Important Disclaimers

- **No existence checks**: Validators confirm syntax only, not whether a resource exists
- **No API calls**: All validation is local
- **No IAM or permission validation**: Access control is not checked
- **No future guarantees**: Google may change naming rules; validators reflect documented conventions at time of release
- **Not affiliated with Google**: This is an independent library
