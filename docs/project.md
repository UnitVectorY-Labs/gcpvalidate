---
layout: default
title: Project
nav_order: 2
permalink: /project
---

# project package

Validators for Google Cloud Project identifiers.

## IsValidProjectID

**Signature**: `project.IsValidProjectID(id string) bool`

**Example**:

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/project"

// Valid project IDs
project.IsValidProjectID("my-project-123")  // true
project.IsValidProjectID("example-gcp")     // true

// Invalid project IDs
project.IsValidProjectID("My-Project")      // false - uppercase not allowed
project.IsValidProjectID("-bad-start")      // false - cannot start with hyphen
project.IsValidProjectID("bad-end-")        // false - cannot end with hyphen
project.IsValidProjectID("short")           // false - too short (min 6 chars)
```

**Rules**:

- Be 6 to 30 characters long
- Start with a lowercase letter
- Contain only lowercase letters, digits, and hyphens
- Not end with a hyphen

**References**:
- [Creating and managing projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects)
- [Resource Manager REST: projects](https://cloud.google.com/resource-manager/reference/rest/v1beta1/projects)

---

## IsValidProjectName

**Signature**: `project.IsValidProjectName(name string) bool`

**Example**:

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/project"

// Valid project names
project.IsValidProjectName("My Project")        // true
project.IsValidProjectName("Test 123!")         // true
project.IsValidProjectName("O'Brien's Project") // true

// Invalid project names
project.IsValidProjectName("abc")               // false - too short (min 4 chars)
project.IsValidProjectName("Way Too Long Project Name Here")  // false - exceeds 30 chars
project.IsValidProjectName("Invalid@Char")      // false - @ not allowed
```

**Rules**:

- Be 4 to 30 characters long
- Contain only letters, numbers, single quotes, hyphens, spaces, or exclamation points

**Note**: This validates the display name shown in the console, not the Project ID used in APIs and resource paths.

## IsValidProjectLocationParent

**Signature**: `project.IsValidProjectLocationParent(parent string) bool`

**Example**:

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/project"

// Valid parent paths
project.IsValidProjectLocationParent(
    "projects/my-project/locations/us-central1") // true

project.IsValidProjectLocationParent(
    "projects/example-project/locations/global") // true

// Invalid parent paths
project.IsValidProjectLocationParent(
    "projects/My-Project/locations/us-central1") // false - invalid project ID

project.IsValidProjectLocationParent(
    "projects/my-project") // false - missing location
```

**Structure**: `projects/{project}/locations/{location}`

**Segment validation**:
- `{project}` - Must satisfy `project.IsValidProjectID`
- `{location}` - Must satisfy `location.IsValidLocation`

**References**:
- [Organization resource management](https://cloud.google.com/resource-manager/docs/organization-resource-management)
- [Google APIs: Resource names](https://cloud.google.com/apis/design/resource_names)

