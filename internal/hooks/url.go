package hooks

import (
	"net/http"
	"strings"
)

type APIURLRequestHook struct {
	CustomDomain string
}

var _ beforeRequestHook = (*APIURLRequestHook)(nil)

// Replace `.konghq.com` with the custom domain set in the `KONG_CUSTOM_DOMAIN` environment variable.
func (i *APIURLRequestHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if i.CustomDomain != "" {
		req.URL.Host = strings.Replace(req.URL.Host, ".konghq.com", "."+i.CustomDomain, 1)
		req.Host = strings.Replace(req.Host, ".konghq.com", "."+i.CustomDomain, 1)
	}

	return req, nil
}
