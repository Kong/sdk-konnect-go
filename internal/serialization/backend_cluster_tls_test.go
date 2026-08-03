package serialization_test

import (
	"testing"

	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/stretchr/testify/require"
)

func TestBackendClusterTLSOmitsUnsetInsecureSkipVerify(t *testing.T) {
	tls := components.BackendClusterTLS{}

	require.Equal(t, map[string]any{"enabled": false}, marshalObject(t, tls))
}

func TestBackendClusterTLSIncludesExplicitInsecureSkipVerify(t *testing.T) {
	for _, enabled := range []bool{false, true} {
		t.Run(boolName(enabled), func(t *testing.T) {
			tls := components.BackendClusterTLS{InsecureSkipVerify: &enabled}

			require.Equal(t, map[string]any{
				"enabled":              false,
				"insecure_skip_verify": enabled,
			}, marshalObject(t, tls))
		})
	}
}
