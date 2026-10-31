# ConnectedApps

## Overview

Connected Apps enable third-party applications to access Konnect APIs on behalf of an organization using OAuth 2.0. When enabled, a reserved auth server is created with predefined scopes.

### Available Operations

* [ListPortalConnectedAppSessions](#listportalconnectedappsessions) - List caller's Portal Connected App sessions
* [RevokePortalConnectedAppSession](#revokeportalconnectedappsession) - Revoke one of caller's Portal App sessions
* [PatchPortalConnectedApps](#patchportalconnectedapps) - Patch Portal Connected Apps configuration
* [ListPortalConnectedAppSessionsAdmin](#listportalconnectedappsessionsadmin) - List Portal Connected App sessions
* [RevokePortalConnectedAppSessionAdmin](#revokeportalconnectedappsessionadmin) - Revoke a Portal Connected App session

## ListPortalConnectedAppSessions

Returns the authenticated caller's active Portal Connected App sessions (active refresh tokens). Restricted to developer principals carrying a portal_id claim. Authenticated but not authorized.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPortalConnectedAppSessions" method="get" path="/v1/connected-apps/portal/me/sessions" -->
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

    res, err := s.ConnectedApps.ListPortalConnectedAppSessions(ctx, operations.ListPortalConnectedAppSessionsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListConnectedAppSessionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ListPortalConnectedAppSessionsRequest](../../models/operations/listportalconnectedappsessionsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.ListPortalConnectedAppSessionsResponse](../../models/operations/listportalconnectedappsessionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## RevokePortalConnectedAppSession

Revokes a single active Portal Connected App session (refresh token) owned by the authenticated caller. Restricted to developer principals carrying a portal_id claim. Authenticated but not authorized.

### Example Usage

<!-- UsageSnippet language="go" operationID="revokePortalConnectedAppSession" method="delete" path="/v1/connected-apps/portal/me/sessions/{sessionId}" -->
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

    res, err := s.ConnectedApps.RevokePortalConnectedAppSession(ctx, "3f1e5c2a-7b8d-4e6f-9a0b-1c2d3e4f5a6b")
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
| `sessionID`                                              | `string`                                                 | :heavy_check_mark:                                       | The session ID                                           | 3f1e5c2a-7b8d-4e6f-9a0b-1c2d3e4f5a6b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RevokePortalConnectedAppSessionResponse](../../models/operations/revokeportalconnectedappsessionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchPortalConnectedApps

Partially updates the given portal's Connected Apps configuration. Only fields present in the request body are changed. Setting `enabled` enables or disables Connected Apps for this portal (same effect as POST /connected-apps; enabling for the first time requires default_url to also be set); omitting it patches dcr and/or default_url on the already-enabled auth server, which returns 404 if Connected Apps has not been enabled yet for this portal. The portal id is immutable and always taken from the path.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPortalConnectedApps" method="patch" path="/v1/connected-apps/portals/{portalId}" -->
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

    res, err := s.ConnectedApps.PatchPortalConnectedApps(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.PatchPortalConnectedApps{
        DefaultURL: sdkkonnectgo.Pointer("https://portal.example.com"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectedAppsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `portalID`                                                                                 | `string`                                                                                   | :heavy_check_mark:                                                                         | The portal ID                                                                              | d32d905a-ed33-46a3-a093-d8f536af9a8a                                                       |
| `patchPortalConnectedApps`                                                                 | [components.PatchPortalConnectedApps](../../models/components/patchportalconnectedapps.md) | :heavy_check_mark:                                                                         | Portal Connected Apps configuration fields to patch                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.PatchPortalConnectedAppsResponse](../../models/operations/patchportalconnectedappsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListPortalConnectedAppSessionsAdmin

Returns all active Portal Connected App sessions (active refresh tokens) for the given portal, across all developers. Optionally filter to a single developer with filter[developer_id].

### Example Usage

<!-- UsageSnippet language="go" operationID="listPortalConnectedAppSessionsAdmin" method="get" path="/v1/connected-apps/portals/{portalId}/sessions" -->
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

    res, err := s.ConnectedApps.ListPortalConnectedAppSessionsAdmin(ctx, operations.ListPortalConnectedAppSessionsAdminRequest{
        PortalID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPortalConnectedAppSessionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.ListPortalConnectedAppSessionsAdminRequest](../../models/operations/listportalconnectedappsessionsadminrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.ListPortalConnectedAppSessionsAdminResponse](../../models/operations/listportalconnectedappsessionsadminresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## RevokePortalConnectedAppSessionAdmin

Revokes a single active Portal Connected App session (refresh token) for the given portal, regardless of which developer owns it.

### Example Usage

<!-- UsageSnippet language="go" operationID="revokePortalConnectedAppSessionAdmin" method="delete" path="/v1/connected-apps/portals/{portalId}/sessions/{sessionId}" -->
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

    res, err := s.ConnectedApps.RevokePortalConnectedAppSessionAdmin(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "3f1e5c2a-7b8d-4e6f-9a0b-1c2d3e4f5a6b")
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
| `portalID`                                               | `string`                                                 | :heavy_check_mark:                                       | The portal ID                                            | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `sessionID`                                              | `string`                                                 | :heavy_check_mark:                                       | The session ID                                           | 3f1e5c2a-7b8d-4e6f-9a0b-1c2d3e4f5a6b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RevokePortalConnectedAppSessionAdminResponse](../../models/operations/revokeportalconnectedappsessionadminresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |