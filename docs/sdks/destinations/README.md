# Destinations

## Overview

Operations related to http destinations and their deliveries. Only an org admin can use these operations. The destination routes exist only when http destinations are enabled.

### Available Operations

* [ListDeliveries](#listdeliveries) - List delivery history
* [ListDestinations](#listdestinations) - List destinations
* [CreateDestination](#createdestination) - Create a destination
* [GetDestination](#getdestination) - Get a destination
* [UpdateDestination](#updatedestination) - Update a destination
* [DeleteDestination](#deletedestination) - Delete a destination
* [GetDestinationDelivery](#getdestinationdelivery) - Get a delivery and its attempts
* [ResendDestinationDelivery](#resenddestinationdelivery) - Resend a delivery
* [RotateDestinationSecret](#rotatedestinationsecret) - Rotate the destination signing secret
* [ListDestinationSubscriptions](#listdestinationsubscriptions) - List destination subscriptions
* [TestDestination](#testdestination) - Test a destination

## ListDeliveries

List the deliveries for all destinations in this organization. To filter to a specific destination, use `filter[destination_id][eq]`

### Example Usage

<!-- UsageSnippet language="go" operationID="list-deliveries" method="get" path="/v1/notifications/deliveries" -->
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

    res, err := s.Destinations.ListDeliveries(ctx, operations.ListDeliveriesRequest{
        PageBefore: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        Filter: &components.OrgDeliveryFilterParameters{
            Status: &components.StatusFieldFilter{
                Eq: sdkkonnectgo.Pointer("?filter[status][eq]=failed"),
                Neq: sdkkonnectgo.Pointer("?filter[status][neq]=failed"),
                Contains: sdkkonnectgo.Pointer("?filter[status][contains]=fail"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeliveryListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListDeliveriesRequest](../../models/operations/listdeliveriesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListDeliveriesResponse](../../models/operations/listdeliveriesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListDestinations

Returns the organization's destinations, with the newest destinations first. Use the event type filter to list destinations with an enabled subscription for a specific event type.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-destinations" method="get" path="/v1/notifications/destinations" -->
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

    res, err := s.Destinations.ListDestinations(ctx, operations.ListDestinationsRequest{
        PageBefore: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DestinationListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListDestinationsRequest](../../models/operations/listdestinationsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListDestinationsResponse](../../models/operations/listdestinationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateDestination

Creates an HTTP destination for organization event notifications. You can create up to 50 destinations for an organization.
The response includes the signing secret. Save it when you create the destination; later responses do not include it.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-destination" method="post" path="/v1/notifications/destinations" -->
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

    res, err := s.Destinations.CreateDestination(ctx, components.CreateDestinationCreatePayloadHTTP(
        components.HTTPDestinationCreatePayload{
            Name: "Production webhooks",
            Type: components.HTTPDestinationCreatePayloadTypeHTTP,
            URL: "https://example.com/webhooks",
            CustomHeaders: map[string]string{
                "X-Environment": "production",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.DestinationWithSecret != nil {
        switch res.DestinationWithSecret.Type {
            case components.DestinationWithSecretTypeHTTP:
                // res.DestinationWithSecret.HTTPDestinationWithSecret is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.DestinationCreatePayload](../../models/components/destinationcreatepayload.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateDestinationResponse](../../models/operations/createdestinationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetDestination

Show a destination's details. Does not return sensitive keys such as signing secret.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-destination" method="get" path="/v1/notifications/destinations/{destinationId}" -->
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

    res, err := s.Destinations.GetDestination(ctx, "d26034f1-0cdd-4b6f-9158-13c03f4cd006")
    if err != nil {
        log.Fatal(err)
    }
    if res.Destination != nil {
        switch res.Destination.Type {
            case components.DestinationTypeHTTP:
                // res.Destination.HTTPDestination is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `destinationID`                                          | `string`                                                 | :heavy_check_mark:                                       | ID of the http destination, a UUID v7.                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetDestinationResponse](../../models/operations/getdestinationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateDestination

Update a destination's parameters

### Example Usage

<!-- UsageSnippet language="go" operationID="update-destination" method="patch" path="/v1/notifications/destinations/{destinationId}" -->
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

    res, err := s.Destinations.UpdateDestination(ctx, "310711be-def8-442b-89d4-1677c4fb3473", components.CreateDestinationUpdatePayloadHTTPDestinationUpdatePayload(
        components.HTTPDestinationUpdatePayload{
            Type: components.HTTPDestinationUpdatePayloadTypeHTTP.ToPointer(),
            Name: sdkkonnectgo.Pointer("Production webhooks"),
            CustomHeaders: map[string]*string{
                "X-Environment": sdkkonnectgo.Pointer("staging"),
            },
            State: components.HTTPDestinationUpdatePayloadStateEnabled.ToPointer(),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Destination != nil {
        switch res.Destination.Type {
            case components.DestinationTypeHTTP:
                // res.Destination.HTTPDestination is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `destinationID`                                                                            | `string`                                                                                   | :heavy_check_mark:                                                                         | ID of the http destination, a UUID v7.                                                     |
| `destinationUpdatePayload`                                                                 | [components.DestinationUpdatePayload](../../models/components/destinationupdatepayload.md) | :heavy_check_mark:                                                                         | Request body schema for updating a destination.                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateDestinationResponse](../../models/operations/updatedestinationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteDestination

Deletes the destination and removes it from all subscriptions that use it.
If this leaves a subscription with no destinations, that subscription is also deleted.
This action also deletes the destination's signing keys and delivery history. You cannot undo this action.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-destination" method="delete" path="/v1/notifications/destinations/{destinationId}" -->
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

    res, err := s.Destinations.DeleteDestination(ctx, "f535bc11-b5c3-4614-b8f4-a9d563f17078")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `destinationID`                                          | `string`                                                 | :heavy_check_mark:                                       | ID of the http destination, a UUID v7.                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteDestinationResponse](../../models/operations/deletedestinationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetDestinationDelivery

Show a notification's delivery status, including information about each delivery attempt.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-destination-delivery" method="get" path="/v1/notifications/destinations/{destinationId}/deliveries/{deliveryId}" -->
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

    res, err := s.Destinations.GetDestinationDelivery(ctx, "4e6f10e2-b5b7-4642-a338-bfeec5a0f74d", "70c31e56-5ade-409c-bac3-95884d81ed86")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeliveryDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `destinationID`                                          | `string`                                                 | :heavy_check_mark:                                       | ID of the http destination, a UUID v7.                   |
| `deliveryID`                                             | `string`                                                 | :heavy_check_mark:                                       | ID of the delivery, a UUID v7.                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetDestinationDeliveryResponse](../../models/operations/getdestinationdeliveryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ResendDestinationDelivery

Trigger a redelivery of a failed notification. The redelivery has a new delivery ID in the "Delivery history" endpoint.
The payload is identical to the original delivery, including the `event_id` which can be used to deduplicate notifications.

### Example Usage

<!-- UsageSnippet language="go" operationID="resend-destination-delivery" method="post" path="/v1/notifications/destinations/{destinationId}/deliveries/{deliveryId}/resend" -->
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

    res, err := s.Destinations.ResendDestinationDelivery(ctx, "0c0bdaad-5a5e-4f8a-9f99-503d665b1ee6", "c5c24211-1c53-478c-9f22-9be0d9e20a22")
    if err != nil {
        log.Fatal(err)
    }
    if res.Delivery != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `destinationID`                                          | `string`                                                 | :heavy_check_mark:                                       | ID of the http destination, a UUID v7.                   |
| `deliveryID`                                             | `string`                                                 | :heavy_check_mark:                                       | ID of the delivery, a UUID v7.                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ResendDestinationDeliveryResponse](../../models/operations/resenddestinationdeliveryresponse.md), error**

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

## RotateDestinationSecret

Creates a new signing secret for the HTTP destination. The response includes the new secret only once. The previous secret remains valid for 24 hours so you can update your integration.

### Example Usage

<!-- UsageSnippet language="go" operationID="rotate-destination-secret" method="post" path="/v1/notifications/destinations/{destinationId}/secret/rotate" -->
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

    res, err := s.Destinations.RotateDestinationSecret(ctx, "3cf186ca-b834-4e41-a770-136cdef73322")
    if err != nil {
        log.Fatal(err)
    }
    if res.DestinationSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `destinationID`                                          | `string`                                                 | :heavy_check_mark:                                       | ID of the http destination, a UUID v7.                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RotateDestinationSecretResponse](../../models/operations/rotatedestinationsecretresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListDestinationSubscriptions

Returns the organization's subscriptions that use this destination.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-destination-subscriptions" method="get" path="/v1/notifications/destinations/{destinationId}/subscriptions" -->
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

    res, err := s.Destinations.ListDestinationSubscriptions(ctx, operations.ListDestinationSubscriptionsRequest{
        DestinationID: "e2be6e2d-f3ab-4205-aa48-7a0052bfec36",
        PageBefore: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgSubscriptionListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListDestinationSubscriptionsRequest](../../models/operations/listdestinationsubscriptionsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListDestinationSubscriptionsResponse](../../models/operations/listdestinationsubscriptionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## TestDestination

Trigger a test delivery to a destination. The test outcome is returned in the "Delivery history" endpoint.
Tests may be triggered every 5 seconds.

### Example Usage

<!-- UsageSnippet language="go" operationID="test-destination" method="post" path="/v1/notifications/destinations/{destinationId}/test" -->
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

    res, err := s.Destinations.TestDestination(ctx, "92133a48-f45f-426c-86c6-a72eae55a903")
    if err != nil {
        log.Fatal(err)
    }
    if res.Delivery != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `destinationID`                                          | `string`                                                 | :heavy_check_mark:                                       | ID of the http destination, a UUID v7.                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TestDestinationResponse](../../models/operations/testdestinationresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |