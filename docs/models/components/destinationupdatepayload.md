# DestinationUpdatePayload

Request body schema for updating a destination.


## Supported Types

### HTTPDestinationUpdatePayload

```go
destinationUpdatePayload := components.CreateDestinationUpdatePayloadHTTPDestinationUpdatePayload(components.HTTPDestinationUpdatePayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch destinationUpdatePayload.Type {
	case components.DestinationUpdatePayloadTypeHTTPDestinationUpdatePayload:
		// destinationUpdatePayload.HTTPDestinationUpdatePayload is populated
}
```
