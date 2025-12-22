# AGENTS.md

Guidance for AI agents contributing to this repository. Keep changes small, testable, and consistent with existing patterns.

## Project scope

- This repository is a Go module that provides small, pure validation helpers.
- Production code should stay dependency-free beyond the Go standard library.
- All external dependencies must be test-only unless there is a strong, reviewed reason.

## Repository layout

- Each domain lives in its own package under the module root (for example `project`, `location`, `storage`, `vertexai`, `resource`).
- Shared helpers live in `internal/` and are not exported.

## Go version and module rules

- Respect the Go version declared in `go.mod`.
- Do not add new module dependencies lightly.
- If you must add a dependency, prefer test-only and justify it in the PR description.

## Public API conventions

- Public validators are predicates and must be named `IsValidX`.
- Public validators must return `bool` only.
- Validators must be deterministic, side-effect free, and must not perform any network calls.

## Implementation conventions

### Validation flow

Follow the established structure across packages (order may vary based on efficiency):

1. Fast rejects first (empty strings, obvious length violations).
2. Whitespace check via `internal.HasTrimmedWhitespace`.
3. Cheap structural checks (character validation, position checks).
4. Regex last, when needed.

Note: The exact ordering between steps 1 and 2 may vary. The key is to fail fast on obviously invalid input before expensive operations.

### Whitespace and path safety

- Use `internal.HasTrimmedWhitespace(s)` to reject leading and trailing whitespace.
- For resource path segments, use `internal.IsValidPathSegment(s)` rather than re-implementing checks.

### Regex guidelines

- Compile regexes once at package scope with `regexp.MustCompile`.
- Anchor patterns with `^` and `$`.
- Prefer bounded quantifiers where practical.
- Avoid unnecessary complexity.

### Style

- Keep functions small and readable.
- Use clear names, minimal comments, and follow existing package patterns.
- Avoid introducing new exported identifiers unless required.

## Testing requirements

### Data-driven tests

- Each validator must have YAML-based test data in `<package>/testdata/`.
- File naming should match the validator subject (snake case).
- YAML schema must be exactly:

```yaml
valid:
  - example

invalid:
  - example
```

- Only the keys `valid` and `invalid`.
- Arrays of strings only.
- Both arrays must be non-empty.

### Test runner

- Use `internal/testutil.RunValidatorTests` for validator tests.
- Tests should be named `TestIsValidX` and should load YAML via `testutil.GetTestDataPath`.

### Coverage expectations

- New validators must include representative valid and invalid cases.
- Include edge cases for whitespace handling and length bounds.

## Adding a new validator

1. Pick the correct package, or create a new one using existing naming patterns.
2. Implement `IsValidX(string) bool` following the validation flow above.
3. Add `testdata/<name>.yaml` with a non-empty `valid` and `invalid` list.
4. Add a `*_test.go` test using `testutil.RunValidatorTests`.
5. Run the full test suite before submitting.

## Local commands

```bash
go mod download
go build ./...
go test ./...
go test -race -coverprofile=coverage.txt -covermode=atomic ./...
```

## Change discipline

- Prefer small commits and focused PRs.
- Do not mix refactors with feature changes unless necessary.
- Keep formatting and lint noise out of functional changes.

## Review checklist

- [ ] Public function name is `IsValidX`
- [ ] Return type is `bool`
- [ ] Uses `internal.HasTrimmedWhitespace` where applicable
- [ ] Length bounds checked early
- [ ] Regex compiled once, anchored, and used last
- [ ] YAML test data added with non-empty `valid` and `invalid`
- [ ] Test uses `testutil.RunValidatorTests`
- [ ] `go test ./...` passes
