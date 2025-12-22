---
layout: default
title: Vertex AI
nav_order: 2
permalink: /vertexai
---

# vertexai package

Validators for Vertex AI resource identifiers.

## IsValidVertexModelName

**Signature**: `vertexai.IsValidVertexModelName(name string) bool`

**Example**:

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/vertexai"

// Valid model names
vertexai.IsValidVertexModelName("MyModel")          // true
vertexai.IsValidVertexModelName("model_123")        // true
vertexai.IsValidVertexModelName("Model-Name")       // true

// Invalid model names
vertexai.IsValidVertexModelName("123model")         // false - must start with letter
vertexai.IsValidVertexModelName("-bad-model")       // false - must start with letter
vertexai.IsValidVertexModelName("model name")       // false - spaces not allowed
```

**Rules**:

- Contain only letters, numbers, dashes, and underscores
- Be case-sensitive
- Start with a letter
- Be no more than 128 characters long

**Note**: This validates the display name you assign, not server-assigned numeric IDs or publisher model IDs.

## IsValidVertexEndpointName

**Signature**: `vertexai.IsValidVertexEndpointName(name string) bool`

**Example**:

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/vertexai"

// Valid endpoint names
vertexai.IsValidVertexEndpointName("MyEndpoint")    // true
vertexai.IsValidVertexEndpointName("endpoint_123")  // true

// Invalid endpoint names
vertexai.IsValidVertexEndpointName("123endpoint")   // false - must start with letter
vertexai.IsValidVertexEndpointName("endpoint name") // false - spaces not allowed
```

**Rules**: Identical to `IsValidVertexModelName` (naming rules are consistent for models and endpoints).

**References**:
- [Vertex AI: Resource names](https://cloud.google.com/vertex-ai/docs/general/resource-naming)

