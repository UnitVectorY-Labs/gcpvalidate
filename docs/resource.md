---
layout: default
title: Resource
nav_order: 2
permalink: /resource
---

# resource package

Validators for Google Cloud resource path structures. Validates required keywords and segment positions.

## IsValidVertexModelResourceName

**Signature**: `resource.IsValidVertexModelResourceName(path string) bool`

**Example**:

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/resource"

// Valid resource paths
resource.IsValidVertexModelResourceName(
    "projects/my-project/locations/us-central1/models/12345") // true

resource.IsValidVertexModelResourceName(
    "projects/my-project/locations/us-central1/publishers/google/models/gemini-2.0") // true

// Invalid resource paths
resource.IsValidVertexModelResourceName(
    "projects/My-Project/locations/us-central1/models/12345") // false - invalid project ID

resource.IsValidVertexModelResourceName(
    "projects/my-project/locations/us-central1/models/") // false - empty model ID

resource.IsValidVertexModelResourceName(
    "projects/my-project/models/12345") // false - missing location segment
```

**Accepted structures**:

1. Model Registry: `projects/{project}/locations/{location}/models/{modelId}`
2. Publisher model: `projects/{project}/locations/{location}/publishers/{publisher}/models/{modelId}`

**Segment validation**:
- `{project}` - Must satisfy `project.IsValidProjectID`
- `{location}` - Must satisfy `location.IsValidLocation`
- `{publisher}` - Non-empty, no slashes, no whitespace
- `{modelId}` - Non-empty, no slashes, no whitespace

**Note**: Model ID formats vary. This function validates structure but does not enforce strict character sets for `{modelId}` or `{publisher}` beyond rejecting slashes and whitespace.

## IsValidProjectLocationParent

**Signature**: `resource.IsValidProjectLocationParent(parent string) bool`

**Example**:

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/resource"

// Valid parent paths
resource.IsValidProjectLocationParent(
    "projects/my-project/locations/us-central1") // true

resource.IsValidProjectLocationParent(
    "projects/example-project/locations/global") // true

// Invalid parent paths
resource.IsValidProjectLocationParent(
    "projects/My-Project/locations/us-central1") // false - invalid project ID

resource.IsValidProjectLocationParent(
    "projects/my-project") // false - missing location
```

**Structure**: `projects/{project}/locations/{location}`

**Segment validation**:
- `{project}` - Must satisfy `project.IsValidProjectID`
- `{location}` - Must satisfy `location.IsValidLocation`

**References**:
- [Vertex AI: Locations](https://cloud.google.com/vertex-ai/docs/general/locations)
- [Google APIs: Resource names](https://cloud.google.com/apis/design/resource_names)

