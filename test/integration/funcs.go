package integration

import "testing"

// NamePrefix returns a prefix for the test name.
func NamePrefix(t *testing.T) string {
	return "sdk-konnect-go-test-integration-" + t.Name()
}

func Labels(t *testing.T) map[string]string {
	return map[string]string{
		"sdk-konnect-go": "true",
		"test_name":      NamePrefix(t),
		"test_run_id":    KonnectTestRunID(t),
	}
}
