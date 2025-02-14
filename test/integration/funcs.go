package integration

import (
	"strings"
	"testing"
)

// NamePrefix returns a prefix for the test name.
func NamePrefix(t *testing.T) string {
	return "sdk-konnect-go-test-integration-" + strings.ReplaceAll(t.Name(), "/", "_")
}

func Labels(t *testing.T) map[string]string {
	return map[string]string{
		"sdk-konnect-go": "true",
		"test_name":      NamePrefix(t),
		"test_run_id":    KonnectTestRunID(t),
	}
}

// Ptr returns a pointer to the given object.
func Ptr[T any](obj T) *T {
	return &obj
}
