package project

import (
	"testing"

	"github.com/UnitVectorY-Labs/gcpvalidate/internal/testutil"
)

func TestIsValidProjectID(t *testing.T) {
	testutil.RunValidatorTests(t, testutil.GetTestDataPath("project_id.yaml"), IsValidProjectID)
}

func TestIsValidProjectName(t *testing.T) {
	testutil.RunValidatorTests(t, testutil.GetTestDataPath("project_name.yaml"), IsValidProjectName)
}
