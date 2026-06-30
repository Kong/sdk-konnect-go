# MTLSAuthCredentials

## Overview

### Available Operations

* [ListMtlsAuthWithConsumerInWorkspace](#listmtlsauthwithconsumerinworkspace) - List all MTLS-auth credentials associated with a Consumer in a workspace
* [CreateMtlsAuthWithConsumerInWorkspace](#createmtlsauthwithconsumerinworkspace) - Create a new MTLS-auth credential associated with a Consumer in a workspace
* [DeleteMtlsAuthWithConsumerInWorkspace](#deletemtlsauthwithconsumerinworkspace) - Delete a a MTLS-auth credential associated with a Consumer in a workspace
* [GetMtlsAuthWithConsumerInWorkspace](#getmtlsauthwithconsumerinworkspace) - Get a MTLS-auth credential associated with a Consumer in a workspace
* [UpsertMtlsAuthWithConsumerInWorkspace](#upsertmtlsauthwithconsumerinworkspace) - Upsert a MTLS-auth credential associated with a Consumer in a workspace
* [ListMtlsAuthInWorkspace](#listmtlsauthinworkspace) - List all MTLS-auth credentials in a workspace
* [GetMtlsAuthInWorkspace](#getmtlsauthinworkspace) - Get a MTLS-auth credential in a workspace
* [ListMtlsAuthWithConsumer](#listmtlsauthwithconsumer) - List all MTLS-auth credentials associated with a Consumer
* [CreateMtlsAuthWithConsumer](#createmtlsauthwithconsumer) - Create a new MTLS-auth credential associated with a Consumer
* [DeleteMtlsAuthWithConsumer](#deletemtlsauthwithconsumer) - Delete a a MTLS-auth credential associated with a Consumer
* [GetMtlsAuthWithConsumer](#getmtlsauthwithconsumer) - Get a MTLS-auth credential associated with a Consumer
* [UpsertMtlsAuthWithConsumer](#upsertmtlsauthwithconsumer) - Upsert a MTLS-auth credential associated with a Consumer
* [ListMtlsAuth](#listmtlsauth) - List all MTLS-auth credentials
* [GetMtlsAuth](#getmtlsauth) - Get a MTLS-auth credential

## ListMtlsAuthWithConsumerInWorkspace

List all MTLS-auth credentials associated with a Consumer in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-mtls-auth-with-consumer-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/mtls-auth" -->
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

    res, err := s.MTLSAuthCredentials.ListMtlsAuthWithConsumerInWorkspace(ctx, operations.ListMtlsAuthWithConsumerInWorkspaceRequest{
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.ListMtlsAuthWithConsumerInWorkspaceRequest](../../models/operations/listmtlsauthwithconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.ListMtlsAuthWithConsumerInWorkspaceResponse](../../models/operations/listmtlsauthwithconsumerinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateMtlsAuthWithConsumerInWorkspace

Create a new MTLS-auth credential associated with a Consumer in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="create-mtls-auth-with-consumer-in-workspace" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/mtls-auth" -->
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

    res, err := s.MTLSAuthCredentials.CreateMtlsAuthWithConsumerInWorkspace(ctx, operations.CreateMtlsAuthWithConsumerInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "",
        Workspace: "team-payments",
        MTLSAuthWithoutParents: components.MTLSAuthWithoutParents{
            CaCertificate: &components.MTLSAuthWithoutParentsCaCertificate{
                ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f260"),
            },
            ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f269"),
            SubjectName: "CA_Subject_Name",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MTLSAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.CreateMtlsAuthWithConsumerInWorkspaceRequest](../../models/operations/createmtlsauthwithconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.CreateMtlsAuthWithConsumerInWorkspaceResponse](../../models/operations/createmtlsauthwithconsumerinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteMtlsAuthWithConsumerInWorkspace

Delete a a MTLS-auth credential associated with a Consumer using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-mtls-auth-with-consumer-in-workspace" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/mtls-auth/{MTLSAuthId}" -->
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

    res, err := s.MTLSAuthCredentials.DeleteMtlsAuthWithConsumerInWorkspace(ctx, operations.DeleteMtlsAuthWithConsumerInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "",
        MTLSAuthID: "",
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.DeleteMtlsAuthWithConsumerInWorkspaceRequest](../../models/operations/deletemtlsauthwithconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.DeleteMtlsAuthWithConsumerInWorkspaceResponse](../../models/operations/deletemtlsauthwithconsumerinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMtlsAuthWithConsumerInWorkspace

Get a MTLS-auth credential associated with a Consumer using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mtls-auth-with-consumer-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/mtls-auth/{MTLSAuthId}" -->
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

    res, err := s.MTLSAuthCredentials.GetMtlsAuthWithConsumerInWorkspace(ctx, operations.GetMtlsAuthWithConsumerInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "",
        MTLSAuthID: "",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MTLSAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetMtlsAuthWithConsumerInWorkspaceRequest](../../models/operations/getmtlsauthwithconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.GetMtlsAuthWithConsumerInWorkspaceResponse](../../models/operations/getmtlsauthwithconsumerinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertMtlsAuthWithConsumerInWorkspace

Create or Update a MTLS-auth credential associated with a Consumer using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-mtls-auth-with-consumer-in-workspace" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/mtls-auth/{MTLSAuthId}" -->
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

    res, err := s.MTLSAuthCredentials.UpsertMtlsAuthWithConsumerInWorkspace(ctx, operations.UpsertMtlsAuthWithConsumerInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "",
        MTLSAuthID: "",
        Workspace: "team-payments",
        MTLSAuthWithoutParents: components.MTLSAuthWithoutParents{
            CaCertificate: &components.MTLSAuthWithoutParentsCaCertificate{
                ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f260"),
            },
            ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f269"),
            SubjectName: "CA_Subject_Name",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MTLSAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.UpsertMtlsAuthWithConsumerInWorkspaceRequest](../../models/operations/upsertmtlsauthwithconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.UpsertMtlsAuthWithConsumerInWorkspaceResponse](../../models/operations/upsertmtlsauthwithconsumerinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMtlsAuthInWorkspace

List all MTLS-auth credentials in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-mtls-auth-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/mtls-auths" -->
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

    res, err := s.MTLSAuthCredentials.ListMtlsAuthInWorkspace(ctx, operations.ListMtlsAuthInWorkspaceRequest{
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListMtlsAuthInWorkspaceRequest](../../models/operations/listmtlsauthinworkspacerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListMtlsAuthInWorkspaceResponse](../../models/operations/listmtlsauthinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetMtlsAuthInWorkspace

Get a MTLS-auth credential using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mtls-auth-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/mtls-auths/{MTLSAuthId}" -->
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

    res, err := s.MTLSAuthCredentials.GetMtlsAuthInWorkspace(ctx, operations.GetMtlsAuthInWorkspaceRequest{
        MTLSAuthID: "",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MTLSAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetMtlsAuthInWorkspaceRequest](../../models/operations/getmtlsauthinworkspacerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetMtlsAuthInWorkspaceResponse](../../models/operations/getmtlsauthinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## ListMtlsAuthWithConsumer

List all MTLS-auth credentials associated with a Consumer

### Example Usage

<!-- UsageSnippet language="go" operationID="list-mtls-auth-with-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/mtls-auth" -->
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

    res, err := s.MTLSAuthCredentials.ListMtlsAuthWithConsumer(ctx, operations.ListMtlsAuthWithConsumerRequest{
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListMtlsAuthWithConsumerRequest](../../models/operations/listmtlsauthwithconsumerrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListMtlsAuthWithConsumerResponse](../../models/operations/listmtlsauthwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateMtlsAuthWithConsumer

Create a new MTLS-auth credential associated with a Consumer

### Example Usage

<!-- UsageSnippet language="go" operationID="create-mtls-auth-with-consumer" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/mtls-auth" -->
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

    res, err := s.MTLSAuthCredentials.CreateMtlsAuthWithConsumer(ctx, operations.CreateMtlsAuthWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        MTLSAuthWithoutParents: components.MTLSAuthWithoutParents{
            CaCertificate: &components.MTLSAuthWithoutParentsCaCertificate{
                ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f260"),
            },
            ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f269"),
            SubjectName: "CA_Subject_Name",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MTLSAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateMtlsAuthWithConsumerRequest](../../models/operations/createmtlsauthwithconsumerrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.CreateMtlsAuthWithConsumerResponse](../../models/operations/createmtlsauthwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteMtlsAuthWithConsumer

Delete a a MTLS-auth credential associated with a Consumer using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-mtls-auth-with-consumer" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/mtls-auth/{MTLSAuthId}" -->
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

    res, err := s.MTLSAuthCredentials.DeleteMtlsAuthWithConsumer(ctx, operations.DeleteMtlsAuthWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        MTLSAuthID: "",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.DeleteMtlsAuthWithConsumerRequest](../../models/operations/deletemtlsauthwithconsumerrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.DeleteMtlsAuthWithConsumerResponse](../../models/operations/deletemtlsauthwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMtlsAuthWithConsumer

Get a MTLS-auth credential associated with a Consumer using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mtls-auth-with-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/mtls-auth/{MTLSAuthId}" -->
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

    res, err := s.MTLSAuthCredentials.GetMtlsAuthWithConsumer(ctx, operations.GetMtlsAuthWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        MTLSAuthID: "",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MTLSAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetMtlsAuthWithConsumerRequest](../../models/operations/getmtlsauthwithconsumerrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GetMtlsAuthWithConsumerResponse](../../models/operations/getmtlsauthwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertMtlsAuthWithConsumer

Create or Update a MTLS-auth credential associated with a Consumer using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-mtls-auth-with-consumer" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/mtls-auth/{MTLSAuthId}" -->
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

    res, err := s.MTLSAuthCredentials.UpsertMtlsAuthWithConsumer(ctx, operations.UpsertMtlsAuthWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        MTLSAuthID: "",
        MTLSAuthWithoutParents: components.MTLSAuthWithoutParents{
            CaCertificate: &components.MTLSAuthWithoutParentsCaCertificate{
                ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f260"),
            },
            ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f269"),
            SubjectName: "CA_Subject_Name",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MTLSAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.UpsertMtlsAuthWithConsumerRequest](../../models/operations/upsertmtlsauthwithconsumerrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.UpsertMtlsAuthWithConsumerResponse](../../models/operations/upsertmtlsauthwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMtlsAuth

List all MTLS-auth credentials

### Example Usage

<!-- UsageSnippet language="go" operationID="list-mtls-auth" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/mtls-auths" -->
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

    res, err := s.MTLSAuthCredentials.ListMtlsAuth(ctx, operations.ListMtlsAuthRequest{
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListMtlsAuthRequest](../../models/operations/listmtlsauthrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListMtlsAuthResponse](../../models/operations/listmtlsauthresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetMtlsAuth

Get a MTLS-auth credential using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mtls-auth" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/mtls-auths/{MTLSAuthId}" -->
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

    res, err := s.MTLSAuthCredentials.GetMtlsAuth(ctx, "", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.MTLSAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `mtlsAuthID`                                                                       | `string`                                                                           | :heavy_check_mark:                                                                 | ID of the MTLS-auth credential to lookup                                           |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetMtlsAuthResponse](../../models/operations/getmtlsauthresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |