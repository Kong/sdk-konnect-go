<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)
	request := operations.ListServerlessCloudGatewayRequest{
		PageSize:   sdkkonnectgo.Int64(10),
		PageNumber: sdkkonnectgo.Int64(1),
		Labels:     sdkkonnectgo.String("filter[labels][eq]=env:prod"),
	}
	ctx := context.Background()
	res, err := s.ServerlessCloudGateways.ListServerlessCloudGateway(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListServerlessCloudGatewaysResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->