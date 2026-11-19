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

	res, err := s.AttributionPlugins.GetCostAttributionPlugin(ctx, "527bc328-6950-4662-a473-78dd7a5beb84")
	if err != nil {
		log.Fatal(err)
	}
	if res.AICostAttributionPlugin != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->