# AIGatewayDatastores

## Overview

Datastores that back AI Gateway policies and models, such as Redis and vector database connections.

### Available Operations

* [ListAiGatewayDatastores](#listaigatewaydatastores) - List AI Gateway Datastores
* [CreateAiGatewayDatastore](#createaigatewaydatastore) - Create an AI Gateway Datastore
* [GetAiGatewayDatastore](#getaigatewaydatastore) - Get an AI Gateway Datastore
* [UpdateAiGatewayDatastore](#updateaigatewaydatastore) - Update an AI Gateway Datastore
* [DeleteAiGatewayDatastore](#deleteaigatewaydatastore) - Delete an AI Gateway Datastore
* [ListAiGatewayDatastoreUsage](#listaigatewaydatastoreusage) - List AI Gateway Datastore Usage

## ListAiGatewayDatastores

Returns a list of Datastores associated with the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-datastores" method="get" path="/v1/ai-gateways/{gatewayId}/datastores" -->
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

    res, err := s.AIGatewayDatastores.ListAiGatewayDatastores(ctx, operations.ListAiGatewayDatastoresRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayDatastoresResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListAiGatewayDatastoresRequest](../../models/operations/listaigatewaydatastoresrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListAiGatewayDatastoresResponse](../../models/operations/listaigatewaydatastoresresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayDatastore

Creates a new Datastore for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-datastore" method="post" path="/v1/ai-gateways/{gatewayId}/datastores" -->
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

    res, err := s.AIGatewayDatastores.CreateAiGatewayDatastore(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.CreateCreateAIGatewayDatastoreRequestRedisEe(
        components.RedisEEDatastore{
            Name: "shared-redis",
            DisplayName: "Shared Redis",
            Description: sdkkonnectgo.Pointer("Redis shared by rate limiting and semantic cache"),
            Config: components.RedisEEDatastoreConfig{
                Host: "127.0.0.1",
                SentinelNodes: []components.SentinelNodes{
                    components.SentinelNodes{
                        Host: "127.0.0.1",
                        Port: 6379,
                    },
                },
                ClusterNodes: []components.ClusterNodes{
                    components.ClusterNodes{
                        IP: "127.0.0.1",
                        Port: 6379,
                    },
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayDatastore != nil {
        switch res.AIGatewayDatastore.Type {
            case components.AIGatewayDatastoreTypeRedisCe:
                // res.AIGatewayDatastore.AIGatewayDatastoreRedisCEDatastore is populated
            case components.AIGatewayDatastoreTypeRedisEe:
                // res.AIGatewayDatastore.AIGatewayDatastoreRedisEEDatastore is populated
            case components.AIGatewayDatastoreTypeVectordb:
                // res.AIGatewayDatastore.AIGatewayDatastoreVectorDBDatastore is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |                                                                                                          |
| `gatewayID`                                                                                              | `string`                                                                                                 | :heavy_check_mark:                                                                                       | The unique ID of the AI Gateway.                                                                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                     |
| `createAIGatewayDatastoreRequest`                                                                        | [components.CreateAIGatewayDatastoreRequest](../../models/components/createaigatewaydatastorerequest.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |                                                                                                          |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |                                                                                                          |

### Response

**[*operations.CreateAiGatewayDatastoreResponse](../../models/operations/createaigatewaydatastoreresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## GetAiGatewayDatastore

Returns the details of a specific AI Gateway Datastore.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-datastore" method="get" path="/v1/ai-gateways/{gatewayId}/datastores/{datastoreIdOrName}" -->
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

    res, err := s.AIGatewayDatastores.GetAiGatewayDatastore(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayDatastore != nil {
        switch res.AIGatewayDatastore.Type {
            case components.AIGatewayDatastoreTypeRedisCe:
                // res.AIGatewayDatastore.AIGatewayDatastoreRedisCEDatastore is populated
            case components.AIGatewayDatastoreTypeRedisEe:
                // res.AIGatewayDatastore.AIGatewayDatastoreRedisEEDatastore is populated
            case components.AIGatewayDatastoreTypeVectordb:
                // res.AIGatewayDatastore.AIGatewayDatastoreVectorDBDatastore is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `datastoreIDOrName`                                      | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway Datastore.       | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiGatewayDatastoreResponse](../../models/operations/getaigatewaydatastoreresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewayDatastore

Updates the configuration of an existing AI Gateway Datastore.

`name` and `type` are immutable after creation; attempting to change either is rejected.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-datastore" method="put" path="/v1/ai-gateways/{gatewayId}/datastores/{datastoreIdOrName}" -->
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

    res, err := s.AIGatewayDatastores.UpdateAiGatewayDatastore(ctx, operations.UpdateAiGatewayDatastoreRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        DatastoreIDOrName: "my-entity-name",
        UpdateAIGatewayDatastoreRequest: components.CreateUpdateAIGatewayDatastoreRequestVectordb(
            components.VectorDBDatastore{
                Name: "shared-redis",
                DisplayName: "Shared Redis",
                Description: sdkkonnectgo.Pointer("Redis shared by rate limiting and semantic cache"),
                Config: components.VectorDBDatastoreConfig{
                    Host: "127.0.0.1",
                    Database: "kong-pgvector",
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayDatastore != nil {
        switch res.AIGatewayDatastore.Type {
            case components.AIGatewayDatastoreTypeRedisCe:
                // res.AIGatewayDatastore.AIGatewayDatastoreRedisCEDatastore is populated
            case components.AIGatewayDatastoreTypeRedisEe:
                // res.AIGatewayDatastore.AIGatewayDatastoreRedisEEDatastore is populated
            case components.AIGatewayDatastoreTypeVectordb:
                // res.AIGatewayDatastore.AIGatewayDatastoreVectorDBDatastore is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateAiGatewayDatastoreRequest](../../models/operations/updateaigatewaydatastorerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateAiGatewayDatastoreResponse](../../models/operations/updateaigatewaydatastoreresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayDatastore

Deletes an AI Gateway Datastore. Rejected with a 409 if the Datastore is still referenced by a policy or model.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-datastore" method="delete" path="/v1/ai-gateways/{gatewayId}/datastores/{datastoreIdOrName}" -->
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

    res, err := s.AIGatewayDatastores.DeleteAiGatewayDatastore(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
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
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `datastoreIDOrName`                                      | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway Datastore.       | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiGatewayDatastoreResponse](../../models/operations/deleteaigatewaydatastoreresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## ListAiGatewayDatastoreUsage

Returns the policies and models that reference this Datastore.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-datastore-usage" method="get" path="/v1/ai-gateways/{gatewayId}/datastores/{datastoreIdOrName}/usage" -->
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

    res, err := s.AIGatewayDatastores.ListAiGatewayDatastoreUsage(ctx, operations.ListAiGatewayDatastoreUsageRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        DatastoreIDOrName: "my-entity-name",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayDatastoreUsageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListAiGatewayDatastoreUsageRequest](../../models/operations/listaigatewaydatastoreusagerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListAiGatewayDatastoreUsageResponse](../../models/operations/listaigatewaydatastoreusageresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |