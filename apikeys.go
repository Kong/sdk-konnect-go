// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkkonnectgo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/hooks"
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"github.com/Kong/sdk-konnect-go/models/sdkerrors"
	"github.com/cenkalti/backoff/v4"
	"io"
	"net/http"
)

type APIKeys struct {
	sdkConfiguration sdkConfiguration
}

func newAPIKeys(sdkConfig sdkConfiguration) *APIKeys {
	return &APIKeys{
		sdkConfiguration: sdkConfig,
	}
}

// ListKeyAuthWithConsumer - List all API-keys associated with a Consumer
// List all API-keys associated with a Consumer
func (s *APIKeys) ListKeyAuthWithConsumer(ctx context.Context, request operations.ListKeyAuthWithConsumerRequest, opts ...operations.Option) (*operations.ListKeyAuthWithConsumerResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "list-key-auth-with-consumer",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/key-auth", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				return nil, backoff.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.ListKeyAuthWithConsumerResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	getRawBody := func() ([]byte, error) {
		rawBody, err := io.ReadAll(httpRes.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}
		httpRes.Body.Close()
		httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
		return rawBody, nil
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := getRawBody()
			if err != nil {
				return nil, err
			}

			var out operations.ListKeyAuthWithConsumerResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			rawBody, err := getRawBody()
			if err != nil {
				return nil, err
			}

			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := getRawBody()
		if err != nil {
			return nil, err
		}

		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := getRawBody()
		if err != nil {
			return nil, err
		}

		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// CreateKeyAuthWithConsumer - Create a new API-key associated with a Consumer
// Create a new API-key associated with a Consumer
func (s *APIKeys) CreateKeyAuthWithConsumer(ctx context.Context, request operations.CreateKeyAuthWithConsumerRequest, opts ...operations.Option) (*operations.CreateKeyAuthWithConsumerResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "create-key-auth-with-consumer",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/key-auth", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "KeyAuthWithoutParents", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				return nil, backoff.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.CreateKeyAuthWithConsumerResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	getRawBody := func() ([]byte, error) {
		rawBody, err := io.ReadAll(httpRes.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}
		httpRes.Body.Close()
		httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
		return rawBody, nil
	}

	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := getRawBody()
			if err != nil {
				return nil, err
			}

			var out components.KeyAuth
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.KeyAuth = &out
		default:
			rawBody, err := getRawBody()
			if err != nil {
				return nil, err
			}

			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := getRawBody()
		if err != nil {
			return nil, err
		}

		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := getRawBody()
		if err != nil {
			return nil, err
		}

		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// DeleteKeyAuthWithConsumer - Delete a an API-key associated with a Consumer
// Delete a an API-key associated with a Consumer using ID.
func (s *APIKeys) DeleteKeyAuthWithConsumer(ctx context.Context, request operations.DeleteKeyAuthWithConsumerRequest, opts ...operations.Option) (*operations.DeleteKeyAuthWithConsumerResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "delete-key-auth-with-consumer",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/key-auth/{KeyAuthId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				return nil, backoff.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.DeleteKeyAuthWithConsumerResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	getRawBody := func() ([]byte, error) {
		rawBody, err := io.ReadAll(httpRes.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}
		httpRes.Body.Close()
		httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
		return rawBody, nil
	}

	switch {
	case httpRes.StatusCode == 204:
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := getRawBody()
		if err != nil {
			return nil, err
		}

		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := getRawBody()
		if err != nil {
			return nil, err
		}

		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// GetKeyAuthWithConsumer - Fetch an API-key associated with a Consumer
// Get an API-key associated with a Consumer using ID.
func (s *APIKeys) GetKeyAuthWithConsumer(ctx context.Context, request operations.GetKeyAuthWithConsumerRequest, opts ...operations.Option) (*operations.GetKeyAuthWithConsumerResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "get-key-auth-with-consumer",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/key-auth/{KeyAuthId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				return nil, backoff.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"404", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.GetKeyAuthWithConsumerResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	getRawBody := func() ([]byte, error) {
		rawBody, err := io.ReadAll(httpRes.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}
		httpRes.Body.Close()
		httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
		return rawBody, nil
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := getRawBody()
			if err != nil {
				return nil, err
			}

			var out components.KeyAuth
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.KeyAuth = &out
		default:
			rawBody, err := getRawBody()
			if err != nil {
				return nil, err
			}

			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := getRawBody()
		if err != nil {
			return nil, err
		}

		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := getRawBody()
		if err != nil {
			return nil, err
		}

		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// UpsertKeyAuthWithConsumer - Upsert an API-key associated with a Consumer
// Create or Update an API-key associated with a Consumer using ID.
func (s *APIKeys) UpsertKeyAuthWithConsumer(ctx context.Context, request operations.UpsertKeyAuthWithConsumerRequest, opts ...operations.Option) (*operations.UpsertKeyAuthWithConsumerResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "upsert-key-auth-with-consumer",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/key-auth/{KeyAuthId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "KeyAuthWithoutParents", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				return nil, backoff.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.UpsertKeyAuthWithConsumerResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	getRawBody := func() ([]byte, error) {
		rawBody, err := io.ReadAll(httpRes.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}
		httpRes.Body.Close()
		httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
		return rawBody, nil
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := getRawBody()
			if err != nil {
				return nil, err
			}

			var out components.KeyAuth
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.KeyAuth = &out
		default:
			rawBody, err := getRawBody()
			if err != nil {
				return nil, err
			}

			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := getRawBody()
		if err != nil {
			return nil, err
		}

		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := getRawBody()
		if err != nil {
			return nil, err
		}

		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// ListKeyAuth - List all API-keys
// List all API-keys
func (s *APIKeys) ListKeyAuth(ctx context.Context, request operations.ListKeyAuthRequest, opts ...operations.Option) (*operations.ListKeyAuthResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "list-key-auth",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/v2/control-planes/{controlPlaneId}/core-entities/key-auths", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				return nil, backoff.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"401", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.ListKeyAuthResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	getRawBody := func() ([]byte, error) {
		rawBody, err := io.ReadAll(httpRes.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}
		httpRes.Body.Close()
		httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
		return rawBody, nil
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := getRawBody()
			if err != nil {
				return nil, err
			}

			var out operations.ListKeyAuthResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			rawBody, err := getRawBody()
			if err != nil {
				return nil, err
			}

			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := getRawBody()
			if err != nil {
				return nil, err
			}

			var out sdkerrors.GatewayUnauthorizedError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			return nil, &out
		default:
			rawBody, err := getRawBody()
			if err != nil {
				return nil, err
			}

			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := getRawBody()
		if err != nil {
			return nil, err
		}

		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := getRawBody()
		if err != nil {
			return nil, err
		}

		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}
