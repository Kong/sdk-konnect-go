package hooks

import (
	"net/http"

	"github.com/Kong/sdk-konnect-go/pkg/metadata"
)

type UserAgentPreRequestHook struct{}

var _ beforeRequestHook = (*UserAgentPreRequestHook)(nil)

func (i *UserAgentPreRequestHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	ua := metadata.GetUserAgent()
	if ua != "" {
		req.Header.Set("User-Agent", ua)
		return req, nil
	}

	return req, nil
}
