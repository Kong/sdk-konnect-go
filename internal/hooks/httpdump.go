package hooks

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// HTTPDumpRequestHook is a hook that dumps the request to stdout.
type HTTPDumpRequestHook struct {
	Enabled bool
}

var _ beforeRequestHook = (*HTTPDumpRequestHook)(nil)

// BeforeRequest dumps the request to stdout if enabled.
func (i *HTTPDumpRequestHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if !i.Enabled {
		return req, nil
	}

	b, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		fmt.Printf("Error dumping request: %v\n", err)
	} else {
		fmt.Printf("request:\n%s\n\n", b)
	}

	return req, nil
}

// HTTPDumpResponseHook is a hook that dumps the response to stdout.
type HTTPDumpResponseHook struct {
	Enabled bool
}

var _ afterSuccessHook = (*HTTPDumpResponseHook)(nil)

// AfterSuccess dumps the response to stdout if enabled.
func (i *HTTPDumpResponseHook) AfterSuccess(hookCtx AfterSuccessContext, res *http.Response) (*http.Response, error) {
	if !i.Enabled {
		return res, nil
	}

	b, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Printf("Error dumping respone: %v\n", err)
	} else {
		fmt.Printf("response:\n%s\n\n", b)
	}

	return res, nil
}
