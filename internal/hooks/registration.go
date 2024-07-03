package hooks

import (
	"net/http"
	"strings"
	"sync"
)

/*
 * This file is only ever generated once on the first generation and then is free to be modified.
 * Any hooks you wish to add should be registered in the InitHooks function. Feel free to define them
 * in this file or in separate files in the hooks package.
 */

func initHooks(h *Hooks) {
	h.registerBeforeRequestHook(&UserAgentPreRequestHook{})
}

var (
	l                sync.RWMutex
	userAgent        string
	defaultUserAgent = "SpeakeasyGoSDK/0.1"
)

func SetUserAgent(ua string) {
	l.Lock()
	defer l.Unlock()
	userAgent = ua
}

func GetUserAgent() string {
	l.RLock()
	defer l.RUnlock()
	return userAgent
}

type UserAgentPreRequestHook struct{}

var _ beforeRequestHook = (*UserAgentPreRequestHook)(nil)

func (i *UserAgentPreRequestHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	req.Header.Set("User-Agent", GetUserAgent())

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
