package hooks

import (
	"net/http"
	"sync"
)

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

	return req, nil
}
