# APISpecPayload


## Supported Types

### APISpecContentPayload

```go
apiSpecPayload := components.CreateAPISpecPayloadAPISpecContentPayload(components.APISpecContentPayload{/* values here */})
```

### APISpecProviderPayload

```go
apiSpecPayload := components.CreateAPISpecPayloadAPISpecProviderPayload(components.APISpecProviderPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch apiSpecPayload.Type {
	case components.APISpecPayloadTypeAPISpecContentPayload:
		// apiSpecPayload.APISpecContentPayload is populated
	case components.APISpecPayloadTypeAPISpecProviderPayload:
		// apiSpecPayload.APISpecProviderPayload is populated
}
```
