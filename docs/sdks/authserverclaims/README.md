# AuthServerClaims

## Overview

Claims are statements about the Client, included in tokens issued by the Auth Server. The management API will give you the ability to create, configure and manage multiple Claims per Auth Server, and include them in tokens based on the requested Scopes.

### Available Operations

* [ListAuthServerClaims](#listauthserverclaims) - List all auth server claims
* [CreateAuthServerClaim](#createauthserverclaim) - Create a new auth server claim
* [GetAuthServerClaim](#getauthserverclaim) - Get an auth server claim
* [UpdateAuthServerClaim](#updateauthserverclaim) - Update an auth server claim
* [DeleteAuthServerClaim](#deleteauthserverclaim) - Delete an auth server claim

## ListAuthServerClaims

List all the claims for a given auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAuthServerClaims" method="get" path="/v1/auth-servers/{authServerId}/claims" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.AuthServerClaims.ListAuthServerClaims(ctx, operations.ListAuthServerClaimsRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClaimList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListAuthServerClaimsRequest](../../models/operations/listauthserverclaimsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListAuthServerClaimsResponse](../../models/operations/listauthserverclaimsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## CreateAuthServerClaim

Create a new claim for a given auth server. The `name` attribute will appear in the JWT access token as a claim with the given case-sensitive name. The `value` attribute will be used to populate the claim value in the JWT access token. The value can be templated, which allows you to use variables like `${ uuidv4 }` to generate dynamic values. Claim values are hydrated based on their types, which may include JSON.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAuthServerClaim" method="post" path="/v1/auth-servers/{authServerId}/claims" -->
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

    res, err := s.AuthServerClaims.CreateAuthServerClaim(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateClaim{
        Name: "<value>",
        Value: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Claim != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |                                                                  |
| `authServerID`                                                   | `string`                                                         | :heavy_check_mark:                                               | The auth server ID                                               | d32d905a-ed33-46a3-a093-d8f536af9a8a                             |
| `createClaim`                                                    | [components.CreateClaim](../../models/components/createclaim.md) | :heavy_check_mark:                                               | Claim to be created                                              |                                                                  |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |                                                                  |

### Response

**[*operations.CreateAuthServerClaimResponse](../../models/operations/createauthserverclaimresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetAuthServerClaim

Get a claim for given auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAuthServerClaim" method="get" path="/v1/auth-servers/{authServerId}/claims/{claimId}" -->
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

    res, err := s.AuthServerClaims.GetAuthServerClaim(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "07d05309-45cc-4b37-92fb-1524846deec3")
    if err != nil {
        log.Fatal(err)
    }
    if res.Claim != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `authServerID`                                           | `string`                                                 | :heavy_check_mark:                                       | The auth server ID                                       | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `claimID`                                                | `string`                                                 | :heavy_check_mark:                                       | The claim ID                                             | 07d05309-45cc-4b37-92fb-1524846deec3                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAuthServerClaimResponse](../../models/operations/getauthserverclaimresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## UpdateAuthServerClaim

Update a claim for a given auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAuthServerClaim" method="patch" path="/v1/auth-servers/{authServerId}/claims/{claimId}" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.AuthServerClaims.UpdateAuthServerClaim(ctx, operations.UpdateAuthServerClaimRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        ClaimID: "07d05309-45cc-4b37-92fb-1524846deec3",
        UpdateClaim: components.UpdateClaim{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Claim != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateAuthServerClaimRequest](../../models/operations/updateauthserverclaimrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateAuthServerClaimResponse](../../models/operations/updateauthserverclaimresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## DeleteAuthServerClaim

Delete a claim for a given auth server. If the claim is included in any scopes, it will be removed from those scopes.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAuthServerClaim" method="delete" path="/v1/auth-servers/{authServerId}/claims/{claimId}" -->
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

    res, err := s.AuthServerClaims.DeleteAuthServerClaim(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "07d05309-45cc-4b37-92fb-1524846deec3")
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
| `authServerID`                                           | `string`                                                 | :heavy_check_mark:                                       | The auth server ID                                       | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `claimID`                                                | `string`                                                 | :heavy_check_mark:                                       | The claim ID                                             | 07d05309-45cc-4b37-92fb-1524846deec3                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAuthServerClaimResponse](../../models/operations/deleteauthserverclaimresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |