# ResourceActionFilterParametersIntegrationInstanceID


## Supported Types

### StringFieldFilterExact

```go
resourceActionFilterParametersIntegrationInstanceID := components.CreateResourceActionFilterParametersIntegrationInstanceIDStringFieldFilterExact(components.StringFieldFilterExact{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceActionFilterParametersIntegrationInstanceID.Type {
	case components.ResourceActionFilterParametersIntegrationInstanceIDTypeStringFieldFilterExact:
		// resourceActionFilterParametersIntegrationInstanceID.StringFieldFilterExact is populated
}
```
