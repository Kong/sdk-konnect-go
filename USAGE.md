<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	sdkkonnectgointernal "github.com/Kong/sdk-konnect-go-internal"
	"github.com/Kong/sdk-konnect-go-internal/models/components"
	"github.com/Kong/sdk-konnect-go-internal/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkkonnectgointernal.New(
		sdkkonnectgointernal.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgointernal.String("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.ControlPlanes.List(ctx, operations.ListControlPlanesRequest{
		FilterNameEq:         sdkkonnectgointernal.String("test"),
		FilterName:           sdkkonnectgointernal.String("test"),
		FilterNameContains:   sdkkonnectgointernal.String("test"),
		FilterNameNeq:        sdkkonnectgointernal.String("test"),
		FilterIDEq:           sdkkonnectgointernal.String("7f9fd312-a987-4628-b4c5-bb4f4fddd5f7"),
		FilterID:             sdkkonnectgointernal.String("7f9fd312-a987-4628-b4c5-bb4f4fddd5f7"),
		FilterIDOeq:          sdkkonnectgointernal.String("some-value,some-other-value"),
		FilterClusterTypeEq:  sdkkonnectgointernal.String("CLUSTER_TYPE_CONTROL_PLANE"),
		FilterClusterType:    sdkkonnectgointernal.String("CLUSTER_TYPE_CONTROL_PLANE"),
		FilterClusterTypeNeq: sdkkonnectgointernal.String("test"),
		Labels:               sdkkonnectgointernal.String("key:value,existCheck"),
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