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
	ctx := context.Background()

	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.CatalogServices.CreateCatalogService(ctx, components.CreateCatalogService{
		Name:        "user-svc",
		DisplayName: "User Service",
		Labels: map[string]string{
			"env": "test",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.CatalogService != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->