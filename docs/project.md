# project package validation rules

[← Back to index](index.md)

The `project` package validates Google Cloud Project identifiers.

## IsValidProjectID

**Signature**: `project.IsValidProjectID(id string) bool`

### Rules

A Project ID must:
- Be 6 to 30 characters long.
- Start with a lowercase letter.
- Contain only ASCII lowercase letters, digits, and hyphens.
- Not end with a hyphen (trailing hyphens prohibited).

### Usage example

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/project"

if !project.IsValidProjectID(cfg.ProjectID) {
    return fmt.Errorf("invalid project id")
}
```

### Notes

- This is a syntactic check only. It does not verify the project exists or is accessible.
- Leading and trailing whitespace is rejected.
- Empty strings are invalid.

## IsValidProjectName

**Signature**: `project.IsValidProjectName(name string) bool`

### Rules

A Project display name must:
- Be 4 to 30 characters long.
- Contain only: letters, numbers, single quotes, hyphens, spaces, or exclamation points.

### Usage example

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/project"

if !project.IsValidProjectName(displayName) {
    return fmt.Errorf("invalid project name")
}
```

### Notes

- This validates the display name used in the console and some tooling, not the Project ID used in resource paths and most APIs.
- Leading and trailing whitespace is rejected.
- Empty strings are invalid.

## References

- [Creating and managing projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects) - Project ID requirements
- [Resource Manager REST: projects](https://cloud.google.com/resource-manager/reference/rest/v1beta1/projects) - Project ID rules, trailing hyphen prohibition
- [Organization resource management](https://cloud.google.com/resource-manager/docs/organization-resource-management) - Project name rules
