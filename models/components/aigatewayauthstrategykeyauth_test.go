package components

import (
	"encoding/json"
	"testing"
)

func TestAIGatewayAuthStrategyKeyAuthIdentityRealmsJSON(t *testing.T) {
	region := "us"

	t.Run("request", func(t *testing.T) {
		tests := []struct {
			name    string
			realms  []AIGatewayAuthStrategyKeyAuthIdentityRealms
			present bool
			length  int
		}{
			{name: "omitted", present: false},
			{name: "empty", realms: []AIGatewayAuthStrategyKeyAuthIdentityRealms{}, present: true},
			{name: "populated", realms: []AIGatewayAuthStrategyKeyAuthIdentityRealms{{Region: &region}}, present: true, length: 1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				assertIdentityRealmsJSON(t, AIGatewayAuthStrategyKeyAuthConfig{IdentityRealms: tt.realms}, tt.present, tt.length)
			})
		}
	})

	t.Run("response", func(t *testing.T) {
		tests := []struct {
			name    string
			realms  []IdentityRealms
			present bool
			length  int
		}{
			{name: "omitted", present: false},
			{name: "empty", realms: []IdentityRealms{}, present: true},
			{name: "populated", realms: []IdentityRealms{{Region: &region}}, present: true, length: 1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				assertIdentityRealmsJSON(t, AIGatewayAuthStrategyKeyAuthResponseConfig{IdentityRealms: tt.realms}, tt.present, tt.length)
			})
		}
	})
}

func assertIdentityRealmsJSON(t *testing.T, config any, present bool, length int) {
	t.Helper()

	body, err := json.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}

	var got map[string]json.RawMessage
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatal(err)
	}

	raw, ok := got["identity_realms"]
	if ok != present {
		t.Fatalf("identity_realms present = %t, want %t; body: %s", ok, present, body)
	}
	if !ok {
		return
	}
	if string(raw) == "null" {
		t.Fatalf("identity_realms must be an array, got null; body: %s", body)
	}

	var realms []json.RawMessage
	if err := json.Unmarshal(raw, &realms); err != nil {
		t.Fatal(err)
	}
	if len(realms) != length {
		t.Fatalf("identity_realms length = %d, want %d; body: %s", len(realms), length, body)
	}
}
