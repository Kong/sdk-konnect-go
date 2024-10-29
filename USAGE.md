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

	ctx := context.Background()
	res, err := s.ControlPlanes.ListControlPlanes(ctx, operations.ListControlPlanesRequest{
		PageSize:   sdkkonnectgo.Int64(10),
		PageNumber: sdkkonnectgo.Int64(1),
		Filter: &components.ControlPlaneFilterParameters{
			CloudGateway: sdkkonnectgo.Bool(true),
		},
		Labels: sdkkonnectgo.String("key:value,existCheck"),
		Sort:   sdkkonnectgo.String("name,created_at desc"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListControlPlanesResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->