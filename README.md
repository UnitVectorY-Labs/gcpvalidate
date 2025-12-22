# gcpvalidate

Client-side syntactic validation for Google Cloud resource identifiers.

## Design Principles

### Goals

- Provide client-side syntactic validation of common Google Cloud identifiers
- Be idiomatic Go, boring, predictable, and easy to adopt
- Avoid runtime dependencies beyond the Go standard library
- Be safe against malformed or adversarial inputs
- Make expansion easy without breaking APIs

### Explicit Non-Goals

- **No existence checks** - Does not verify resources exist
- **No API calls** - All validation is local
- **No IAM or permission validation** - Access control is not checked
- **No future guarantees** - Google may change naming rules

### Stability Promise

- Validators reflect documented conventions at time of release
- Changes in Google docs may require new versions of this library
- Backward-incompatible changes require a major version bump

## Package Layout

```
gcpvalidate/
├── project/          // GCP project identifiers and metadata
├── location/         // Regions, zones, and global identifiers
├── storage/          // Cloud Storage
├── vertexai/         // Vertex AI
├── resource/         // Generic resource path validation helpers
├── internal/         // Shared helpers (not exported)
└── docs/             // Rendered documentation site
```

### Package Naming Rules

- Lowercase
- Singular where possible
- Match Google product naming where reasonable
- No abbreviations unless canonical (e.g. ai, iam)

## Public API Design

### Boolean Return Semantics

All public validators return `bool` (not errors).

**Reasons:**
- This is a validation predicate, not a failure explanation
- Keeps API friction extremely low
- Prevents error string bikeshedding
- Encourages use in guards, flags, config parsing

### Naming Conventions

Public functions must:
- Start with `IsValid`
- Include the full resource name
- Be explicit about what is being validated

**Examples:**
- `IsValidProjectID`
- `IsValidProjectName`
- `IsValidBucketName`
- `IsValidVertexModelName`

**Avoid:**
- `ValidateX`
- `CheckX`
- `ParseX`

### Usage Example

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/project"

if !project.IsValidProjectID(cfg.ProjectID) {
    return fmt.Errorf("invalid project id")
}
```

## Security and Performance

### Regex Safety Guarantees

All validators meet these requirements:

1. **Length checked before regex** - Reject immediately if outside documented bounds
2. **No unbounded repetition** - Avoid `.*`, `.+`, or catastrophic backtracking patterns
3. **Regexes must be anchored** - Always `^...$`
4. **No look-arounds** - No `(?=)`, `(?!`, etc.
5. **Prefer rune iteration over regex where possible**

### Input Handling Rules

- **Leading and trailing whitespace** - MUST be rejected unless explicitly allowed by Google docs
- **Unicode** - Unless explicitly documented, validators assume ASCII
- **Empty strings** - Always invalid

### Complexity Constraints

- Validation must be O(n)
- No allocations proportional to regex complexity
- No recursion

## Testing Strategy

### Philosophy

- Fully data-driven
- No hard-coded test cases in Go files
- Easy to add new examples without writing code

### YAML Test Files

Test data is stored in YAML files under each package's `testdata/` directory.

**Test data layout:**

```
<package>/
├── testdata/
│   ├── <validator_name>.yaml
```

**YAML schema:**

```yaml
valid:
  - my-bucket
  - example.bucket.name

invalid:
  - MyBucket
  - -bad
  - bad-
```

**Rules:**
- Only two keys: `valid`, `invalid`
- Arrays of strings only
- No metadata, no comments relied upon programmatically

### Test Execution Model

For each public validator:
- Exactly one YAML file
- Test runner:
  - Loads YAML
  - Iterates `valid` → expect `true`
  - Iterates `invalid` → expect `false`
- Tests must fail if:
  - YAML missing
  - YAML malformed
  - Either list empty

## Adding a New Validator

1. **Choose the package** - Use existing package or create new one following naming rules
2. **Implement the validator** - Follow the public API naming convention (`IsValidX`)
3. **Create test data** - Add a YAML file in `<package>/testdata/`
4. **Add test** - Create a test function that calls `testutil.RunValidatorTests`
5. **Document** - Add to appropriate documentation page in `docs/`
6. **Security review** - Ensure regex safety guarantees are met

## Documentation

- **User documentation** - See [docs/](docs/) for end-user documentation
- **Architecture** - This README serves as the architecture and contribution guide

## Building and Testing

```bash
# Download dependencies
go mod download

# Build all packages
go build -v ./...

# Run all tests
go test -v ./...

# Run tests with coverage
go test -race -coverprofile=coverage.txt -covermode=atomic ./...
```

## Contributing

Contributions are welcome! Please ensure:

1. All validators follow the security and performance requirements
2. Test data is provided in YAML format
3. Documentation is updated
4. All tests pass

## License

See [LICENSE](LICENSE) for details.

## Disclaimer

This library is not affiliated with Google. It validates format and conventions only, based on publicly available Google Cloud documentation.

