# DestinationCreatePayload

Request body schema for creating a destination.


## Supported Types

### HTTPDestinationCreatePayload

```go
destinationCreatePayload := components.CreateDestinationCreatePayloadHTTP(components.HTTPDestinationCreatePayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch destinationCreatePayload.Type {
	case components.DestinationCreatePayloadTypeHTTP:
		// destinationCreatePayload.HTTPDestinationCreatePayload is populated
}
```
