# IntegrationDisplayName


## Supported Types

### StringFieldFilter

```go
integrationDisplayName := components.CreateIntegrationDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationDisplayName.Type {
	case components.IntegrationDisplayNameTypeStringFieldFilter:
		// integrationDisplayName.StringFieldFilter is populated
}
```
