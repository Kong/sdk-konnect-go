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

	res, err := s.CatalogServiceAPIMappings.ListServiceMappingsForAPI(ctx, operations.ListServiceMappingsForAPIRequest{
		APIID:      "d687f4ea-aa04-4157-b446-34519e5b18a7",
		PageSize:   sdkkonnectgo.Pointer[int64](10),
		PageNumber: sdkkonnectgo.Pointer[int64](1),
		Sort:       sdkkonnectgo.Pointer("created_at desc"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListCatalogServiceAPIMappingsResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->