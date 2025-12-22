// Package testutil provides common testing utilities for gcpvalidate packages.
package testutil

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

// TestData represents the structure of YAML test data files.
type TestData struct {
	Valid   []string `yaml:"valid"`
	Invalid []string `yaml:"invalid"`
}

// LoadTestData loads test data from a YAML file.
func LoadTestData(t *testing.T, filename string) TestData {
	t.Helper()

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read test data file %s: %v", filename, err)
	}

	var testData TestData
	if err := yaml.Unmarshal(data, &testData); err != nil {
		t.Fatalf("Failed to parse test data file %s: %v", filename, err)
	}

	// Ensure both lists are non-empty
	if len(testData.Valid) == 0 {
		t.Fatalf("Test data file %s has no valid test cases", filename)
	}
	if len(testData.Invalid) == 0 {
		t.Fatalf("Test data file %s has no invalid test cases", filename)
	}

	return testData
}

// RunValidatorTests runs a validator function against test data.
func RunValidatorTests(t *testing.T, yamlFile string, validator func(string) bool) {
	t.Helper()

	// Get the absolute path to the test data file
	testData := LoadTestData(t, yamlFile)

	// Test valid cases
	for _, testCase := range testData.Valid {
		t.Run("valid:"+testCase, func(t *testing.T) {
			if !validator(testCase) {
				t.Errorf("Expected %q to be valid, but got false", testCase)
			}
		})
	}

	// Test invalid cases
	for _, testCase := range testData.Invalid {
		t.Run("invalid:"+testCase, func(t *testing.T) {
			if validator(testCase) {
				t.Errorf("Expected %q to be invalid, but got true", testCase)
			}
		})
	}
}

// GetTestDataPath returns the path to a testdata file relative to the caller's package.
func GetTestDataPath(filename string) string {
	return filepath.Join("testdata", filename)
}
