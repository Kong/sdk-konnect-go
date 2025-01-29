<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkkonnectgo.New()

	res, err := s.CloudGateways.GetAvailabilityJSON(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.AvailabilityDocument != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->