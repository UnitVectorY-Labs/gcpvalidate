package vertexai

import (
	"testing"

	"github.com/UnitVectorY-Labs/gcpvalidate/internal/testutil"
)

func TestIsValidVertexModelName(t *testing.T) {
	testutil.RunValidatorTests(t, testutil.GetTestDataPath("vertex_model_name.yaml"), IsValidVertexModelName)
}

func TestIsValidVertexEndpointName(t *testing.T) {
	testutil.RunValidatorTests(t, testutil.GetTestDataPath("vertex_endpoint_name.yaml"), IsValidVertexEndpointName)
}

func TestIsValidVertexModelResourceName(t *testing.T) {
	testutil.RunValidatorTests(t, testutil.GetTestDataPath("vertex_model_resource_name.yaml"), IsValidVertexModelResourceName)
}
