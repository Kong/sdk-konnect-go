# IntegrationInstanceAuthCredentials

## Overview

Represents the credentials use to authorize an integration instance.
You will want to configure the integration instance settings and authorization configuration before authorizing the instance.
This will inform the authorization process on how to reach and authorize the account.
Once the integration instance is authorized, the system will automatically discover all the relevant resources from the account.
The integration instance's auth credentials can be removed or updated while retaining all resources which have already been discovered.


### Available Operations

* [CreateIntegrationInstanceAuthCredential](#createintegrationinstanceauthcredential) - Create Integration Instance Auth Credential
* [GetIntegrationInstanceAuthCredential](#getintegrationinstanceauthcredential) - Get Integration Instance Auth Credential
* [DeleteIntegrationInstanceAuthCredential](#deleteintegrationinstanceauthcredential) - Delete Integration Instance Auth Credential

## CreateIntegrationInstanceAuthCredential

Creates an auth credential scoped to the given integration instance.
Auth credentials are singleton resources that have a 1-to-1 relationship with
an integration instance.

An attempt to create subsequent auth credentials for an instance will result in a 409 response.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-integration-instance-auth-credential" method="post" path="/v1/integration-instances/{integrationInstanceId}/auth-credential" -->
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

    res, err := s.IntegrationInstanceAuthCredentials.CreateIntegrationInstanceAuthCredential(ctx, "3f51fa25-310a-421d-bd1a-007f859021a3", components.CreateCreateIntegrationInstanceAuthCredentialMultiKeyAuth(
        components.MultiKeyAuth{
            Config: components.CreateMultiKeyAuthCredentialConfig{
                Headers: []components.Headers{
                    components.Headers{
                        Name: "x-api-key",
                        Key: "9f2a3b4c8d6e7f00112233445566778899aabbccddeeff001122334455667788",
                    },
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationInstanceAuthCredential != nil {
        switch res.IntegrationInstanceAuthCredential.Type {
            case components.IntegrationInstanceAuthCredentialTypeOauth:
                // res.IntegrationInstanceAuthCredential.Oauth is populated
            case components.IntegrationInstanceAuthCredentialTypeGithubAppInstallation:
                // res.IntegrationInstanceAuthCredential.GithubAppInstallation is populated
            case components.IntegrationInstanceAuthCredentialTypeMultiKeyAuthCredential:
                // res.IntegrationInstanceAuthCredential.MultiKeyAuthCredential is populated
            case components.IntegrationInstanceAuthCredentialTypeAWSRoleDelegationAuthCredential:
                // res.IntegrationInstanceAuthCredential.AWSRoleDelegationAuthCredential is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              | Example                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |                                                                                                                          |
| `integrationInstanceID`                                                                                                  | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | The `id` of the integration instance.                                                                                    | 3f51fa25-310a-421d-bd1a-007f859021a3                                                                                     |
| `createIntegrationInstanceAuthCredential`                                                                                | [components.CreateIntegrationInstanceAuthCredential](../../models/components/createintegrationinstanceauthcredential.md) | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |                                                                                                                          |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |                                                                                                                          |

### Response

**[*operations.CreateIntegrationInstanceAuthCredentialResponse](../../models/operations/createintegrationinstanceauthcredentialresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetIntegrationInstanceAuthCredential

Fetches the auth credential scoped to the given integration instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-integration-instance-auth-credential" method="get" path="/v1/integration-instances/{integrationInstanceId}/auth-credential" -->
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

    res, err := s.IntegrationInstanceAuthCredentials.GetIntegrationInstanceAuthCredential(ctx, "3f51fa25-310a-421d-bd1a-007f859021a3")
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationInstanceAuthCredential != nil {
        switch res.IntegrationInstanceAuthCredential.Type {
            case components.IntegrationInstanceAuthCredentialTypeOauth:
                // res.IntegrationInstanceAuthCredential.Oauth is populated
            case components.IntegrationInstanceAuthCredentialTypeGithubAppInstallation:
                // res.IntegrationInstanceAuthCredential.GithubAppInstallation is populated
            case components.IntegrationInstanceAuthCredentialTypeMultiKeyAuthCredential:
                // res.IntegrationInstanceAuthCredential.MultiKeyAuthCredential is populated
            case components.IntegrationInstanceAuthCredentialTypeAWSRoleDelegationAuthCredential:
                // res.IntegrationInstanceAuthCredential.AWSRoleDelegationAuthCredential is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `integrationInstanceID`                                  | *string*                                                 | :heavy_check_mark:                                       | The `id` of the integration instance.                    | 3f51fa25-310a-421d-bd1a-007f859021a3                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetIntegrationInstanceAuthCredentialResponse](../../models/operations/getintegrationinstanceauthcredentialresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteIntegrationInstanceAuthCredential

Deletes the auth credential scoped to the given integration instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-integration-instance-auth-credential" method="delete" path="/v1/integration-instances/{integrationInstanceId}/auth-credential" -->
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

    res, err := s.IntegrationInstanceAuthCredentials.DeleteIntegrationInstanceAuthCredential(ctx, "3f51fa25-310a-421d-bd1a-007f859021a3")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `integrationInstanceID`                                  | *string*                                                 | :heavy_check_mark:                                       | The `id` of the integration instance.                    | 3f51fa25-310a-421d-bd1a-007f859021a3                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteIntegrationInstanceAuthCredentialResponse](../../models/operations/deleteintegrationinstanceauthcredentialresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |