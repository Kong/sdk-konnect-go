package hooks

import (
	"net/http"
	"strings"
)

type APIURLRequestHook struct{}

var _ beforeRequestHook = (*APIURLRequestHook)(nil)

func (i *APIURLRequestHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	switch hookCtx.OperationID {
	case "get-organizations-me":
		// NOTE(pmalek): This is because we merge OpenAPI specs and /organizations/me
		// is only served by the global API.
		// @mheap mentioned that we can add operation specific URLs to do away with this.
		if strings.HasSuffix(req.URL.Host, "api.konghq.tech") {
			req.URL.Host = "global.api.konghq.tech"
		} else {
			req.URL.Host = "global.api.konghq.com"
		}
	}
	return req, nil
}
