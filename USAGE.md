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
			PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.Notifications.ListUserConfigurations(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.UserConfigurationListResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->