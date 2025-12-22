# gcpvalidate - Technical Architecture and Conventions

This document describes the technical architecture, coding conventions, and contribution guidelines for gcpvalidate. For user-facing documentation, see [README.md](README.md) and [docs/](docs/).

## Package Structure

```
gcpvalidate/
├── project/          // GCP project identifiers and metadata
├── location/         // Regions, zones, and global identifiers
├── storage/          // Cloud Storage bucket names
├── vertexai/         // Vertex AI model and endpoint names
├── resource/         // Generic resource path validation
├── internal/         // Shared helpers (not exported)
│   ├── whitespace.go // Whitespace validation helper
│   └── testutil/     // Test infrastructure
└── docs/             // User-facing documentation
```

### Package Naming Rules

- Lowercase
- Singular where possible
- Match Google product naming where reasonable
- No abbreviations unless canonical (e.g., `ai`, `iam`)
- One package per Google Cloud product or logical domain

## Public API Conventions

### Function Naming

All public validator functions must:
- Start with `IsValid`
- Include the full resource type name
- Be explicit about what is being validated

**Examples:**
- `IsValidProjectID`
- `IsValidProjectName`
- `IsValidBucketName`
- `IsValidVertexModelName`

**Avoid:**
- `ValidateX` (imperative, not predicate)
- `CheckX` (ambiguous)
- `ParseX` (implies transformation)

### Return Type

All validators **MUST** return `bool` only.

**Reasons:**
- This is a validation predicate, not a failure explanation
- Keeps API friction extremely low for guard clauses
- Prevents error string bikeshedding
- Encourages use in config parsing and validation

## Security Requirements

### Whitespace Handling

All validators **MUST** use `internal.HasTrimmedWhitespace()` to reject leading/trailing whitespace.

This catches:
- Spaces (` `)
- Tabs (`\t`)
- Newlines (`\n`, `\r`)
- Other whitespace characters

**Never** check only for space character with `s[0] == ' '`.

### Regex Safety

All validators must meet these requirements:

1. **Length checked before regex**
   - Check minimum and maximum length bounds before applying regex
   - Prevents regex engine from processing oversized inputs

2. **Bounded quantifiers**
   - Use `{min,max}` instead of `*` or `+` where possible
   - Example: `[0-9]{1,3}` instead of `[0-9]+`

3. **Anchored patterns**
   - Always use `^` and `$` to anchor entire string
   - Prevents partial matches

4. **No look-arounds**
   - Avoid `(?=)`, `(?!)`, `(?<=)`, `(?<!)`
   - These can cause performance issues

### Input Validation Order

Standard validation flow:

```go
func IsValidX(input string) bool {
    // 1. Empty string check
    if input == "" {
        return false
    }
    
    // 2. Whitespace check (catches tabs, newlines, etc.)
    if !internal.HasTrimmedWhitespace(input) {
        return false
    }
    
    // 3. Length bounds
    if len(input) < MIN || len(input) > MAX {
        return false
    }
    
    // 4. Character-specific checks (if needed)
    if input[0] == '-' {
        return false
    }
    
    // 5. Regex validation
    return regex.MatchString(input)
}
```

## Testing Strategy

### Data-Driven Testing

All validators use YAML-based test data:

**File location:** `<package>/testdata/<validator_name>.yaml`

**Schema:**
```yaml
valid:
  - test-case-1
  - test-case-2

invalid:
  - bad-case-1
  - bad-case-2
```

**Rules:**
- Only two keys: `valid` and `invalid`
- Arrays of strings only
- Both arrays must be non-empty
- No metadata or programmatic comments

### Test Execution

Use `internal/testutil.RunValidatorTests`:

```go
func TestIsValidProjectID(t *testing.T) {
    testutil.RunValidatorTests(t, 
        testutil.GetTestDataPath("project_id.yaml"), 
        IsValidProjectID)
}
```

### Test Naming

The test runner uses **index-based naming** to avoid issues with special characters:

```
TestIsValidProjectID/valid/0
TestIsValidProjectID/valid/1
TestIsValidProjectID/invalid/0
TestIsValidProjectID/invalid/1
```

This prevents:
- Slash characters causing subtest nesting
- Empty strings creating ambiguous names
- Special characters breaking test output

## Adding a New Validator

### Step 1: Choose Package

Use existing package if logical fit, or create new package following naming rules.

### Step 2: Implement Validator

```go
// Package mypackage provides validators for...
package mypackage

import (
    "regexp"
    "github.com/UnitVectorY-Labs/gcpvalidate/internal"
)

var myRegex = regexp.MustCompile(`^[a-z][a-z0-9-]*$`)

// IsValidMyResource validates...
//
// Rules:
//   - Must be 3-63 characters
//   - Start with lowercase letter
//   - Contain only lowercase letters, digits, hyphens
func IsValidMyResource(name string) bool {
    if name == "" {
        return false
    }
    
    if !internal.HasTrimmedWhitespace(name) {
        return false
    }
    
    if len(name) < 3 || len(name) > 63 {
        return false
    }
    
    return myRegex.MatchString(name)
}
```

### Step 3: Create Test Data

Create `mypackage/testdata/my_resource.yaml`:

```yaml
valid:
  - valid-name-1
  - valid-name-2

invalid:
  - Invalid-Name
  - -bad-start
  - ""
  - "name\n"
  - "name\t"
```

### Step 4: Add Test

Create `mypackage/mypackage_test.go`:

```go
package mypackage

import (
    "testing"
    "github.com/UnitVectorY-Labs/gcpvalidate/internal/testutil"
)

func TestIsValidMyResource(t *testing.T) {
    testutil.RunValidatorTests(t, 
        testutil.GetTestDataPath("my_resource.yaml"), 
        IsValidMyResource)
}
```

### Step 5: Update Documentation

1. Create `docs/mypackage.md` with:
   - Function signature
   - Validation rules
   - Code examples (both valid and invalid)
   - Links to Google Cloud documentation

2. Add entry to `docs/README.md` table

3. Run tests: `go test ./mypackage/...`

## Documentation Conventions

### docs/README.md

Contains a **table** listing all packages and validators with brief descriptions. This is the overview/index page.

### docs/<package>.md

Each package documentation page must include:

1. **Function signature**
2. **Code example** showing both success and failure cases
3. **Validation rules** in plain English
4. **Links to authoritative Google Cloud documentation**

**Example structure:**

```markdown
# mypackage package

## IsValidMyResource

**Signature**: `mypackage.IsValidMyResource(name string) bool`

**Example**:

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/mypackage"

// Valid
if mypackage.IsValidMyResource("my-resource") {
    // Success
}

// Invalid
if !mypackage.IsValidMyResource("Invalid-Name") {
    // Fails validation
}
```

**Rules**:
- Must be 3-63 characters
- Start with lowercase letter
- Contain only lowercase letters, digits, hyphens

**References**:
- [Google Cloud Resource Naming](https://cloud.google.com/...)
```

### Style Guidelines

Documentation should be:
- **Information dense** - No fluff, only critical details
- **Concise** - Short sentences, bullet points
- **Factual** - State what is validated, not what is "obvious"
- **Example-driven** - Show, don't just tell

## Build and Test Commands

```bash
# Download dependencies
go mod download

# Build all packages
go build -v ./...

# Run all tests
go test ./...

# Run tests with race detection and coverage
go test -race -coverprofile=coverage.txt -covermode=atomic ./...

# Run tests for specific package
go test ./project/...
```

## Code Review Checklist

Before submitting:

- [ ] Validator function named `IsValidX`
- [ ] Returns `bool` only
- [ ] Uses `internal.HasTrimmedWhitespace()` for whitespace checking
- [ ] Length checked before regex
- [ ] Regex uses bounded quantifiers where possible
- [ ] YAML test data created with valid and invalid cases
- [ ] Test includes cases for whitespace (tabs, newlines)
- [ ] Documentation updated in `docs/<package>.md`
- [ ] Entry added to `docs/README.md` table
- [ ] All tests pass: `go test ./...`

## Stability Promise

- Validators reflect documented Google Cloud conventions at time of release
- Breaking changes require major version bump
- Google may change naming rules - library will be updated accordingly
