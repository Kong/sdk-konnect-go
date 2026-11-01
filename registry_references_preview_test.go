package sdkkonnectgo_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	sdk "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
)

type registryReferenceCapture struct {
	path  string
	calls int
	t     *testing.T
}

func (c *registryReferenceCapture) Do(request *http.Request) (*http.Response, error) {
	c.calls++
	if request.Method != http.MethodGet || request.URL.Host != "preview.invalid" || request.URL.Path != c.path {
		c.t.Fatal("Registry read changed its configured origin, path or HTTP method")
	}
	q := request.URL.Query()
	if len(q) != 2 || q.Get("page[size]") != "1" || q.Get("page[number]") != "2" {
		c.t.Fatalf("Registry pagination parameters = %v", q)
	}
	return &http.Response{StatusCode: 200, Header: http.Header{"Content-Type": {"application/json"}},
		Body:    io.NopCloser(strings.NewReader(`{"data":[],"meta":{"page":{"number":2,"size":1,"total":0}}}`)),
		Request: request}, nil
}

func TestPreviewRegistryReferenceReadsUseConfiguredOriginAndBracketedPagination(t *testing.T) {
	for _, principals := range []bool{false, true} {
		capture := &registryReferenceCapture{path: "/v1/agents", t: t}
		client := sdk.New(sdk.WithClient(capture), sdk.WithServerURL("https://preview.invalid"))
		if principals {
			capture.path += "/731559d3-39bd-48c7-b33c-eb192847eb20/principals"
			response, err := client.RegistryAgentReferences.ListAgentPrincipals(t.Context(), operations.ListAgentPrincipalsRequest{
				AgentID: "731559d3-39bd-48c7-b33c-eb192847eb20", PageSize: new(int64(1)), PageNumber: new(int64(2)),
			}, operations.WithServerURL("https://preview.invalid"))
			if err != nil || response == nil || response.RegistryReferenceAgentPrincipalList == nil {
				t.Fatal("missing typed principal association response", err)
			}
		} else {
			response, err := client.RegistryAgentReferences.ListAgents(t.Context(), operations.ListAgentsRequest{
				PageSize: new(int64(1)), PageNumber: new(int64(2)),
			}, operations.WithServerURL("https://preview.invalid"))
			if err != nil || response == nil || response.RegistryReferenceAgentList == nil {
				t.Fatal("missing typed Agent response", err)
			}
		}
		if capture.calls != 1 {
			t.Fatalf("expected one read, got %d", capture.calls)
		}
	}
}
