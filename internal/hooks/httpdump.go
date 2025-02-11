package hooks

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

// HTTPDumpRequestHook is a hook that dumps the request to stdout.
type HTTPDumpRequestHook struct{}

var _ beforeRequestHook = (*HTTPDumpRequestHook)(nil)

// BeforeRequest dumps the request to stdout if enabled.
func (i *HTTPDumpRequestHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if req == nil {
		return nil, nil
	}

	// Dumping the body consumes it, so we need to clone the request and copy
	// the body as Clone() only performs a shallow copy.

	cpy := req.Clone(hookCtx.Context)
	var (
		b   []byte
		err error
	)
	if req.Body != nil {
		b, err = io.ReadAll(req.Body)
		if err != nil {
			return req, fmt.Errorf("error reading request body: %w", err)
		}
		cpy.Body = io.NopCloser(bytes.NewReader(b))
	}

	dump, err := httputil.DumpRequestOut(cpy, true)
	if err != nil {
		fmt.Printf("Error dumping request: %v\n", err)
	} else {
		fmt.Printf("request:\n%s\n\n", dump)
	}
	req.Body = io.NopCloser(bytes.NewReader(b))

	return req, nil
}

// HTTPDumpResponseHook is a hook that dumps the response to stdout.
type HTTPDumpResponseHook struct{}

var _ afterSuccessHook = (*HTTPDumpResponseHook)(nil)

// AfterSuccess dumps the response to stdout if enabled.
func (i *HTTPDumpResponseHook) AfterSuccess(hookCtx AfterSuccessContext, res *http.Response) (*http.Response, error) {
	return dumpResponse(res), nil
}

// HTTPDumpResponseErrorHook is a hook that dumps the error response to stdout.
type HTTPDumpResponseErrorHook struct{}

var _ afterErrorHook = (*HTTPDumpResponseErrorHook)(nil)

// AfterError dumps the error response to stdout if enabled.
func (i *HTTPDumpResponseErrorHook) AfterError(hookCtx AfterErrorContext, res *http.Response, err error) (*http.Response, error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	return dumpResponse(res), err
}

func dumpResponse(res *http.Response) *http.Response {
	if res == nil {
		return nil
	}

	b, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Printf("Error dumping respone: %v\n", err)
	} else {
		fmt.Printf("response:\n%s\n\n", b)
	}

	return res
}
