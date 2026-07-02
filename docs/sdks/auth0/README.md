# Auth0

## Overview

### Available Operations

* [PostAuth0UserMfaSettingsInternal](#postauth0usermfasettingsinternal) - Get Auth0 User MFA Settings (Internal)

## PostAuth0UserMfaSettingsInternal

Returns whether Multi-Factor Authentication (MFA) is enabled for an Auth0 user. A user's MFA status is determined from their Konnect organization's settings. If the user is associated with multiple organizations, MFA is considered enabled if any of their organizations have MFA enabled. Although a user is generally uniquely identified by their email address, this endpoint requires the user's Auth0 ID to manage cases where the user logs in with multiple authentication providers but chooses not to link their accounts (in which case a single email address may correspond to multiple Auth0 user IDs).
Note: The email address is not a part of the URL path to avoid PII data in request logs/traces.

### Example Usage

<!-- UsageSnippet language="go" operationID="post-auth0-user-mfa-settings-internal" method="post" path="/v3/internal/auth0/users/mfa-settings" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Auth0.PostAuth0UserMfaSettingsInternal(ctx, components.GetMFASettingsRequest{
        Email: "user@example.com",
        Auth0id: "auth0|1234567890",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetMFASettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.GetMFASettingsRequest](../../models/components/getmfasettingsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PostAuth0UserMfaSettingsInternalResponse](../../models/operations/postauth0usermfasettingsinternalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |