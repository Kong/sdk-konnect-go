# MapActionPayloadService


## Supported Types

### MapByName

```go
mapActionPayloadService := components.CreateMapActionPayloadServiceMapByName(components.MapByName{/* values here */})
```

### MapByID

```go
mapActionPayloadService := components.CreateMapActionPayloadServiceMapByID(components.MapByID{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mapActionPayloadService.Type {
	case components.MapActionPayloadServiceTypeMapByName:
		// mapActionPayloadService.MapByName is populated
	case components.MapActionPayloadServiceTypeMapByID:
		// mapActionPayloadService.MapByID is populated
}
```
