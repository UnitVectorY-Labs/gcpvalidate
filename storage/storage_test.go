package storage

import (
	"testing"

	"github.com/UnitVectorY-Labs/gcpvalidate/internal/testutil"
)

func TestIsValidBucketName(t *testing.T) {
	testutil.RunValidatorTests(t, testutil.GetTestDataPath("bucket_name.yaml"), IsValidBucketName)
}
