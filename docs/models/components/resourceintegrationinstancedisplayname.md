# ResourceIntegrationInstanceDisplayName


## Supported Types

### StringFieldFilter

```go
resourceIntegrationInstanceDisplayName := components.CreateResourceIntegrationInstanceDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceIntegrationInstanceDisplayName.Type {
	case components.ResourceIntegrationInstanceDisplayNameTypeStringFieldFilter:
		// resourceIntegrationInstanceDisplayName.StringFieldFilter is populated
}
```
