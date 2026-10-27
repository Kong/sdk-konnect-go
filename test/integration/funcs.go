package integration

import (
	"strings"
	"testing"
)

// workspace is the Konnect core-entities workspace targeted by the workspace-scoped endpoints.
const workspace = "default"

// NamePrefix returns a prefix for the test name.
func NamePrefix(t *testing.T) string {
	ret := "sdk-konnect-go-test-integration-" + strings.ReplaceAll(t.Name(), "/", "_")
	if len(ret) > 63 {
		ret = ret[:63]
	}
	return strings.TrimSuffix(ret, "_")
}

func Labels(t *testing.T) map[string]string {
	return map[string]string{
		"sdk-konnect-go": "true",
		"test_name":      NamePrefix(t),
		"test_run_id":    KonnectTestRunID(t),
	}
}
