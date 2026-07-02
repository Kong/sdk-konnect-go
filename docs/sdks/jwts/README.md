# JWTs

## Overview

### Available Operations

* [ListJwtWithConsumerInWorkspace](#listjwtwithconsumerinworkspace) - List all JWTs associated with a Consumer in a workspace
* [CreateJwtWithConsumerInWorkspace](#createjwtwithconsumerinworkspace) - Create a new JWT associated with a Consumer in a workspace
* [DeleteJwtWithConsumerInWorkspace](#deletejwtwithconsumerinworkspace) - Delete a a JWT associated with a Consumer in a workspace
* [GetJwtWithConsumerInWorkspace](#getjwtwithconsumerinworkspace) - Get a JWT associated with a Consumer in a workspace
* [UpsertJwtWithConsumerInWorkspace](#upsertjwtwithconsumerinworkspace) - Upsert a JWT associated with a Consumer in a workspace
* [ListJwtInWorkspace](#listjwtinworkspace) - List all JWTs in a workspace
* [GetJwtInWorkspace](#getjwtinworkspace) - Get a JWT in a workspace
* [ListJwtWithConsumer](#listjwtwithconsumer) - List all JWTs associated with a Consumer
* [CreateJwtWithConsumer](#createjwtwithconsumer) - Create a new JWT associated with a Consumer
* [DeleteJwtWithConsumer](#deletejwtwithconsumer) - Delete a a JWT associated with a Consumer
* [GetJwtWithConsumer](#getjwtwithconsumer) - Get a JWT associated with a Consumer
* [UpsertJwtWithConsumer](#upsertjwtwithconsumer) - Upsert a JWT associated with a Consumer
* [ListJwt](#listjwt) - List all JWTs
* [GetJwt](#getjwt) - Get a JWT

## ListJwtWithConsumerInWorkspace

List all JWTs associated with a Consumer in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-jwt-with-consumer-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/jwt" -->
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

    res, err := s.JWTs.ListJwtWithConsumerInWorkspace(ctx, operations.ListJwtWithConsumerInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ListJwtWithConsumerInWorkspaceRequest](../../models/operations/listjwtwithconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.ListJwtWithConsumerInWorkspaceResponse](../../models/operations/listjwtwithconsumerinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateJwtWithConsumerInWorkspace

Create a new JWT associated with a Consumer in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="create-jwt-with-consumer-in-workspace" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/jwt" -->
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

    res, err := s.JWTs.CreateJwtWithConsumerInWorkspace(ctx, operations.CreateJwtWithConsumerInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "",
        Workspace: "team-payments",
        JWTWithoutParents: &components.JWTWithoutParents{
            ID: sdkkonnectgo.Pointer("75695322-e8a0-4109-aed4-5416b0308d85"),
            Key: sdkkonnectgo.Pointer("YJdmaDvVTJxtcWRCvkMikc8oELgAVNcz"),
            Secret: sdkkonnectgo.Pointer("C50k0bcahDhLNhLKSUBSR1OMiFGzNZ7X"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Jwt != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.CreateJwtWithConsumerInWorkspaceRequest](../../models/operations/createjwtwithconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.CreateJwtWithConsumerInWorkspaceResponse](../../models/operations/createjwtwithconsumerinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteJwtWithConsumerInWorkspace

Delete a a JWT associated with a Consumer using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-jwt-with-consumer-in-workspace" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/jwt/{JWTId}" -->
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

    res, err := s.JWTs.DeleteJwtWithConsumerInWorkspace(ctx, operations.DeleteJwtWithConsumerInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "",
        JWTID: "4a7f5faa-8c96-46d6-8214-c87573ef2ac4",
        Workspace: "team-payments",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeleteJwtWithConsumerInWorkspaceRequest](../../models/operations/deletejwtwithconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.DeleteJwtWithConsumerInWorkspaceResponse](../../models/operations/deletejwtwithconsumerinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetJwtWithConsumerInWorkspace

Get a JWT associated with a Consumer using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-jwt-with-consumer-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/jwt/{JWTId}" -->
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

    res, err := s.JWTs.GetJwtWithConsumerInWorkspace(ctx, operations.GetJwtWithConsumerInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "",
        JWTID: "4a7f5faa-8c96-46d6-8214-c87573ef2ac4",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Jwt != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetJwtWithConsumerInWorkspaceRequest](../../models/operations/getjwtwithconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.GetJwtWithConsumerInWorkspaceResponse](../../models/operations/getjwtwithconsumerinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertJwtWithConsumerInWorkspace

Create or Update a JWT associated with a Consumer using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-jwt-with-consumer-in-workspace" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/jwt/{JWTId}" -->
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

    res, err := s.JWTs.UpsertJwtWithConsumerInWorkspace(ctx, operations.UpsertJwtWithConsumerInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "",
        JWTID: "4a7f5faa-8c96-46d6-8214-c87573ef2ac4",
        Workspace: "team-payments",
        JWTWithoutParents: &components.JWTWithoutParents{
            ID: sdkkonnectgo.Pointer("75695322-e8a0-4109-aed4-5416b0308d85"),
            Key: sdkkonnectgo.Pointer("YJdmaDvVTJxtcWRCvkMikc8oELgAVNcz"),
            Secret: sdkkonnectgo.Pointer("C50k0bcahDhLNhLKSUBSR1OMiFGzNZ7X"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Jwt != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.UpsertJwtWithConsumerInWorkspaceRequest](../../models/operations/upsertjwtwithconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpsertJwtWithConsumerInWorkspaceResponse](../../models/operations/upsertjwtwithconsumerinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListJwtInWorkspace

List all JWTs in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-jwt-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/jwts" -->
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

    res, err := s.JWTs.ListJwtInWorkspace(ctx, operations.ListJwtInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListJwtInWorkspaceRequest](../../models/operations/listjwtinworkspacerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListJwtInWorkspaceResponse](../../models/operations/listjwtinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetJwtInWorkspace

Get a JWT using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-jwt-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/jwts/{JWTId}" -->
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

    res, err := s.JWTs.GetJwtInWorkspace(ctx, operations.GetJwtInWorkspaceRequest{
        JWTID: "4a7f5faa-8c96-46d6-8214-c87573ef2ac4",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Jwt != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetJwtInWorkspaceRequest](../../models/operations/getjwtinworkspacerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetJwtInWorkspaceResponse](../../models/operations/getjwtinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## ListJwtWithConsumer

List all JWTs associated with a Consumer

### Example Usage

<!-- UsageSnippet language="go" operationID="list-jwt-with-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/jwt" -->
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

    res, err := s.JWTs.ListJwtWithConsumer(ctx, operations.ListJwtWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListJwtWithConsumerRequest](../../models/operations/listjwtwithconsumerrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListJwtWithConsumerResponse](../../models/operations/listjwtwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateJwtWithConsumer

Create a new JWT associated with a Consumer

### Example Usage

<!-- UsageSnippet language="go" operationID="create-jwt-with-consumer" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/jwt" -->
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

    res, err := s.JWTs.CreateJwtWithConsumer(ctx, operations.CreateJwtWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        JWTWithoutParents: &components.JWTWithoutParents{
            ID: sdkkonnectgo.Pointer("75695322-e8a0-4109-aed4-5416b0308d85"),
            Key: sdkkonnectgo.Pointer("YJdmaDvVTJxtcWRCvkMikc8oELgAVNcz"),
            Secret: sdkkonnectgo.Pointer("C50k0bcahDhLNhLKSUBSR1OMiFGzNZ7X"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Jwt != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateJwtWithConsumerRequest](../../models/operations/createjwtwithconsumerrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateJwtWithConsumerResponse](../../models/operations/createjwtwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteJwtWithConsumer

Delete a a JWT associated with a Consumer using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-jwt-with-consumer" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/jwt/{JWTId}" -->
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

    res, err := s.JWTs.DeleteJwtWithConsumer(ctx, operations.DeleteJwtWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        JWTID: "4a7f5faa-8c96-46d6-8214-c87573ef2ac4",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.DeleteJwtWithConsumerRequest](../../models/operations/deletejwtwithconsumerrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.DeleteJwtWithConsumerResponse](../../models/operations/deletejwtwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetJwtWithConsumer

Get a JWT associated with a Consumer using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-jwt-with-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/jwt/{JWTId}" -->
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

    res, err := s.JWTs.GetJwtWithConsumer(ctx, operations.GetJwtWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        JWTID: "4a7f5faa-8c96-46d6-8214-c87573ef2ac4",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Jwt != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetJwtWithConsumerRequest](../../models/operations/getjwtwithconsumerrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetJwtWithConsumerResponse](../../models/operations/getjwtwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertJwtWithConsumer

Create or Update a JWT associated with a Consumer using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-jwt-with-consumer" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/jwt/{JWTId}" -->
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

    res, err := s.JWTs.UpsertJwtWithConsumer(ctx, operations.UpsertJwtWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        JWTID: "4a7f5faa-8c96-46d6-8214-c87573ef2ac4",
        JWTWithoutParents: &components.JWTWithoutParents{
            ID: sdkkonnectgo.Pointer("75695322-e8a0-4109-aed4-5416b0308d85"),
            Key: sdkkonnectgo.Pointer("YJdmaDvVTJxtcWRCvkMikc8oELgAVNcz"),
            Secret: sdkkonnectgo.Pointer("C50k0bcahDhLNhLKSUBSR1OMiFGzNZ7X"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Jwt != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpsertJwtWithConsumerRequest](../../models/operations/upsertjwtwithconsumerrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpsertJwtWithConsumerResponse](../../models/operations/upsertjwtwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListJwt

List all JWTs

### Example Usage

<!-- UsageSnippet language="go" operationID="list-jwt" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/jwts" -->
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

    res, err := s.JWTs.ListJwt(ctx, operations.ListJwtRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.ListJwtRequest](../../models/operations/listjwtrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.ListJwtResponse](../../models/operations/listjwtresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetJwt

Get a JWT using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-jwt" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/jwts/{JWTId}" -->
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

    res, err := s.JWTs.GetJwt(ctx, "4a7f5faa-8c96-46d6-8214-c87573ef2ac4", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.Jwt != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `jwtID`                                                                            | `string`                                                                           | :heavy_check_mark:                                                                 | ID of the JWT to lookup                                                            | 4a7f5faa-8c96-46d6-8214-c87573ef2ac4                                               |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetJwtResponse](../../models/operations/getjwtresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |