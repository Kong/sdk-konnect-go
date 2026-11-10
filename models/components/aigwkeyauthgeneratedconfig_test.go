package components

import (
	"encoding/json"
	"testing"
)

func TestAIGWKeyAuthGeneratedConfigIdentityRealmsJSON(t *testing.T) {
	region := "us"
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
			body, err := json.Marshal(AIGWKeyAuthGeneratedConfig{IdentityRealms: tt.realms})
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

			var realms []json.RawMessage
			if err := json.Unmarshal(raw, &realms); err != nil {
				t.Fatal(err)
			}
			if len(realms) != tt.length {
				t.Fatalf("identity_realms length = %d, want %d; body: %s", len(realms), tt.length, body)
			}
		})
	}
}
