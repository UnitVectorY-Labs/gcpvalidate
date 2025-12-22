package location

import (
	"testing"

	"github.com/UnitVectorY-Labs/gcpvalidate/internal/testutil"
)

func TestIsValidRegion(t *testing.T) {
	testutil.RunValidatorTests(t, testutil.GetTestDataPath("region.yaml"), IsValidRegion)
}

func TestIsValidZone(t *testing.T) {
	testutil.RunValidatorTests(t, testutil.GetTestDataPath("zone.yaml"), IsValidZone)
}

func TestIsValidLocation(t *testing.T) {
	testutil.RunValidatorTests(t, testutil.GetTestDataPath("location.yaml"), IsValidLocation)
}
