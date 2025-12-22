# storage package validation rules

[← Back to index](index.md)

The `storage` package validates Google Cloud Storage identifiers.

## IsValidBucketName

**Signature**: `storage.IsValidBucketName(name string) bool`

### Rules

A Cloud Storage bucket name must:
- Contain only lowercase letters, digits, dashes (`-`), underscores (`_`), and dots (`.`). Spaces are not allowed.
- Start and end with a number or letter.
- Be 3 to 63 characters long.
- If the name contains dots:
  - Total length can be up to 222 characters, and
  - Each dot-separated component can be no longer than 63 characters.
- Not contain consecutive dots (e.g., `..` is invalid).
- Not look like an IP address (e.g., `192.168.1.1` is invalid).

### Usage example

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/storage"

if !storage.IsValidBucketName(cfg.BucketName) {
    return fmt.Errorf("invalid bucket name")
}
```

### Notes

- Bucket names are globally unique, but uniqueness cannot be validated locally.
- This validates format only, not whether the bucket exists or is available.
- Names containing dots require domain verification when used with Cloud Storage.
- Leading and trailing whitespace is rejected.
- Empty strings are invalid.

### Future extensions

A future version may include `IsValidDNSBucketName` for stricter DNS-compliant bucket names.

## References

- [About Cloud Storage buckets](https://cloud.google.com/storage/docs/buckets) - Official naming requirements
- [Bucket and Object Naming Guidelines](https://cloud.google.com/storage/docs/naming-buckets) - Detailed naming rules
- [Domain name verification](https://cloud.google.com/storage/docs/domain-name-verification) - For dotted bucket names
