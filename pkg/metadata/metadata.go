package metadata

import "sync"

var (
	l         sync.RWMutex
	userAgent string
)

// SetUserAgent sets the user agent to be used by the SDK.
// This is then used by UserAgentPreRequestHook in internal/hooks/user_agent.go
func SetUserAgent(ua string) {
	l.Lock()
	defer l.Unlock()
	userAgent = ua
}

// GetUserAgent returns the user agent to be used by the SDK.
// This is then used by UserAgentPreRequestHook in internal/hooks/user_agent.go
func GetUserAgent() string {
	l.RLock()
	defer l.RUnlock()
	return userAgent
}
