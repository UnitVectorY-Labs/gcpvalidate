# gcpvalidate validation rules

This documentation describes the client-side validation rules implemented by gcpvalidate. These checks validate format and conventions only, not existence, permissions, or availability.

## Important disclaimers

- **No existence checks**: Validators confirm syntax only, not whether a resource exists.
- **No API calls**: All validation is local.
- **No IAM or permission validation**: Access control is not checked.
- **No future guarantees**: Google may change naming rules; validators reflect documented conventions at time of release.
- **Not affiliated with Google**: This is an independent library.

## Packages documented here

- [project](project.md) - GCP project identifiers and metadata
- [location](location.md) - Regions, zones, and global identifiers
- [storage](storage.md) - Cloud Storage bucket names
- [vertexai](vertexai.md) - Vertex AI model and endpoint names
- [resource](resource.md) - Generic resource path validation

## Design principles

- **Boolean return semantics**: All validators return `bool` (no error explanations)
- **Security first**: Length checked before regex, no unbounded repetition
- **Predictable**: Follows documented Google Cloud naming conventions
- **Boring**: Standard library only, no runtime dependencies beyond `gopkg.in/yaml.v3` for testing
