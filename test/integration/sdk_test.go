package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReplaceFirstSegmentWithGlobal(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "valid URL with https",
			input:       "https://example.api.konghq.com",
			expected:    "https://global.api.konghq.com",
			expectError: false,
		},
		{
			name:        "valid URL with http",
			input:       "http://example.api.konghq.com",
			expected:    "http://global.api.konghq.com",
			expectError: false,
		},
		{
			name:        "valid URL without protocol",
			input:       "example.api.konghq.com",
			expected:    "global.api.konghq.com",
			expectError: false,
		},
		{
			name:        "empty URL",
			input:       "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid URL format",
			input:       "invalid-url",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := replaceFirstSegmentWithGlobal(tt.input)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, result)
			}
		})
	}
}
