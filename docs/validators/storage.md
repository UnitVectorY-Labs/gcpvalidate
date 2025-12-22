---
layout: default
title: Storage
parent: Validators
nav_order: 3
permalink: /validators/storage
---

# Storage

Validators for Google Cloud Storage identifiers.

## IsValidBucketName

**Signature**: `storage.IsValidBucketName(name string) bool`

**Example**:

```go
import "github.com/UnitVectorY-Labs/gcpvalidate/storage"

// Valid bucket names
storage.IsValidBucketName("my-bucket")              // true
storage.IsValidBucketName("bucket-123")             // true
storage.IsValidBucketName("bucket_with_underscores") // true
storage.IsValidBucketName("bucket.with.dots")       // true
storage.IsValidBucketName("my-bucket.example.com")  // true

// Invalid bucket names
storage.IsValidBucketName("My-Bucket")              // false - uppercase not allowed
storage.IsValidBucketName("-bad-start")             // false - must start with letter/number
storage.IsValidBucketName("bad-end-")               // false - must end with letter/number
storage.IsValidBucketName("ab")                     // false - too short (min 3 chars)
storage.IsValidBucketName("bucket..dots")           // false - consecutive dots not allowed
storage.IsValidBucketName("192.168.1.1")            // false - cannot look like IP address
```

**Rules**:

- Contain only lowercase letters, digits, dashes (`-`), underscores (`_`), and dots (`.`)
- Start and end with a letter or number
- Be 3 to 63 characters long
- If the name contains dots:
  - Total length can be up to 222 characters
  - Each dot-separated component can be no longer than 63 characters
- Not contain consecutive dots
- Not look like an IP address

**Note**: Bucket names are globally unique, but uniqueness cannot be validated locally. Names containing dots require domain verification.

**References**:
- [About Cloud Storage buckets](https://cloud.google.com/storage/docs/buckets)
- [Bucket naming guidelines](https://cloud.google.com/storage/docs/naming-buckets)

