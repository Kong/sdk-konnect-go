package components

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestKongctlSparseRequests(t *testing.T) {
	for name, request := range map[string]any{
		"portal":                   UpdatePortal{},
		"webhook":                  UpdatePortalAuditLogWebhook{},
		"email":                    PatchCustomPortalEmailTemplatePayload{},
		"identity provider":        UpdateIdentityProvider{},
		"portal identity provider": PortalUpdateIdentityProvider{},
		"DCR":                      UpdateDcrConfigHTTPInRequest{},
		"portal team":              PortalUpdateTeamRequest{},
		"organization team":        UpdateTeam{},
	} {
		t.Run(name, func(t *testing.T) {
			data, err := json.Marshal(request)
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != "{}" {
				t.Fatalf("unset fields serialized: %s", data)
			}
		})
	}
	// An embedded request must not consume kongctl's sibling metadata.
	var team struct {
		CreateTeam
		Ref string `json:"ref"`
	}
	if err := json.Unmarshal([]byte(`{"name":"team","ref":"local-team"}`), &team); err != nil {
		t.Fatal(err)
	}
	if team.Ref != "local-team" || team.Name != "team" {
		t.Fatalf("team metadata or omitted membership management changed: %+v", team)
	}
}

func TestKongctlAIGateway22Surfaces(t *testing.T) {
	for name, value := range map[string]any{
		"passthrough": AIGatewayModelFormat{Type: AIGatewayModelFormatTypePassthrough.ToPointer()},
		"skills":      CapabilitiesSkills,
		"decisions":   AIGatewayModelModelCapabilitiesDecisions,
		"aliases":     AIGatewayModelSelectorConfig{Values: []string{"primary", "alias"}},
		"target":      CreateAIGatewayTargetConfigTypesafe(AIGatewayTargetTypesafeConfig{}),
	} {
		t.Run(name, func(t *testing.T) {
			data, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			roundTrip := reflect.New(reflect.TypeOf(value))
			if err := json.Unmarshal(data, roundTrip.Interface()); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(value, roundTrip.Elem().Interface()) {
				t.Fatalf("AI Gateway 2.2 value changed during round trip: %s", data)
			}
		})
	}
	for _, data := range []string{
		`{"type":"typesafe","name":"jev","display_name":"Jev","config":{"auth":{"type":"basic"}}}`,
	} {
		var provider CreateAIGatewayModelProviderRequest
		if err := json.Unmarshal([]byte(data), &provider); err != nil {
			t.Fatal(err)
		}
		if provider.AIGatewayModelProviderTypesafe == nil {
			t.Fatal("typesafe provider lost")
		}
	}
	for _, data := range []string{
		`{"type":"installed","name":"custom","display_name":"Custom","schema":"return {}"}`,
		`{"type":"streaming","name":"custom","display_name":"Custom","schema":"return {}","handler":"return {}"}`,
	} {
		var policy CreateAIGatewayCustomPolicyRequest
		if err := json.Unmarshal([]byte(data), &policy); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(policy)
		if err != nil {
			t.Fatal(err)
		}
		var update UpdateAIGatewayCustomPolicyRequest
		if err := json.Unmarshal(encoded, &update); err != nil {
			t.Fatal(err)
		}
	}
}
