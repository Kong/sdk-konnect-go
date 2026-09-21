# DataPlaneNodeLogLevelOperations

## Overview

### Available Operations

* [ListDataplaneNodeLogLevelOperations](#listdataplanenodelogleveloperations) - List Data Plane Node Log Level Operations
* [CreateDataplaneNodeLogLevelOperation](#createdataplanenodelogleveloperation) - Create a data plane node log level operation
* [GetDataplaneNodeLogLevelOperation](#getdataplanenodelogleveloperation) - Get a data plane node log level operation
* [ListDataplaneNodeLogLevelOperationResults](#listdataplanenodelogleveloperationresults) - List results of a data plane node log level operation
* [GetDataplaneNodeLogLevelOperationResult](#getdataplanenodelogleveloperationresult) - Get result of a data plane node log level operation

## ListDataplaneNodeLogLevelOperations

Returns the log level operations requested for the data plane nodes connected to this control plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-dataplane-node-log-level-operations" method="get" path="/v2/control-planes/{controlPlaneId}/nodes/log-level-operations" -->
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

    res, err := s.DataPlaneNodeLogLevelOperations.ListDataplaneNodeLogLevelOperations(ctx, operations.ListDataplaneNodeLogLevelOperationsRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DataPlaneNodeLogLevelOperationList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.ListDataplaneNodeLogLevelOperationsRequest](../../models/operations/listdataplanenodelogleveloperationsrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.ListDataplaneNodeLogLevelOperationsResponse](../../models/operations/listdataplanenodelogleveloperationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateDataplaneNodeLogLevelOperation

Temporarily overrides the log level on one or more data plane nodes connected to this control plane. The override reverts to the node's configured log level once the requested `ttl` elapses.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-dataplane-node-log-level-operation" method="post" path="/v2/control-planes/{controlPlaneId}/nodes/log-level-operations" -->
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

    res, err := s.DataPlaneNodeLogLevelOperations.CreateDataplaneNodeLogLevelOperation(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.DataPlaneNodeLogLevelOperationRequest{
        LogLevel: components.GatewayLogLevelInfo,
        Targets: components.CreateTargetsDataPlaneNodeLogLevelTargetNodeIds(
            components.DataPlaneNodeLogLevelTargetNodeIds{
                NodeIds: []string{
                    "9524ec7d-36d9-465d-a8c5-83a3c9390458",
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DataPlaneNodeLogLevelOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          | Example                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |                                                                                                                      |
| `controlPlaneID`                                                                                                     | `string`                                                                                                             | :heavy_check_mark:                                                                                                   | The UUID of your control plane. This variable is available in the Konnect manager.                                   | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                 |
| `dataPlaneNodeLogLevelOperationRequest`                                                                              | [components.DataPlaneNodeLogLevelOperationRequest](../../models/components/dataplanenodelogleveloperationrequest.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |                                                                                                                      |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |                                                                                                                      |

### Response

**[*operations.CreateDataplaneNodeLogLevelOperationResponse](../../models/operations/createdataplanenodelogleveloperationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDataplaneNodeLogLevelOperation

Returns the details of a data plane node log level operation requested for the data plane nodes connected to this control plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-dataplane-node-log-level-operation" method="get" path="/v2/control-planes/{controlPlaneId}/nodes/log-level-operations/{operationId}" -->
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

    res, err := s.DataPlaneNodeLogLevelOperations.GetDataplaneNodeLogLevelOperation(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.DataPlaneNodeLogLevelOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `operationID`                                                                      | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of the log level override operation.                                      | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetDataplaneNodeLogLevelOperationResponse](../../models/operations/getdataplanenodelogleveloperationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDataplaneNodeLogLevelOperationResults

Returns the results of a data plane node log level operation requested for the data plane nodes connected to this control plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-dataplane-node-log-level-operation-results" method="get" path="/v2/control-planes/{controlPlaneId}/nodes/log-level-operations/{operationId}/results" -->
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

    res, err := s.DataPlaneNodeLogLevelOperations.ListDataplaneNodeLogLevelOperationResults(ctx, operations.ListDataplaneNodeLogLevelOperationResultsRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        OperationID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DataPlaneNodeLogLevelOperationResultList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.ListDataplaneNodeLogLevelOperationResultsRequest](../../models/operations/listdataplanenodelogleveloperationresultsrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `opts`                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.ListDataplaneNodeLogLevelOperationResultsResponse](../../models/operations/listdataplanenodelogleveloperationresultsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDataplaneNodeLogLevelOperationResult

Returns the result of a data plane node log level operation requested for a specific data plane node connected to this control plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-dataplane-node-log-level-operation-result" method="get" path="/v2/control-planes/{controlPlaneId}/nodes/log-level-operations/{operationId}/results/{nodeId}" -->
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

    res, err := s.DataPlaneNodeLogLevelOperations.GetDataplaneNodeLogLevelOperationResult(ctx, operations.GetDataplaneNodeLogLevelOperationResultRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        OperationID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        NodeID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DataPlaneNodeLogLevelOperationResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.GetDataplaneNodeLogLevelOperationResultRequest](../../models/operations/getdataplanenodelogleveloperationresultrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.GetDataplaneNodeLogLevelOperationResultResponse](../../models/operations/getdataplanenodelogleveloperationresultresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |