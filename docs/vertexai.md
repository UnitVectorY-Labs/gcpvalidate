# vertexai package validation rules

[← Back to index](index.md)

The `vertexai` package validates Vertex AI resource identifiers.

Vertex AI uses the same naming rules for model names and endpoint names (these are display names you provide when creating resources).

## IsValidVertexModelName

**Signature**: `vertexai.IsValidVertexModelName(name string) bool`

### Rules

A Vertex AI model name must:
- Contain only letters, numbers, dashes, and underscores.
- Be case-sensitive.
- Start with a letter.
- Be no more than 128 characters long.

### Usage example

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/vertexai"

if !vertexai.IsValidVertexModelName(cfg.ModelName) {
    return fmt.Errorf("invalid model name")
}
```

### Notes

- This function validates the name you assign (display name).
- It does NOT validate:
  - Server-assigned numeric IDs for Model Registry resources
  - Publisher model IDs or versioned model identifiers
- Leading and trailing whitespace is rejected.
- Empty strings are invalid.

## IsValidVertexEndpointName

**Signature**: `vertexai.IsValidVertexEndpointName(name string) bool`

### Rules

Identical to `IsValidVertexModelName`, because the Vertex AI documentation states the naming rules are consistent for both models and endpoints.

A Vertex AI endpoint name must:
- Contain only letters, numbers, dashes, and underscores.
- Be case-sensitive.
- Start with a letter.
- Be no more than 128 characters long.

### Usage example

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/vertexai"

if !vertexai.IsValidVertexEndpointName(cfg.EndpointName) {
    return fmt.Errorf("invalid endpoint name")
}
```

### Notes

- This function validates the name you assign (display name).
- Leading and trailing whitespace is rejected.
- Empty strings are invalid.

## References

- [Vertex AI: Resource names](https://cloud.google.com/vertex-ai/docs/general/resource-naming) - Naming rules for Vertex AI resources
