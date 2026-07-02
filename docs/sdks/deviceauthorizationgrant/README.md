# DeviceAuthorizationGrant

## Overview

### Available Operations

* [PostOauthDeviceAuthorize](#postoauthdeviceauthorize) - Device authorization request
* [PostOauthDeviceToken](#postoauthdevicetoken) - Device access token request
* [PostOauthDeviceAuthorizeUser](#postoauthdeviceauthorizeuser) - User device authorization request
* [PatchOauthDeviceConfirm](#patchoauthdeviceconfirm) - Device confirmation request

## PostOauthDeviceAuthorize

Initiates a device authorization workflow, generating and returning a unique device verification code.
See https://www.rfc-editor.org/rfc/rfc8628#section-3.1 for details.

### Example Usage

<!-- UsageSnippet language="go" operationID="post-oauth-device-authorize" method="post" path="/v3/internal/oauth/device/authorize" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New()

    res, err := s.DeviceAuthorizationGrant.PostOauthDeviceAuthorize(ctx, operations.PostOauthDeviceAuthorizeRequestBody{
        ClientID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeviceAuthorizationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PostOauthDeviceAuthorizeRequestBody](../../models/operations/postoauthdeviceauthorizerequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.PostOauthDeviceAuthorizeResponse](../../models/operations/postoauthdeviceauthorizeresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| sdkerrors.DeviceAuthorizationGrantAuthorizeBadRequest | 400                                                   | application/json                                      |
| sdkerrors.SDKError                                    | 4XX, 5XX                                              | \*/\*                                                 |

## PostOauthDeviceToken

This endpoint provides the machine client a means of being notified when a request for authorization is granted or rejected.
It is expected for the client to try the access token request repeatedly in a polling fashion based on the error code in the response.
See https://www.rfc-editor.org/rfc/rfc8628#section-3.4 for details.

### Example Usage

<!-- UsageSnippet language="go" operationID="post-oauth-device-token" method="post" path="/v3/internal/oauth/device/token" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New()

    res, err := s.DeviceAuthorizationGrant.PostOauthDeviceToken(ctx, operations.PostOauthDeviceTokenRequestBody{
        GrantType: "<value>",
        DeviceCode: "<value>",
        ClientID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeviceAccessTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PostOauthDeviceTokenRequestBody](../../models/operations/postoauthdevicetokenrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PostOauthDeviceTokenResponse](../../models/operations/postoauthdevicetokenresponse.md), error**

### Errors

| Error Type                                        | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| sdkerrors.DeviceAuthorizationGrantTokenBadRequest | 400                                               | application/json                                  |
| sdkerrors.SDKError                                | 4XX, 5XX                                          | \*/\*                                             |

## PostOauthDeviceAuthorizeUser

Marks the device code as authorized and is a means to provide the interactive UI flow with the necessary request metadata for the user to confirm the request.

### Example Usage

<!-- UsageSnippet language="go" operationID="post-oauth-device-authorize-user" method="post" path="/v3/internal/oauth/device/authorize/user" -->
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
            KonnectAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.DeviceAuthorizationGrant.PostOauthDeviceAuthorizeUser(ctx, components.DeviceAuthorizationUserRequest{
        UserCode: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeviceAuthorizationUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [components.DeviceAuthorizationUserRequest](../../models/components/deviceauthorizationuserrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PostOauthDeviceAuthorizeUserResponse](../../models/operations/postoauthdeviceauthorizeuserresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchOauthDeviceConfirm

Confirms the authorization request by marking the device code as confirmed.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-oauth-device-confirm" method="patch" path="/v3/internal/oauth/device/confirm" -->
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
            KonnectAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.DeviceAuthorizationGrant.PatchOauthDeviceConfirm(ctx, components.DeviceConfirmationRequest{
        UserCode: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.DeviceConfirmationRequest](../../models/components/deviceconfirmationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchOauthDeviceConfirmResponse](../../models/operations/patchoauthdeviceconfirmresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |