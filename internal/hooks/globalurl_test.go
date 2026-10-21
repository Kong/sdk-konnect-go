package hooks

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGlobalAPIURLRequestHook_BeforeRequest(t *testing.T) {
	const identityOp = "get-organizations-me"

	tests := []struct {
		name         string
		operationID  string
		baseURL      string
		customDomain string
		wantHost     string
		wantPath     string
	}{
		{
			name:        "default server list keeps the global host",
			operationID: identityOp,
			baseURL:     "https://global.api.konghq.com/",
			wantHost:    "global.api.konghq.com",
			wantPath:    "/v3/organizations/me",
		},
		{
			name:        "public regional host goes to the production global",
			operationID: identityOp,
			baseURL:     "https://us.api.konghq.com/",
			wantHost:    "global.api.konghq.com",
			wantPath:    "/v3/organizations/me",
		},
		{
			name:        "public dev host goes to the dev global",
			operationID: identityOp,
			baseURL:     "https://us.api.konghq.tech/",
			wantHost:    "global.api.konghq.tech",
			wantPath:    "/v3/organizations/me",
		},
		{
			name:        "explicit port on a dev host still goes to the dev global",
			operationID: identityOp,
			baseURL:     "https://us.api.konghq.tech:443/",
			wantHost:    "global.api.konghq.tech",
			wantPath:    "/v3/organizations/me",
		},
		{
			name:        "a base URL with a path is left alone",
			operationID: identityOp,
			baseURL:     "https://us.svc.konghq.tech/api",
			wantHost:    "us.svc.konghq.tech",
			wantPath:    "/api/v3/organizations/me",
		},
		{
			name:        "a trailing slash on the path is still a path",
			operationID: identityOp,
			baseURL:     "https://us.svc.konghq.tech/api/",
			wantHost:    "us.svc.konghq.tech",
			wantPath:    "/api/v3/organizations/me",
		},
		{
			name:        "an explicit port on a path-carrying base URL is left alone",
			operationID: identityOp,
			baseURL:     "https://us.svc.konghq.tech:8443/api",
			wantHost:    "us.svc.konghq.tech:8443",
			wantPath:    "/api/v3/organizations/me",
		},
		{
			name:        "a lone slash is not a path, so the rewrite still happens",
			operationID: identityOp,
			baseURL:     "/",
			wantHost:    "global.api.konghq.com",
			wantPath:    "/v3/organizations/me",
		},
		{
			name:        "another operation is not touched",
			operationID: "list-acl",
			baseURL:     "https://us.api.konghq.com/",
			wantHost:    "us.api.konghq.com",
			wantPath:    "/v3/organizations/me",
		},
		{
			name:         "the custom domain hook still rewrites the global host",
			operationID:  identityOp,
			baseURL:      "https://us.api.konghq.com/",
			customDomain: "konghq.tech",
			wantHost:     "global.api.konghq.tech",
			wantPath:     "/v3/organizations/me",
		},
		{
			name:         "the custom domain hook still applies to a private host",
			operationID:  identityOp,
			baseURL:      "https://us.svc.konghq.com/api",
			customDomain: "konghq.tech",
			wantHost:     "us.svc.konghq.tech",
			wantPath:     "/api/v3/organizations/me",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opURL, err := url.JoinPath(tt.baseURL, "/v3/organizations/me")
			require.NoError(t, err)

			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, opURL, nil)
			require.NoError(t, err)

			hookCtx := BeforeRequestContext{
				HookContext: HookContext{
					Context:     context.Background(),
					BaseURL:     tt.baseURL,
					OperationID: tt.operationID,
				},
			}

			req, err = (&GlobalAPIURLRequestHook{}).BeforeRequest(hookCtx, req)
			require.NoError(t, err)

			if tt.customDomain != "" {
				req, err = (&APIURLRequestHook{CustomDomain: tt.customDomain}).BeforeRequest(hookCtx, req)
				require.NoError(t, err)
			}

			require.Equal(t, tt.wantHost, req.URL.Host)
			require.Equal(t, tt.wantPath, req.URL.Path)
		})
	}
}
