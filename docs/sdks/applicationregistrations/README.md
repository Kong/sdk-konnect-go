# ApplicationRegistrations

## Overview

APIs related to Konnect Developer Portal Application Registrations.

### Available Operations

* [ListRegistrations](#listregistrations) - List Registrations by Portal
* [ListRegistrationsByApplication](#listregistrationsbyapplication) - List Registrations by Application
* [GetApplicationRegistration](#getapplicationregistration) - Get a Registration
* [UpdateApplicationRegistration](#updateapplicationregistration) - Update Registration
* [DeleteApplicationRegistration](#deleteapplicationregistration) - Delete Registration

## ListRegistrations

Lists all of the application registrations and their current status (e.g., approved or pending) for this portal. Each registration is associated with a single API. Access is provided through the credentials issued to the application that contains each registration.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-registrations" method="get" path="/v3/portals/{portalId}/application-registrations" example="Example 1" -->
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

    res, err := s.ApplicationRegistrations.ListRegistrations(ctx, operations.ListRegistrationsRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListApplicationRegistrationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListRegistrationsRequest](../../models/operations/listregistrationsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListRegistrationsResponse](../../models/operations/listregistrationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListRegistrationsByApplication

Lists each API that this application is registered for and their current status (e.g., pending, approved, rejected, revoked).

### Example Usage

<!-- UsageSnippet language="go" operationID="list-registrations-by-application" method="get" path="/v3/portals/{portalId}/applications/{applicationId}/registrations" example="Example 1" -->
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

    res, err := s.ApplicationRegistrations.ListRegistrationsByApplication(ctx, operations.ListRegistrationsByApplicationRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ApplicationID: "5e6d301a-4935-48f9-9525-817ba1a91ce3",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListApplicationRegistrationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ListRegistrationsByApplicationRequest](../../models/operations/listregistrationsbyapplicationrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.ListRegistrationsByApplicationResponse](../../models/operations/listregistrationsbyapplicationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetApplicationRegistration

Returns information about an application's registration status for a particular API.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-application-registration" method="get" path="/v3/portals/{portalId}/applications/{applicationId}/registrations/{registrationId}" example="Example 1" -->
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

    res, err := s.ApplicationRegistrations.GetApplicationRegistration(ctx, operations.GetApplicationRegistrationRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ApplicationID: "09efcbdf-b833-49ac-b5c9-a0dc1ec0ef6a",
        RegistrationID: "4f131ab6-5c14-41b9-b22c-347366ac1ac0",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicationRegistrationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetApplicationRegistrationRequest](../../models/operations/getapplicationregistrationrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetApplicationRegistrationResponse](../../models/operations/getapplicationregistrationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateApplicationRegistration

Updates the status of a particular application registration to an API. Approved application registrations will allow API traffic to the corresponding API. Revoked, rejected, or pending will not allow API traffic.

### Example Usage: Approved

<!-- UsageSnippet language="go" operationID="update-application-registration" method="patch" path="/v3/portals/{portalId}/applications/{applicationId}/registrations/{registrationId}" example="Approved" -->
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

    res, err := s.ApplicationRegistrations.UpdateApplicationRegistration(ctx, operations.UpdateApplicationRegistrationRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ApplicationID: "8531a3ec-718e-40e8-b5ec-e12e6e00b6e1",
        RegistrationID: "f7e27313-022a-404f-b641-ad7fca249f28",
        UpdateApplicationRegistrationRequest: components.UpdateApplicationRegistrationRequest{
            Status: components.ApplicationRegistrationStatusApproved.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicationRegistrationResponse != nil {
        // handle response
    }
}
```
### Example Usage: Example 1

<!-- UsageSnippet language="go" operationID="update-application-registration" method="patch" path="/v3/portals/{portalId}/applications/{applicationId}/registrations/{registrationId}" example="Example 1" -->
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

    res, err := s.ApplicationRegistrations.UpdateApplicationRegistration(ctx, operations.UpdateApplicationRegistrationRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ApplicationID: "ee68e99c-bd3a-4396-a9ee-a17dce84d550",
        RegistrationID: "b937480e-01a8-4b5c-9e26-76d376e62f55",
        UpdateApplicationRegistrationRequest: components.UpdateApplicationRegistrationRequest{
            Status: components.ApplicationRegistrationStatusRejected.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicationRegistrationResponse != nil {
        // handle response
    }
}
```
### Example Usage: NotFoundExample

<!-- UsageSnippet language="go" operationID="update-application-registration" method="patch" path="/v3/portals/{portalId}/applications/{applicationId}/registrations/{registrationId}" example="NotFoundExample" -->
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

    res, err := s.ApplicationRegistrations.UpdateApplicationRegistration(ctx, operations.UpdateApplicationRegistrationRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ApplicationID: "24acd281-7019-4d8f-9e94-e900d83c9e71",
        RegistrationID: "0475f524-437f-4af4-b7a7-6e2c3f909ab0",
        UpdateApplicationRegistrationRequest: components.UpdateApplicationRegistrationRequest{
            Status: components.ApplicationRegistrationStatusRejected.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicationRegistrationResponse != nil {
        // handle response
    }
}
```
### Example Usage: Pending

<!-- UsageSnippet language="go" operationID="update-application-registration" method="patch" path="/v3/portals/{portalId}/applications/{applicationId}/registrations/{registrationId}" example="Pending" -->
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

    res, err := s.ApplicationRegistrations.UpdateApplicationRegistration(ctx, operations.UpdateApplicationRegistrationRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ApplicationID: "1a7a1699-6816-470f-b652-dc045af229de",
        RegistrationID: "619b15a7-357b-481e-baa1-b06410b42196",
        UpdateApplicationRegistrationRequest: components.UpdateApplicationRegistrationRequest{
            Status: components.ApplicationRegistrationStatusPending.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicationRegistrationResponse != nil {
        // handle response
    }
}
```
### Example Usage: Rejected

<!-- UsageSnippet language="go" operationID="update-application-registration" method="patch" path="/v3/portals/{portalId}/applications/{applicationId}/registrations/{registrationId}" example="Rejected" -->
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

    res, err := s.ApplicationRegistrations.UpdateApplicationRegistration(ctx, operations.UpdateApplicationRegistrationRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ApplicationID: "cd96be86-5548-4f11-9752-605e4daafde8",
        RegistrationID: "1d7446b9-6174-472a-8d14-2750bc372841",
        UpdateApplicationRegistrationRequest: components.UpdateApplicationRegistrationRequest{
            Status: components.ApplicationRegistrationStatusRejected.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicationRegistrationResponse != nil {
        // handle response
    }
}
```
### Example Usage: Revoked

<!-- UsageSnippet language="go" operationID="update-application-registration" method="patch" path="/v3/portals/{portalId}/applications/{applicationId}/registrations/{registrationId}" example="Revoked" -->
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

    res, err := s.ApplicationRegistrations.UpdateApplicationRegistration(ctx, operations.UpdateApplicationRegistrationRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ApplicationID: "72904692-d677-4cc3-a196-3ebb309d356c",
        RegistrationID: "741d9097-1faf-4536-a52d-8fb48d72bf9c",
        UpdateApplicationRegistrationRequest: components.UpdateApplicationRegistrationRequest{
            Status: components.ApplicationRegistrationStatusRevoked.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicationRegistrationResponse != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="update-application-registration" method="patch" path="/v3/portals/{portalId}/applications/{applicationId}/registrations/{registrationId}" example="Unauthorized" -->
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

    res, err := s.ApplicationRegistrations.UpdateApplicationRegistration(ctx, operations.UpdateApplicationRegistrationRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ApplicationID: "88eadf83-49d4-44a2-86ec-1376a185a864",
        RegistrationID: "aba7a860-6fae-45bc-9e8e-81f1684a88d3",
        UpdateApplicationRegistrationRequest: components.UpdateApplicationRegistrationRequest{
            Status: components.ApplicationRegistrationStatusRejected.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicationRegistrationResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-application-registration" method="patch" path="/v3/portals/{portalId}/applications/{applicationId}/registrations/{registrationId}" example="UnauthorizedExample" -->
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

    res, err := s.ApplicationRegistrations.UpdateApplicationRegistration(ctx, operations.UpdateApplicationRegistrationRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ApplicationID: "fe006d44-2485-4c82-ab54-ddc1b7427bac",
        RegistrationID: "72b2aca8-229a-4a12-aab3-a490c9f41cc4",
        UpdateApplicationRegistrationRequest: components.UpdateApplicationRegistrationRequest{
            Status: components.ApplicationRegistrationStatusRejected.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicationRegistrationResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateApplicationRegistrationBadRequestExample1

<!-- UsageSnippet language="go" operationID="update-application-registration" method="patch" path="/v3/portals/{portalId}/applications/{applicationId}/registrations/{registrationId}" example="UpdateApplicationRegistrationBadRequestExample1" -->
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

    res, err := s.ApplicationRegistrations.UpdateApplicationRegistration(ctx, operations.UpdateApplicationRegistrationRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ApplicationID: "5d8b3e89-b759-4ce5-b708-322b575f974b",
        RegistrationID: "c7ff2e33-99e7-4e9d-9861-7686dea6686b",
        UpdateApplicationRegistrationRequest: components.UpdateApplicationRegistrationRequest{
            Status: components.ApplicationRegistrationStatusRejected.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicationRegistrationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.UpdateApplicationRegistrationRequest](../../models/operations/updateapplicationregistrationrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.UpdateApplicationRegistrationResponse](../../models/operations/updateapplicationregistrationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteApplicationRegistration

Deletes an application registration, which if currently approved will immediately block API traffic to the API.
Note: Developers can request a new application registration for the given API as long as they have RBAC access to consume.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-application-registration" method="delete" path="/v3/portals/{portalId}/applications/{applicationId}/registrations/{registrationId}" -->
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

    res, err := s.ApplicationRegistrations.DeleteApplicationRegistration(ctx, operations.DeleteApplicationRegistrationRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ApplicationID: "eb118d8f-e927-40ac-8a2b-a918daeed1dc",
        RegistrationID: "2d31c856-eb91-4ec9-a008-ce29dff66288",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.DeleteApplicationRegistrationRequest](../../models/operations/deleteapplicationregistrationrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.DeleteApplicationRegistrationResponse](../../models/operations/deleteapplicationregistrationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |