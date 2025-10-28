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
	ctx := context.Background()

	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.APIGatewayDataPlaneCertificates.ListAPIGatewayDataPlaneCertificates(ctx, operations.ListAPIGatewayDataPlaneCertificatesRequest{
		GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
		PageSize:  sdkkonnectgo.Pointer[int64](10),
		PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListAPIGatewayDataPlaneCertificatesResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->