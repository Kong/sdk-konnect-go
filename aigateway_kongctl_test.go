package sdkkonnectgo_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	sdk "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
)

type customPolicyClient struct {
	t       *testing.T
	methods []string
}

func (c *customPolicyClient) Do(req *http.Request) (*http.Response, error) {
	c.methods = append(c.methods, req.Method)
	if !strings.HasPrefix(req.URL.Path, "/v1/ai-gateways/gateway/custom-policies") {
		c.t.Fatalf("unexpected custom-policy path: %s", req.URL.Path)
	}
	status := http.StatusOK
	body := `{"type":"installed","name":"custom","display_name":"Custom","schema":"return {}","id":"custom","created_at":"2026-09-25T00:00:00Z","updated_at":"2026-09-25T00:00:00Z"}`
	if req.Method == http.MethodPost {
		status = http.StatusCreated
	}
	if req.Method == http.MethodDelete {
		status = http.StatusNoContent
		body = ""
	}
	if req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/custom-policies") {
		body = `{"data":[],"meta":{"page":{"next":null,"size":0}}}`
	}
	return &http.Response{StatusCode: status, Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(body)), Request: req}, nil
}

func TestKongctlCustomPolicyCRUD(t *testing.T) {
	client := &customPolicyClient{t: t}
	api := sdk.New(sdk.WithServerURL("https://example.invalid"), sdk.WithClient(client)).AIGatewayCustomPolicies
	ctx := t.Context()
	if _, err := api.ListAiGatewayCustomPolicies(ctx, operations.ListAiGatewayCustomPoliciesRequest{GatewayID: "gateway"}); err != nil {
		t.Fatal(err)
	}
	request := components.CreateCreateAIGatewayCustomPolicyRequestInstalled(components.CreateAIGatewayCustomPolicyInstalledRequest{
		Name: "custom", DisplayName: "Custom", Type: components.CreateAIGatewayCustomPolicyInstalledRequestTypeInstalled, Schema: "return {}",
	})
	if _, err := api.CreateAiGatewayCustomPolicy(ctx, "gateway", request); err != nil {
		t.Fatal(err)
	}
	if _, err := api.GetAiGatewayCustomPolicy(ctx, "gateway", "custom"); err != nil {
		t.Fatal(err)
	}
	update := components.CreateUpdateAIGatewayCustomPolicyRequestInstalled(components.UpdateAIGatewayCustomPolicyInstalledRequest{
		Name: "custom", DisplayName: "Custom", Type: components.UpdateAIGatewayCustomPolicyInstalledRequestTypeInstalled, Schema: "return {}",
	})
	if _, err := api.UpdateAiGatewayCustomPolicy(ctx, operations.UpdateAiGatewayCustomPolicyRequest{
		GatewayID: "gateway", CustomPolicyIDOrName: "custom", UpdateAIGatewayCustomPolicyRequest: update,
	}); err != nil {
		t.Fatal(err)
	}
	if _, err := api.DeleteAiGatewayCustomPolicy(ctx, "gateway", "custom"); err != nil {
		t.Fatal(err)
	}
	if got := strings.Join(client.methods, ","); got != "GET,POST,GET,PUT,DELETE" {
		t.Fatalf("unexpected CRUD methods: %s", got)
	}
}
