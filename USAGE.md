<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)
	request := components.CreateControlPlaneRequest{
		Name:         "Test Control Plane",
		Description:  sdkkonnectgo.String("A test control plane for exploration."),
		ClusterType:  components.ClusterTypeClusterTypeControlPlane.ToPointer(),
		AuthType:     components.AuthTypePinnedClientCerts.ToPointer(),
		CloudGateway: sdkkonnectgo.Bool(false),
		ProxyUrls: []components.ProxyURL{
			components.ProxyURL{
				Host:     "example.com",
				Port:     443,
				Protocol: "https",
			},
		},
		Labels: map[string]string{
			"env": "test",
		},
	}
	ctx := context.Background()
	res, err := s.ControlPlanes.CreateControlPlane(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.ControlPlane != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->