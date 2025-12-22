package resource

import (
	"testing"

	"github.com/UnitVectorY-Labs/gcpvalidate/internal/testutil"
)

func TestIsValidVertexModelResourceName(t *testing.T) {
	testutil.RunValidatorTests(t, testutil.GetTestDataPath("vertex_model_resource_name.yaml"), IsValidVertexModelResourceName)
}

func TestIsValidProjectLocationParent(t *testing.T) {
	testutil.RunValidatorTests(t, testutil.GetTestDataPath("project_location_parent.yaml"), IsValidProjectLocationParent)
}
