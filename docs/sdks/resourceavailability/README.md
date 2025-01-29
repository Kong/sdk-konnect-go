# ResourceAvailability
(*ResourceAvailability*)

## Overview

### Available Operations

* [GetAvailabilityJSON](#getavailabilityjson) - Get Resource Availability JSON

## GetAvailabilityJSON

Get Cloud Gateways Availability JSON document for describing cloud provider and region availability, pricing,
gateway version availability, and instance type information.


### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New()

    res, err := s.ResourceAvailability.GetAvailabilityJSON(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.AvailabilityDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAvailabilityJSONResponse](../../models/operations/getavailabilityjsonresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |