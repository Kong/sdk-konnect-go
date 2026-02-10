# Portals

## Overview

APIs related to configuration of Konnect Developer Portals.

### Available Operations

* [ListPortals](#listportals) - List Portals
* [CreatePortal](#createportal) - Create Portal
* [GetPortal](#getportal) - Get a Portal
* [UpdatePortal](#updateportal) - Update Portal
* [DeletePortal](#deleteportal) - Delete Portal

## ListPortals

Lists developer portals defined in this region for this organization. Each developer portal is available at a unique address and has isolated configuration, customization, developers, and applications.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-portals" method="get" path="/v3/portals" example="Example 1" -->
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

    res, err := s.Portals.ListPortals(ctx, operations.ListPortalsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.PortalFilterParameters{
            AuthenticationEnabled: sdkkonnectgo.Pointer(true),
            RbacEnabled: sdkkonnectgo.Pointer(true),
            AutoApproveDevelopers: sdkkonnectgo.Pointer(true),
            AutoApproveApplications: sdkkonnectgo.Pointer(true),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPortalsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListPortalsRequest](../../models/operations/listportalsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListPortalsResponse](../../models/operations/listportalsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePortal

Creates a new developer portal scoped in this region for this organization.

### Example Usage: Authentication Disabled

<!-- UsageSnippet language="go" operationID="create-portal" method="post" path="/v3/portals" example="Authentication Disabled" -->
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

    res, err := s.Portals.CreatePortal(ctx, components.CreatePortal{
        Name: "MyDevPortal",
        AuthenticationEnabled: sdkkonnectgo.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: Auto Approve Settings

<!-- UsageSnippet language="go" operationID="create-portal" method="post" path="/v3/portals" example="Auto Approve Settings" -->
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

    res, err := s.Portals.CreatePortal(ctx, components.CreatePortal{
        Name: "MyDevPortal",
        AutoApproveDevelopers: sdkkonnectgo.Pointer(true),
        AutoApproveApplications: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: Enable RBAC

<!-- UsageSnippet language="go" operationID="create-portal" method="post" path="/v3/portals" example="Enable RBAC" -->
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

    res, err := s.Portals.CreatePortal(ctx, components.CreatePortal{
        Name: "MyDevPortal",
        RbacEnabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: Example 1

<!-- UsageSnippet language="go" operationID="create-portal" method="post" path="/v3/portals" example="Example 1" -->
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

    res, err := s.Portals.CreatePortal(ctx, components.CreatePortal{
        Name: "MyDevPortal",
        RbacEnabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="create-portal" method="post" path="/v3/portals" example="Unauthorized" -->
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

    res, err := s.Portals.CreatePortal(ctx, components.CreatePortal{
        Name: "<value>",
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-portal" method="post" path="/v3/portals" example="UnauthorizedExample" -->
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

    res, err := s.Portals.CreatePortal(ctx, components.CreatePortal{
        Name: "<value>",
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdatePortalBadRequestExample1

<!-- UsageSnippet language="go" operationID="create-portal" method="post" path="/v3/portals" example="UpdatePortalBadRequestExample1" -->
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

    res, err := s.Portals.CreatePortal(ctx, components.CreatePortal{
        Name: "<value>",
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdatePortalBadRequestExample2

<!-- UsageSnippet language="go" operationID="create-portal" method="post" path="/v3/portals" example="UpdatePortalBadRequestExample2" -->
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

    res, err := s.Portals.CreatePortal(ctx, components.CreatePortal{
        Name: "<value>",
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdatePortalBadRequestExample3

<!-- UsageSnippet language="go" operationID="create-portal" method="post" path="/v3/portals" example="UpdatePortalBadRequestExample3" -->
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

    res, err := s.Portals.CreatePortal(ctx, components.CreatePortal{
        Name: "<value>",
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdatePortalBadRequestExample4

<!-- UsageSnippet language="go" operationID="create-portal" method="post" path="/v3/portals" example="UpdatePortalBadRequestExample4" -->
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

    res, err := s.Portals.CreatePortal(ctx, components.CreatePortal{
        Name: "<value>",
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdatePortalBadRequestExample5

<!-- UsageSnippet language="go" operationID="create-portal" method="post" path="/v3/portals" example="UpdatePortalBadRequestExample5" -->
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

    res, err := s.Portals.CreatePortal(ctx, components.CreatePortal{
        Name: "<value>",
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [components.CreatePortal](../../models/components/createportal.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.CreatePortalResponse](../../models/operations/createportalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPortal

Returns the configuration for a single developer portal, including the current visibility, access, and domain settings.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal" method="get" path="/v3/portals/{portalId}" example="Example 1" -->
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

    res, err := s.Portals.GetPortal(ctx, "1672e313-9f8e-4485-a62b-716aad351889")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPortalResponse](../../models/operations/getportalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortal

Updates the configuration for a single portal including the visibility, access, and custom domain settings.

### Example Usage: Authentication Disabled

<!-- UsageSnippet language="go" operationID="update-portal" method="patch" path="/v3/portals/{portalId}" example="Authentication Disabled" -->
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

    res, err := s.Portals.UpdatePortal(ctx, "06a06521-85cd-4704-8a94-628a34be93a4", components.UpdatePortal{
        AuthenticationEnabled: sdkkonnectgo.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: Auto Approve Settings

<!-- UsageSnippet language="go" operationID="update-portal" method="patch" path="/v3/portals/{portalId}" example="Auto Approve Settings" -->
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

    res, err := s.Portals.UpdatePortal(ctx, "c33fb6aa-eefa-4b2c-839d-ae0c33151e65", components.UpdatePortal{
        AutoApproveDevelopers: sdkkonnectgo.Pointer(true),
        AutoApproveApplications: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: Enable RBAC

<!-- UsageSnippet language="go" operationID="update-portal" method="patch" path="/v3/portals/{portalId}" example="Enable RBAC" -->
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

    res, err := s.Portals.UpdatePortal(ctx, "89fecbfd-98af-4be4-a961-762758c1cda0", components.UpdatePortal{
        RbacEnabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: Example 1

<!-- UsageSnippet language="go" operationID="update-portal" method="patch" path="/v3/portals/{portalId}" example="Example 1" -->
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

    res, err := s.Portals.UpdatePortal(ctx, "efe853b6-a55c-4929-9b80-c39a22b01c7a", components.UpdatePortal{
        AuthenticationEnabled: sdkkonnectgo.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: NotFoundExample

<!-- UsageSnippet language="go" operationID="update-portal" method="patch" path="/v3/portals/{portalId}" example="NotFoundExample" -->
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

    res, err := s.Portals.UpdatePortal(ctx, "11e433bf-d599-4a1d-95d3-d79b8659de0e", components.UpdatePortal{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="update-portal" method="patch" path="/v3/portals/{portalId}" example="Unauthorized" -->
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

    res, err := s.Portals.UpdatePortal(ctx, "1bbd136e-032f-41e0-ac85-574f86daa79b", components.UpdatePortal{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-portal" method="patch" path="/v3/portals/{portalId}" example="UnauthorizedExample" -->
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

    res, err := s.Portals.UpdatePortal(ctx, "02bccb4f-0606-4eb1-afe4-af597707eae1", components.UpdatePortal{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdatePortalBadRequestExample1

<!-- UsageSnippet language="go" operationID="update-portal" method="patch" path="/v3/portals/{portalId}" example="UpdatePortalBadRequestExample1" -->
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

    res, err := s.Portals.UpdatePortal(ctx, "b646206d-66de-4dee-8106-381967797727", components.UpdatePortal{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdatePortalBadRequestExample2

<!-- UsageSnippet language="go" operationID="update-portal" method="patch" path="/v3/portals/{portalId}" example="UpdatePortalBadRequestExample2" -->
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

    res, err := s.Portals.UpdatePortal(ctx, "e2802368-801b-4ad3-994f-f01954a4f348", components.UpdatePortal{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdatePortalBadRequestExample3

<!-- UsageSnippet language="go" operationID="update-portal" method="patch" path="/v3/portals/{portalId}" example="UpdatePortalBadRequestExample3" -->
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

    res, err := s.Portals.UpdatePortal(ctx, "0abeae52-4205-4262-bf1c-ce816dc7730f", components.UpdatePortal{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdatePortalBadRequestExample4

<!-- UsageSnippet language="go" operationID="update-portal" method="patch" path="/v3/portals/{portalId}" example="UpdatePortalBadRequestExample4" -->
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

    res, err := s.Portals.UpdatePortal(ctx, "a92fd996-1a8b-4414-a3fb-9333666a001f", components.UpdatePortal{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdatePortalBadRequestExample5

<!-- UsageSnippet language="go" operationID="update-portal" method="patch" path="/v3/portals/{portalId}" example="UpdatePortalBadRequestExample5" -->
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

    res, err := s.Portals.UpdatePortal(ctx, "f0c9976f-d0f3-42ad-9f78-b7cd2eb521ae", components.UpdatePortal{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `portalID`                                                         | *string*                                                           | :heavy_check_mark:                                                 | ID of the portal.                                                  |
| `updatePortal`                                                     | [components.UpdatePortal](../../models/components/updateportal.md) | :heavy_check_mark:                                                 | Update a portal's settings.                                        |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.UpdatePortalResponse](../../models/operations/updateportalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePortal

Deletes a single portal, along with all related entities.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-portal" method="delete" path="/v3/portals/{portalId}" -->
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

    res, err := s.Portals.DeletePortal(ctx, "304c8bf0-1c10-414d-af00-7cb8c3204da0", operations.DeletePortalQueryParamForceFalse.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                 | Type                                                                                                                                                                                      | Required                                                                                                                                                                                  | Description                                                                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                        | The context to use for the request.                                                                                                                                                       |
| `portalID`                                                                                                                                                                                | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | ID of the portal.                                                                                                                                                                         |
| `force`                                                                                                                                                                                   | [*operations.DeletePortalQueryParamForce](../../models/operations/deleteportalqueryparamforce.md)                                                                                         | :heavy_minus_sign:                                                                                                                                                                        | If true, the portal will be deleted, automatically deleting all API publications. If the force param is not set, the deletion will only succeed if there are no APIs currently published. |
| `opts`                                                                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                        | The options for this request.                                                                                                                                                             |

### Response

**[*operations.DeletePortalResponse](../../models/operations/deleteportalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |