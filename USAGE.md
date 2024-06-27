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
	request := operations.ListControlPlanesRequest{
		PageSize:             sdkkonnectgo.Int64(10),
		PageNumber:           sdkkonnectgo.Int64(1),
		FilterNameEq:         sdkkonnectgo.String("test"),
		FilterName:           sdkkonnectgo.String("test"),
		FilterNameContains:   sdkkonnectgo.String("test"),
		FilterNameNeq:        sdkkonnectgo.String("test"),
		FilterIDEq:           sdkkonnectgo.String("7f9fd312-a987-4628-b4c5-bb4f4fddd5f7"),
		FilterID:             sdkkonnectgo.String("7f9fd312-a987-4628-b4c5-bb4f4fddd5f7"),
		FilterIDOeq:          sdkkonnectgo.String("some-value,some-other-value"),
		FilterClusterTypeEq:  sdkkonnectgo.String("CLUSTER_TYPE_CONTROL_PLANE"),
		FilterClusterType:    sdkkonnectgo.String("CLUSTER_TYPE_CONTROL_PLANE"),
		FilterClusterTypeNeq: sdkkonnectgo.String("test"),
		Labels:               sdkkonnectgo.String("key:value,existCheck"),
	}
	ctx := context.Background()
	res, err := s.ControlPlanes.ListControlPlanes(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListControlPlanesResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->