package hooks

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPDumpRequestHook_BeforeRequest(t *testing.T) {
	const expectedBody = `{"test": "body"}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.Equal(t, expectedBody, string(body))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL, bytes.NewBufferString(expectedBody))
	require.NoError(t, err)

	var (
		hook    = &HTTPDumpRequestHook{}
		hookCtx = BeforeRequestContext{
			HookContext: HookContext{
				Context:     context.Background(),
				OperationID: "test-123",
			},
		}
	)
	req, err = hook.BeforeRequest(hookCtx, req)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
}
