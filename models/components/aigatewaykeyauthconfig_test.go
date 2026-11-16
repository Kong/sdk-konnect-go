package components

import (
	"encoding/json"
	"testing"
)

func TestAIGatewayKeyAuthConfigIdentityRealmsJSON(t *testing.T) {
	strategies := []struct {
		name   string
		config func([]byte) (any, error)
	}{
		{name: "request", config: func(data []byte) (any, error) {
			var strategy AIGatewayAuthStrategyKeyAuth
			err := json.Unmarshal(data, &strategy.Config)
			return strategy.Config, err
		}},
		{name: "response", config: func(data []byte) (any, error) {
			var strategy AIGatewayAuthStrategyKeyAuthResponse
			err := json.Unmarshal(data, &strategy.Config)
			return strategy.Config, err
		}},
	}
	tests := []struct {
		name    string
		config  string
		present bool
		length  int
	}{
		{name: "omitted", config: `{}`, present: false},
		{name: "empty", config: `{"identity_realms":[]}`, present: true},
		{name: "populated", config: `{"identity_realms":[{"region":"us"}]}`, present: true, length: 1},
	}

	for _, strategy := range strategies {
		t.Run(strategy.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					config, err := strategy.config([]byte(tt.config))
					if err != nil {
						t.Fatal(err)
					}
					body, err := json.Marshal(config)
					if err != nil {
						t.Fatal(err)
					}

					var got map[string]json.RawMessage
					if err := json.Unmarshal(body, &got); err != nil {
						t.Fatal(err)
					}

					raw, present := got["identity_realms"]
					if present != tt.present {
						t.Fatalf("identity_realms present = %t, want %t; body: %s", present, tt.present, body)
					}
					if !present {
						return
					}
					if string(raw) == "null" {
						t.Fatalf("identity_realms must be an array, got null; body: %s", body)
					}

					var realms []json.RawMessage
					if err := json.Unmarshal(raw, &realms); err != nil {
						t.Fatal(err)
					}
					if len(realms) != tt.length {
						t.Fatalf("identity_realms length = %d, want %d; body: %s", len(realms), tt.length, body)
					}
				})
			}
		})
	}
}
