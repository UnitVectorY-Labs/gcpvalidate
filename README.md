# gcpvalidate

Client-side syntactic validation for Google Cloud resource identifiers in Go.

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/UnitVectorY-Labs/gcpvalidate/project"
)

func main() {
    projectID := "my-gcp-project"
    
    if project.IsValidProjectID(projectID) {
        fmt.Println("Valid project ID")
    } else {
        fmt.Println("Invalid project ID")
    }
}
```

## What It Does

Validates the **format** of Google Cloud identifiers like project IDs, bucket names, and resource paths. All validation is local—no API calls, no existence checks, no IAM verification.

```go
//  Valid project IDs
project.IsValidProjectID("my-project-123")  // true
project.IsValidProjectID("example-gcp")     // true

//  Invalid project IDs  
project.IsValidProjectID("My-Project")      // false (uppercase)
project.IsValidProjectID("-bad-start")      // false (starts with hyphen)
project.IsValidProjectID("bad-end-")        // false (ends with hyphen)
```

## Installation

```bash
go get github.com/UnitVectorY-Labs/gcpvalidate
```

## Packages

- **[project](docs/project.md)** - Project IDs and display names
- **[location](docs/location.md)** - Regions, zones, and location identifiers
- **[storage](docs/storage.md)** - Cloud Storage bucket names
- **[vertexai](docs/vertexai.md)** - Vertex AI model and endpoint names
- **[resource](docs/resource.md)** - Full resource path validation

See [complete documentation](docs/README.md) for detailed validation rules and examples.

## Design Principles

### What It Validates
-  **Syntax** - Format matches documented Google Cloud conventions
-  **Length** - Identifier length is within allowed bounds
-  **Character sets** - Only permitted characters are used

### What It Does NOT Validate
-  **Existence** - Does not check if resource exists
-  **Permissions** - No IAM or access control checks
-  **Availability** - No API calls to verify resource state

### API Design

All validators return `bool`:

```go
if !storage.IsValidBucketName(name) {
    return fmt.Errorf("invalid bucket name: %s", name)
}
```

No error messages means:
- Low friction for guard clauses and config validation
- Fast and predictable
- No error string bikeshedding

## Simplicity

- Length bounds enforced before regex evaluation
- Anchored regex patterns (`^...$`) with explicit length checks
- Uses Go's RE2 engine which avoids catastrophic backtracking
- Whitespace (including tabs, newlines) rejected unless explicitly allowed

## Disclaimer

Not affiliated with Google. Validates format based on publicly documented conventions. Google may change naming rules—this library reflects conventions at time of release.
