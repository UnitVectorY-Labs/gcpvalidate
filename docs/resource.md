# resource package validation rules

[← Back to index](index.md)

The `resource` package validates the structure of common Google Cloud resource paths. It validates required keywords and segment positions and delegates segment validation to other packages.

## IsValidVertexModelResourceName

**Signature**: `resource.IsValidVertexModelResourceName(path string) bool`

### Accepted path structures

At minimum, the validator recognizes these Vertex AI model resource name forms:

1. **Model Registry model resource**:
   - `projects/{project}/locations/{location}/models/{modelId}`

2. **Publisher model resource**:
   - `projects/{project}/locations/{location}/publishers/{publisher}/models/{modelId}`

The Vertex AI locations documentation shows the canonical `projects/PROJECT/locations/us-central1/...` resource style.

### Segment rules

- `{project}` must satisfy `project.IsValidProjectID`.
- `{location}` must satisfy `location.IsValidLocation` (syntax-based).
- `{publisher}` must be non-empty and contain no slashes.
- `{modelId}` must be non-empty and contain no slashes.

### Usage example

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/resource"

if !resource.IsValidVertexModelResourceName(cfg.ModelPath) {
    return fmt.Errorf("invalid model resource name")
}
```

### Notes

- `{modelId}` formats vary. This function does not enforce a narrow regex for `{modelId}`.
- This validates structure only, not whether the resource exists.
- Leading and trailing whitespace is rejected.
- Empty strings are invalid.

## IsValidProjectLocationParent

**Signature**: `resource.IsValidProjectLocationParent(parent string) bool`

### Accepted structure

- `projects/{project}/locations/{location}`

This structure is widely used in Google APIs and commonly documented as the parent format.

### Segment rules

- `{project}` must satisfy `project.IsValidProjectID`.
- `{location}` must satisfy `location.IsValidLocation`.

### Usage example

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/resource"

if !resource.IsValidProjectLocationParent(cfg.Parent) {
    return fmt.Errorf("invalid project location parent")
}
```

### Notes

- This validates structure only, not whether the resource exists.
- Leading and trailing whitespace is rejected.
- Empty strings are invalid.

## References

- [Vertex AI: Locations](https://cloud.google.com/vertex-ai/docs/general/locations) - Shows `projects/PROJECT/locations/us-central1/datasets/...` usage
- [Cloud Run API](https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services/list) - Documents `projects/{project}/locations/{location}` format
- [Google APIs: Resource names](https://cloud.google.com/apis/design/resource_names) - General resource naming design
- [Resource Manager: Project ID rules](https://cloud.google.com/resource-manager/reference/rest/v1beta1/projects) - For `{project}` segment validation
